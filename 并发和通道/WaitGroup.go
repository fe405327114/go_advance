package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main(){
	runtime.GOMAXPROCS(2)
	//wg用来等待程序完成
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutine")
	go func(){
	defer wg.Done()
		for count:=0;count<3;count++ {
			for char:='a';char<'a'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()
	go func (){
		defer wg.Done()
	for count:=0;count<3;count++{
		for char:='A';char<'A'+26;char++{
			fmt.Printf("%c",char)
		}
	}
	}()
	fmt.Println("Waiting Finish")
	wg.Wait()
}
