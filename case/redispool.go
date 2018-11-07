package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

//定义一个全局的 pool
var pool *redis.Pool

//当启动程序时，初始化连接池
func init(){
	pool=& redis.Pool{
		MaxIdle:8,  //最大空闲连接数
		MaxActive:0,  //最大连接数，0表示不限制
		IdleTimeout:100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","localhost:6379")
		},

	}
}

func main() {
	//先从连接池中取出一个连接
   conn:=pool.Get()
   defer conn.Close()

   _,err:=conn.Do("Set","name","Tom")
    if err!=nil{
    	fmt.Println("conn Do err",err)
    	return
	}

	//取出
	r,err:=redis.String(conn.Do("Get","name"))
	if err!=nil{
		fmt.Println("conn Do err",err)
		return
	}
	fmt.Println(r)
   }
