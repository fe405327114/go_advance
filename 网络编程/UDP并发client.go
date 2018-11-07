package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//请求连接
	conn, err := net.Dial("udp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//读取从键盘输入的内容
	b := make([]byte, 1024)
	go func() {
		for {
			n, err := os.Stdin.Read(b)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(b[:n]))
			conn.Write(b[:n]) //给服务器发送从键盘里读取到的内容,此处省去返回值
		}
	}()
	c := make([]byte, 1024)
	for {
	n,err:=conn.Read(c)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println("从服务器接收：",string(c[:n]))
	}

}
