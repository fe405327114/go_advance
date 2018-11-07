package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	//1111111111发起链接请求
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("conn err")
		return
	}
	defer conn.Close()
	//22222222222读取标准输入设备Stdin中的内容，写入接口
	go func() {
		buf := make([]byte, 4096)
		for {
		n, _ := os.Stdin.Read(buf)
		if n==0{
			return
		}
		conn.Write(buf[:n])
		}
	}()
	//33333333第三次握手，接收服务器返回的消息
	buf2:=make([]byte,4096)
	for {
		n,_:=conn.Read(buf2)
		if n==0{
			return
		}
		fmt.Printf("从服务器接收的数据%s",string(buf2[:n]))
	}
}
