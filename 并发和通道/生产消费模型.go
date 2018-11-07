package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Order struct {
	id int
}

func producer1(ch chan<- Order) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 10; i++ {
		order:=Order{id:i+1}
		ch <- order
	}
}
func consumer1(ch <-chan Order) {
	defer wg.Done()
	for num:=range ch {
		fmt.Println(num.id)
	}
}
func main() {
	wg.Add(2)
	ch := make(chan Order)
	go producer1(ch)
	go consumer1(ch)
	wg.Wait()
}
