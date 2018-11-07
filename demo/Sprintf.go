package main

import (
	"fmt"
	"bytes"
)

func main(){
 	var name string="Amy"
 	//把格式化的字符串写入某个字符串缓冲区
 	name1:=fmt.Sprintf("I am %s,age %d\n",name,23)//与Printf比多了个返回值
 	name2:=fmt.Sprintln(name," HELLO world")
 	fmt.Println(name1,name2)
 	var b bytes.Buffer
 	b.Write([]byte("hello word"))
 	fmt.Fprint(&b,"haha")

 }
