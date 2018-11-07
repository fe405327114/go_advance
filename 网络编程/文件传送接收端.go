package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("accept err", err)
		return
	}
	defer conn.Close()
	//接收文件名
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fileName := string(buf[:n])
	//回复一个ok
	conn.Write([]byte("ok"))
	//接收文件内容
	receiveFile(conn, fileName)
}
func receiveFile(conn net.Conn, fileName string) {
	//创建文件
	f,err:=os.Create("E:/学习资料/讲义/"+fileName)
	if err!=nil{
		fmt.Println("create err",err)
		return
	}
	defer f.Close()
	//从接口中读取数据
	buf := make([]byte, 4096)
	for {
		n, err:= conn.Read(buf)
		if n == 0 {
			fmt.Println("received success")
			return
		}
		if err!=nil{
			fmt.Println("read err",err)
		}
		f.Write(buf[:n])
	}
}
