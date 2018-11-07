package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg sync.WaitGroup

const (
	tasknum = 10
	tinenum = 4
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	task := make(chan string, tasknum)
	wg.Add(tinenum)
	//以工人为编号，启动goroutine工作
	//每个工人进入goroutine后，判断是否有任务，开始工作，后面的工人也依次进入
	for gr:=1;gr<=tinenum;gr++ {//注意此处是<=,因为主程序在等待4个goroutine
		go worker(task,gr)
	}
	//以任务为编号，向通道缓冲区中写入任务
	for rw:=1;rw<=tasknum;rw++{
		//fmt.Sprintf返回值是字符串
		task<-fmt.Sprintf("Task:%d",rw)
	}
	close(task)
	wg.Wait()
	}
func worker(task chan string,gr int){
	defer wg.Done()
	for{
		//判断通道缓冲区中是否有任务,此处会阻塞，等待任务
		value,ok:=<-task
		if!ok{
			fmt.Printf("%d worker: shutdown\n",gr)
			return
		}
		//有任务开始工作
		fmt.Printf("%d woker: start work\n",gr)
		time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		//工作结束
		fmt.Printf("%d worker: finished work %s\n",gr,value)
	}
}