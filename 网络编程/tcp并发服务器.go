package main

import (
	"net"
	"fmt"
	"strings"
)

func main(){
	listener,err:=net.Listen("tcp","127.0.0.1:8004")
	if err!=nil{
		fmt.Println("listen err",err)
		return
	}
	defer listener.Close()
	for{
		fmt.Println("waiting for connect")
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("conn err",err)
			return
		}
		defer conn.Close()
		go ConnHindle(conn)
	}
}
func ConnHindle(conn net.Conn){
	defer conn.Close()
	addr:=conn.RemoteAddr()
	fmt.Println(addr,"conncet success")
	buf:=make([]byte,4096)
	for{
		n,_:=conn.Read(buf)
		if n==0{
			return
		}
		fmt.Println("recive data",string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
