package main

import "fmt"

func demo(s []string) {
	var s1 []string
	for i := 0; i < len(s); i++ {
		if s[i]!=""{
			s1=append(s1,s[i])
		}
	}
	fmt.Println(s1)
}
func main() {
	str := []string{"red", "", "black", "", "", "pink", "blue"}
	demo(str)
}
