package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

//同步等待组是为了每条通道一个一个卖票，以防每个goroutine都卖100、
// 可是在判定tickt>0 后，该gorotuine睡着，导致后边的goroutine进来判定依然执行，卖到了负数
//因此需要上锁，但是要注意要给"程序结束"方向去的goroutine也解锁

var tickt int = 10
var wg sync.WaitGroup
var mat sync.Mutex

func Saletickt(name string) {
	rand.Seed(time.Now().UnixNano())

	for {
		mat.Lock()
		if tickt > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "卖掉了", tickt)
			tickt--
		} else {
			fmt.Println(name, "售票结束")
			mat.Unlock()
			break
		}
		mat.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(4)
	go Saletickt("售票口1")
	go Saletickt("售票口2")
	go Saletickt("售票口3")
	go Saletickt("售票口4")
	wg.Wait()
	fmt.Println("程序结束")

}
