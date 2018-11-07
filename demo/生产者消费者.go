package main

import (
	"fmt"
)

type order struct {
	id int
}
func producer00(out chan<-order){
	defer close(out)
	for i:=0;i<10;i++  {
		order:=order{i+1}
		out<-order
	}
}
func consumer00(in <-chan order){
	for orderid:=range in {
		fmt.Println(orderid.id)
	}
}
func main(){
	ch:=make(chan order)
	go producer00(ch)
	consumer00(ch)
}