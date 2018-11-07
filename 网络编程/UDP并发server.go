package main

import (
	"net"
	"fmt"
)

func main(){
//指定监听的地址
	addr,err:=net.ResolveUDPAddr("udp","127.0.0.1:8002")//resolve指定
	if err!=nil{
		fmt.Println(err)
		return
	}
	//客户端连接
	conn,err:=net.ListenUDP("udp",addr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//读取从客户端接收的数据
	for{
		buf:=make([]byte,1024)
		n,raddr,err:=conn.ReadFromUDP(buf)
		if err!=nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("从客户端%s接收的:%s",raddr,string(buf[:n]))
		conn.WriteToUDP([]byte(""),addr)
	}
}
