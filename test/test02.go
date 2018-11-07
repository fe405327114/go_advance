package main

import (
	"net"
	"fmt"
	"os"
	"strings"
)

//编一个TCP并发服务器，可同时支持N个客户端访问。服务器接收客户端发送内容，将内容按单词逆置，回发给客户端。
// 如： 客户端发送：this is a socket test 服务器回复：test socket a is this
func myerr1(err error,info string){
	if err!=nil{
		fmt.Println(err,info)
		os.Exit(1)
	}
}

func HandleData(conn net.Conn){
	defer conn.Close()
	//读取客户端发送内容
	var result string
	buf:=make([]byte,4096)
	for{
		n,err:=conn.Read(buf)
		myerr1(err,"read conn err")
		if n==0{
			break
		}
		result=string(buf[:n])
		//将内容按单词逆置
		slice:=SortData(result)
		//发送给客户端
		resultData:=strings.Join(slice," ")
		conn.Write([]byte(resultData))
	}
}
func SortData(result string)(slice  []string){
	slice=strings.Split(result," ")
	n:=len(slice)
	for i:=0;i<len(slice);i++{
		if i==n{
			break
		}
		n--
		slice[i]=slice[n]
	}
	return
}
func main(){
	//创建监听SOCKET
	listener,err:=net.Listen("tcp","127.0.0.1:8008")
	myerr1(err,"listener err")
	defer listener.Close()
	//循环监听，建立连接
	 for{
	 	conn,err:=listener.Accept()
	 	myerr1(err,"accept err")
	 	//进入处理连接协程
	 	go HandleData(conn)
	 }
}