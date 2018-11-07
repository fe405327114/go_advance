package main

import (
	"os"
	"fmt"
)

func main(){
	var a string
	for i:=0;i<len(os.Args);i++{
		a=""
		a+=os.Args[i]
	}
	fmt.Println(a)
}
