package main

import "fmt"

//练习2：写一个函数，就地消除[]string中重复字符串，如：
//	{"red", "black", "red", "pink", "blue", "pink", "blue"}
//	——>	{"red", "black", "pink", "blue"}
func demo1(s []string) {
	var s1 []string
	for i := 0; i < len(s); i++ {
		count:=0
		for j := 0; j < len(s1); j++ {
			if s[i] == s1[j] {
				count++
				break
			}
		}
		if count==0{
			s1=append(s1,s[i])
		}
	}
	fmt.Println(s1)
}
func main() {
	s := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	demo1(s)
}
