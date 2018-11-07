package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn,err:=redis.Dial("tcp",":6379")
	defer conn.Close()
	if err!=nil{
		fmt.Println("redis 连接失败")
		return
	}
	data,_:=conn.Do("get","name")

	fmt.Println(data)
}
