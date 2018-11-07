package main

import (
	"sync"
	"fmt"
	"math/rand"
	"time"
)
var C sync.Cond
func Reproducer(ch chan<- int, i int) {
	for {
		C.L.Lock()
		for len(ch) == 5 {
			C.Wait()
		}
		num:=rand.Intn(100)
		ch <- num
		fmt.Printf("%d write in %d ,length %d\n",i+1,num,len(ch))
		C.L.Unlock()
		C.Signal()
		//time.Sleep(time.Second)
	}
}
func Reconsumer(ch <-chan int, j int) {
	for {
		C.L.Lock()
		for len(ch) == 0 {
			C.Wait()
		}
		num := <-ch
		fmt.Printf("%d read out %d ,length %d\n",j+1,num,len(ch))
		C.L.Unlock()
		C.Signal()
		//time.Sleep(time.Millisecond*500)
	}
}
func main() {
	C.L=new(sync.Mutex)
	ch := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())
	quit:=make(chan bool)
	for i := 0; i < 5; i++ {
		go Reproducer(ch, i)
	}
	for j := 0; j < 3; j++ {
		go Reconsumer(ch, j)
	}
	<-quit
}
