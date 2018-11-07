package main

import (
	"os"
	"fmt"
)

func main(){
	list:=os.Args
	if len(list)!=2{
		fmt.Println("geshi err")
		return
	}
	fileinfo,err:=os.Stat(list[1])
	if err!=nil{
		fmt.Println("stat err",err)
		return
	}
	fmt.Println(fileinfo.Name(),fileinfo.Size())
}
