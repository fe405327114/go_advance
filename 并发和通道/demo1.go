package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("func goroutine:", i)
		}
		quit <- true
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		select {
		case num := <-ch:
			fmt.Println("main goroutine:", num)
		case <-quit:
			fmt.Println("quit")
		}
	}
}
