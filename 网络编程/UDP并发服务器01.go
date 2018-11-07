package main

import (
	"net"
	"fmt"
)

func main(){
	serAddr,err:=net.ResolveUDPAddr("udp","127.0.0.1:8005")
	if err!=nil{
		fmt.Println("resolve err",err)
		return
	}
	conn,err1:=net.ListenUDP("udp",serAddr)
	if err1!=nil{
		fmt.Println("listen err",err1)
		return
	}
	defer conn.Close()
	for{
		buf:=make([]byte,4096)
		n,cltAddr,err:=conn.ReadFromUDP(buf)
		if err!=nil{
			fmt.Println("read err",err)
			return
		}
		fmt.Println("server received",string(buf[:n]))
		go func() {
			_,err:=conn.WriteToUDP(buf,cltAddr)
			if err!=nil{
				fmt.Println("write err",err)
				return
			}
		}()
	}
}
