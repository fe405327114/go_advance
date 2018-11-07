package main

import (
	"sync"
	"fmt"
	"runtime"
)

var wg sync.WaitGroup

func main() {
	//分配一个逻辑处理器给调度器
	runtime.GOMAXPROCS(runtime.NumCPU())
	//获取可用的物理处理器的数量
	//计数器+2，表示要等待两个goroutine
	wg.Add(2)
	//创建两个goroutine
	fmt.Println("Create Goroutine")
	go printPrime("A")
	go printPrime("B")
	//等待gorountine结束
	fmt.Println("Wait Goroutine")
	wg.Wait()
}

//显示5000以内的素数
func printPrime(prefix string) {
	defer wg.Done()
next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
