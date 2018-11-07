package main

import "fmt"

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	defer close(out)
}
func consumer(in <-chan int) {
	//for i :=0;i<len(in);i++{  // channel不能用for遍历，必须用range，返回值为值
	//	fmt.Println(<-in)
	//}
	for num:=range in{  //range便利通道返回值只有一个，那就是值，不可以用for i++遍历
		fmt.Println(num,<-in)
	}
}
func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

}