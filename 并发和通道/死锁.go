package main

//func main(){
//	ch:=make(chan int)
//	ch<-123
//	fmt.Println(<-ch)
//}
// 死锁2
//func main()  {
//	ch := make(chan int)
//	num := <- ch
//	go func() {
//		ch <- 789
//	}()
//	fmt.Println("num = ", num)
//}
func main()  {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {				// 子
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	for {
		select {
		case num := <- ch2:
			ch1 <- num
		}
	}
}