package main

//练习3：	封装 wcFunc() 函数。接收一段英文字符串str。返回一个map，记录str中每个“词”出现次数的。
//如："I love my work and I love my family too"
import (
"strings"
"fmt"
)

func wcFunc(str string) map[string]int {
	m1 := make(map[string]int)
	s1 := strings.Split(str, " ")
	fmt.Println(s1)
	for _, v := range s1 {
		m1[v]++
	}
	return m1
}
func main() {
	str := "I love my work and I love my family too"
	m1 := wcFunc(str)
	fmt.Println(m1)
}
