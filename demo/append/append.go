package main

import "fmt"

func test(s1 []int){
	s1=append(s1,10)
	fmt.Printf("===%p\n",s1)
}
func main(){
	s1:=make([]int,3,10)
	s1[0]=1
	fmt.Printf("%p\n",s1)
	s1=append(s1,2)
	fmt.Printf("%p\n",s1)
	test(s1)
}
