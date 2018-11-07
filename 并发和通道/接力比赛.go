package main

import (
	"sync"
	"fmt"

	"time"
)

var wg sync.WaitGroup
func main() {
	wg.Add(1) //Add和Wait之间是要等待执行的goroutine
	baton := make(chan int)
	//此处如果用4条goroutine，则无法控制每一条goroutine执行的顺序
	go Runner(baton)
	baton<-1//在开始等待之前，将数据写入通道内
	wg.Wait()//等待比赛结束
}
func Runner(baton chan int){
	runner:=<-baton
	var newrunner int
	fmt.Printf("%d runner running with the baton\n",runner)

	if runner!=4{
		newrunner=runner+1
		fmt.Printf("%d runner to the line\n",newrunner)
		go Runner(baton)//等待baton中被写入新的内容
	}
	//围绕跑道跑步
	time.Sleep(100*time.Millisecond)

	if runner==4{
		fmt.Printf("%d runner finished race",runner)
		wg.Done()
		return
	}
	//将接力棒交给下一位
	fmt.Printf("%d runner exchange with %d runner\n",runner,newrunner)
	baton<-newrunner
}