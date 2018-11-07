package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	r:=time.NewTimer(time.Second*2)
	nowtime:=<-r.C
	fmt.Println(r)
	fmt.Println(nowtime)
	nowtome1:=<-time.After(2*time.Second)
	fmt.Println(nowtome1)


}