package main

import "fmt"

//练习2：写一个函数，就地消除[]string中重复字符串，如：
//	{"red", "black", "red", "pink", "blue", "pink", "blue"}
//	——>	{"red", "black", "pink", "blue"}

func delete01(str []string) []string {
	var s1 []string
	for _, word := range str {
		i:=0
		for ; i < len(s1); i++ {
			if word == s1[i] {
				break
			}
		}
		if i==len(s1) {
			s1 = append(s1, word)
		}
	}
	return s1
}
func main() {
	s2 := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	s1 := delete01(s2)
	fmt.Println(s1)
}
