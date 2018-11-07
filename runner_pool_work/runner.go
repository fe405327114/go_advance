package runner

import (
	"os"
	"time"
	"errors"
	"os/signal"
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

//创建中断的错误返回
var ErrInterrupt = errors.New("received interrupt")
//创建超时错误
var ErrTimeout = errors.New("received timeout")
//New方法返回一个新Runner 类型
// 参数是时间类型
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add方法将任务附加到Runner上，此任务是一个函数，函数以int类型为参数
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		//停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
