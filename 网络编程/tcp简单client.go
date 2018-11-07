package main

import (
	"net"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("Dial err", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("are you ready?"))
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
}
