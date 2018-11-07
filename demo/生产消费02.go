package main

import "fmt"

type OrderInfo struct{
	id int
}
func producer1(out chan<-OrderInfo){
	for i:=0;i<10;i++{
		order:=OrderInfo{id:i+1}
		out<-order
	}
	 defer close(out)
}
func consumer1(in <-chan OrderInfo){
	for order:=range in{
		fmt.Println(order.id)
	}
}
func main(){
	ch:=make(chan OrderInfo)
	go producer1(ch)
	consumer1(ch)
}