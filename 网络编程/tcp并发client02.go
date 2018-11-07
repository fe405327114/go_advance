package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("dial err")
		return
	}
	defer conn.Close()
	go  Readstdin(conn)
	buf:=make([]byte,4096)
	for {
		n,_:=conn.Read(buf)
		if n==0{
			break
		}
		fmt.Printf("receive data: %s",string(buf[:n]))
	}

}
func Readstdin(conn net.Conn) {
	buf := make([]byte, 4096)
	for {
		n, _ := os.Stdin.Read(buf)
		if n == 0 {
			return
		}
		conn.Write(buf[:n])
	}
}
