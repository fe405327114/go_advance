package main

import "fmt"

func main(){
	arr:=[8]int{1,2,3,4,5,6,7,8}
	s:=arr[2:5] //切片在截取时如果没有指定容量，则cap=原数组cap-low
	s1:=make([]int,6)//切片创建时cap=len
	fmt.Println(s)
	fmt.Println(cap(s))  //8-2
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s[0]=100
	s2:=s[1:2]
	fmt.Println(s2)
	s2[0]=999
	fmt.Println(arr)
}
