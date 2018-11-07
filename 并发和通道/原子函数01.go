package main

import (
	"sync"
	"fmt"
	"time"
	"sync/atomic"
)

var wg sync.WaitGroup
var shutdown int64

func main() {
	wg.Add(2)
	fmt.Println("Start Goroutine")
	go dowork("A")
	go dowork("B")
	time.Sleep(1 * time.Second)
	//利用原子函数atomic.Store
	atomic.StoreInt64(&shutdown,1)
	wg.Wait()
}
func dowork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("%s正在工作", name)
		time.Sleep(250 * time.Millisecond)
		// 利用atomic.LoadInt64来检查shutdown的值
		if atomic.LoadInt64(&shutdown) == 1 {//1秒后 Store函数给shutdown 赋值为1
			fmt.Printf("%s Finish", name)
			break
		}
	}
}