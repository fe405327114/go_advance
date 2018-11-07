package main

import (
	"math/rand"
	"sync"
	"fmt"
	"time"
)

var value int
var rwlock sync.RWMutex

func Readgo() {
	rwlock.RLock()
	num := value
	fmt.Printf("======%d 读", num)
	rwlock.RUnlock()
}
func Writego() {
	for {
		num := rand.Intn(1000)
		rwlock.RUnlock()
		value = num
		fmt.Printf("%d 写", num)
		rwlock.RUnlock()
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<5;i++{
		go 	Readgo()
	}
	for j:=0;j<5;j++{
		go Writego()
	}
	for{
		;
	}
}
