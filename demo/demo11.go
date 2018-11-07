package main

import (
	"strconv"
	"fmt"
)

func main (){
	//将其他类型转换为字符串Format
	a:=strconv.FormatFloat(3.14,'f',2,64)
	fmt.Println(a)
	b,_:=strconv.ParseFloat("12.123",64)
	c,_:=strconv.Atoi("45")
	fmt.Println(b,c)
	s:=make([]byte,0,1024)
	s=strconv.AppendInt(s,124,10)
	s=strconv.AppendQuote(s,"fdsgsdg")
	fmt.Println(string(s))
}
