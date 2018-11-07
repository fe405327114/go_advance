package main

import (
	"net"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("conn err", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
	conn.Write([]byte("yeah come on"))
}
