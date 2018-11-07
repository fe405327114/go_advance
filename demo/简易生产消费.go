package main

import "fmt"

func producer (out chan<-int){
	for i:=0;i<10;i++ {
		out<-i
	}
	defer close(out)
}
func consumer (in <-chan int){
	for num:=range in {  //从通道中返回的值没有下标
		fmt.Println(num)
	}
}
func main(){
	c:=make(chan int)

	go producer(c)
	consumer(c)

}