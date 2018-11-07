package main

import "fmt"

func main()  {
	var i  int
	go func(){
	i++
	}()
	fmt.Println(i)
}
