package main

import "fmt"

//练习3：要删除slice中间的某个元素并保存原有的元素顺序， 如：
//	{5, 6, 7, 8, 9} ——> {5, 6, 8, 9}

func remove(s []int,i int)[]int{
	copy(s[i:],s[i+1:])
	return s
}
func main(){
	s:=[]int{5, 6, 7, 8, 9}
	s1:=remove(s,2)
	fmt.Println(s[:len(s1)-1])
	s3:=[]int{1,2}
	s4:=[]int{3,4,5,6}
	copy(s3,s4)
	fmt.Println(s3)
}
