package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup // 同步等待组是为了先让子协程运行，防止主程序先结束
var mutex sync.Mutex
var ticket int = 10

func main() {
	wg.Add(4)
	fmt.Println("Start Sale")

	go SaleTicket("window1")
	go SaleTicket("window2")
	go SaleTicket("window3")
	go SaleTicket("window4")
	wg.Wait()
	fmt.Println("Over")
}
func SaleTicket(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		mutex.Lock()
		if ticket <= 0 {
			fmt.Printf("%s Finish\n", name)
			mutex.Unlock()
			break
		}
		fmt.Printf("%s Saling %d\n", name, ticket)
		ticket--
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		mutex.Unlock()
	}
	wg.Done()
}
