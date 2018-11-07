package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	//先listen等待客户端发起链接请求，返回conn接口
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//连接成功后进入方法处理数据
		go HindleConn(conn)
	}
}
func HindleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String() //conn接口中的RemoteAddr方法,返回值为Addr
	//接收数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s:%s", addr, string(buf[:n]))

	if string(buf[:n-1]) == "exit" { //去掉\r\n两个字符
		fmt.Println(addr, "exit")
		return
	}
	// 向客户端发送数据
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}
