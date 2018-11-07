package 指针切片map结构体

import (
	"strings"
	"fmt"
)

//练习3：	封装 wcFunc() 函数。接收一段英文字符串str。返回一个map，记录str中每个“词”出现次数的。
//如："I love my work and I love my family too"
//输出：
//family : 1
//too : 1
//I : 2
//love : 2
//my : 2
//work : 1
//and : 1
//提示：使用 strings.Fields() 函数可提高效率。
func wcFunc(str string) map[string]int {
	map1 := make(map[string]int)
	s1 := strings.Fields(str)
	for _, v := range s1 {
		value, ok := map1[v]
		if ok {
			value++
		} else {
			value = 1
		}
		map1[v] = value
	}
	return map1
}
func main() {
	str := "I love my work and I love my family too"
	map1 := wcFunc(str)
	fmt.Println(map1)
}
