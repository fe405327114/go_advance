package main

import (
	"fmt"
	"runtime"
)

func demo12(ch chan string,quit chan bool){
	 defer close(ch)
	for i:=0;i<5;i++  {
		ch<-"hello world"
	}
	quit<-true
	runtime.Goexit()
}
func main() {
	ch:=make(chan string,3)
	quit :=make(chan bool)
	go demo12(ch,quit)
	for{
	select {
	case num:=<-ch:
		fmt.Println(num)
	case <-quit:
		return
	}
}
}