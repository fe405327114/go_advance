package main

import (
	"time"
	"fmt"
)

func main (){
	time1:=time.NewTimer(time.Second*2)
	t1:=time.Now()
	fmt.Printf("%v\n",t1)

	t2:=<-time1.C  //2秒后从通道中接收值
	fmt.Printf("%v\n",t2)

	time.Sleep(time.Second*2)
	fmt.Printf("%v\n",t2)

	t3:=<-time.After(time.Second*2)  //等于先time.NewTimer(),延迟后再从time.C中接收
	fmt.Println(t3)
}
