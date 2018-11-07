package main

import (
	"fmt"
	"strconv"
)

func test(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {

	go fmt.Println("ss："+strconv.Itoa(test(10,20)))

	for i:=0;i<5;i++{
		fmt.Println(i)
	}
}

