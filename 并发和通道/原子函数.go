package main

import (
	"sync"
	"fmt"
	"runtime"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64
func main (){
	wg.Add(2)
	fmt.Println("Start Goroutine")
	go inCounter(1)
	go inCounter(2)
	wg.Wait()
	fmt.Println("Final Counter",counter)
}
func inCounter(id int){
	defer wg.Done()
	for count:=0;count<2;count++{
		atomic.AddInt64(&counter,1)
		//counter++
	}
	runtime.Gosched()
}
