package 指针切片map结构体

import "fmt"

//练习1：给定一个字符串列表，在原有slice上返回不包含空字符串的列表， 如：
//	{"red", "", "black", "", "", "pink", "blue"}
//——> {"red", "black", "pink", "blue"}
func main() {
	str := []string{"red", "", "black", "", "", "pink", "blue"}
	var s1 []string
	for i := 0; i < len(str); i++ {
		if str[i] == "" {
			s1 = append(str[:i], str[i+1:]...)//如果str[i]为空格则将后一个值覆盖str[i],
			// i--是再次判定此值
			i--
		}
	}
	fmt.Println(s1)
}
