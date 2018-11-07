package main

import (
	"regexp"
	"fmt"

)

func main() {

	str := `<title>标题</title>
	<div>过年来吃鸡啊</div>jkbjvhjv
	<div>hello regexp</div>  knk

	<div>你在吗？</div>
	<div>
		2块钱啥时候还？
	过了年再说吧！
	刚买了车，没钱。。。
	</div>
	<body>呵呵</body>`
//250-255  200-250 100-200 0-99`
	var re = regexp.MustCompile(`<div>(?s:(.*?))</div>`)
	result:=re.FindAllStringSubmatch(str, -1)
	for _,data:=range result{
		fmt.Println(data[0])
		fmt.Println(data[1])
	}
	//fmt.Println(result)
}
