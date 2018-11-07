package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("conn err")
			return
		}
		go HindleConn2(conn)
	}

}
func HindleConn2(conn net.Conn) {
	addrr := conn.RemoteAddr()
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			return
		}
		if "exit" == string(buf[:n-1]) {
			fmt.Println(addrr, "exit")
			return
		}
		//将用户发送的内容打印出来
		fmt.Printf("%s send %s", addrr, string(buf[:n]))
		//转成大写发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
