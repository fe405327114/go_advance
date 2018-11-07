package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch =make(chan int)
func printer(s string) {
	for _, v := range s {
		fmt.Printf("%c", v)
	}
}
func person1() {
	defer wg.Done()
	printer("hello")
	fmt.Println(<-ch)
}
func person2() {
	defer wg.Done()
	ch<-8
	printer("world")
}
func main() {
	wg.Add(2)
	go person1()
	go person2()
	wg.Wait()

}
