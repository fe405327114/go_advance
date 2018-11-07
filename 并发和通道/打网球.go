package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	court := make(chan int)
	go HitBall("David", court)
	go HitBall("Amy", court)
	court <- 1
	wg.Wait()
}
func HitBall(name string, court chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for {
		ball, ok := <-court //判断通道里是否有值
		if !ok {
			fmt.Printf("%s Win!", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s Missed The Ball\n", name)
			close(court)
			return
		}
		fmt.Printf("%s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}