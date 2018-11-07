package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var cond sync.Cond

func Pro(out chan<- int) {
	for {
		cond.L.Lock()
		for len(out) == 3 {
			cond.Wait()
		}
		num := rand.Intn(100)
		out <- num
		fmt.Printf("Write in %d\n", num)
		cond.L.Unlock()
		cond.Signal()
	}
}

func Con(in <-chan int) {
	for {
		cond.L.Lock()
		for len(in) == 0 {
			cond.Wait()
		}
		fmt.Printf("read out %d\n", <-in)
		cond.L.Unlock()
		cond.Signal()
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int, 3)
	cond.L = new(sync.Mutex)
	for i := 0; i < 3; i++ {
		go Pro(ch)
	}
	for i := 0; i < 5; i++ {
		go Con(ch)
	}
	for{
		;
	}
}
