package main

import "fmt"

type order struct {
	id int
}
func producer(out chan<- order) {
	defer close(out)
	for i:=0;i<10;i++{
		order:=order{i+1}
		out<-order
		fmt.Println("write in",order.id)
}

}
func consumer(in <-chan order){
	//for i:=0;i<10;i++{
	//	order:=<-in
	//	fmt.Println("receive",order.id)
	//}
	for order:=range in{
		fmt.Println("receive",order.id)
	}
}
func main(){
	ch:=make(chan order)
	//quit:=make(chan bool)
	go producer(ch)
	consumer(ch)
}