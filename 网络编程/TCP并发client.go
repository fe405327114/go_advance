package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	//客户端发起请求，返回conn接口
	conn,err:=net.Dial("tcp","127.0.0.1:8001")
	if err!= nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go func(){
		//利用子协程处理用户键盘输入的内容，用os.Stdin.包获取
		s:=make([]byte,1024)
	for {
		n,err:=os.Stdin.Read(s)
		if err!=nil{
			fmt.Println(err)
			return
		}
		//将读到的数据，写入，发送到服务器
		_,err=conn.Write(s[:n])
		if err!=nil{
			fmt.Println(err)
			return
		}
	}
	}()
	//从服务器接收数据，并打印至屏幕
	buf:=make([]byte,1024)
	for  {
		n,err:=conn.Read(buf)
		if err!=nil{
			return //break?
		}

		fmt.Printf("接收从服务器发回的数据：%s",string(buf[:n]))
	}
}