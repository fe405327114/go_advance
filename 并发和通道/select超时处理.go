package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan bool)
	go func() {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-time.After(3 * time.Second):
			fmt.Println("time out")
			goto lable
			return
		}
	lable:
	}()
	for i := 0; i < 3; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
}