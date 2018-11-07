package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

var con sync.Cond // 声明全局条件变量
func producer(out chan<- int, i int) {
	for {
		con.L.Lock()
		for len(out) == 5 {
			con.Wait()
		}
		num := rand.Intn(100)
		out <- num
		fmt.Printf("第%d个生产者正在写入数据%d，剩余空间%d\n", i+1, num, len(out))
		con.L.Unlock()
		con.Signal()
		time.Sleep(time.Second)
	}
}
func consumer(in <-chan int, i int) {
	for {
		con.L.Lock()
		for len(in) == 0 {
			con.Wait()
		}
		num := <-in
		fmt.Printf("第%d个消费者正在读取数据%d，剩余空间%d\n", i+1, num, len(in))
		con.L.Unlock()
		con.Signal()
		time.Sleep(time.Millisecond*500)
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	con.L = new(sync.Mutex) //生成互斥锁
	product := make(chan int,5) //此处如果是无缓冲区的通道，则无法执行
	quit := make(chan bool)
	for i := 0; i < 5; i++ {
		go producer(product, i)
	}
	for i := 0; i < 3; i++ {
		go consumer(product, i)
	}
	<-quit
}
