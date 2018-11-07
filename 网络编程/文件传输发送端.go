package main

import (
	"os"
	"fmt"
	"net"
	"io"
)

func main() {
	//list := os.Args
	//if len(list) != 2 {
	//	fmt.Println("geshi err")
	//	return
	//}
	//path := list[1]
	//获取文件名字
	path:="E:/学习资料/证件照.JPG"
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("stat err", err)
		return
	}
	fileName := fileInfo.Name()
	//建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("dial err")
		return
	}
	defer conn.Close()
	//发送文件名字
	_, err1 := conn.Write([]byte(fileName))
	if err1 != nil {
		fmt.Println("write err", err1)
		return
	}
	//接收返回的数据
	buf := make([]byte, 4096)
	n, _ := conn.Read(buf)
	if n == 0 {
		return
	}
	if string(buf[:n]) == "ok" {
		sendFile(conn, path)
	}
}
func sendFile(conn net.Conn, path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Open err", err)
		return
	}
	defer f.Close()
	//读取文件
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err !=nil {
			if err==io.EOF{
				fmt.Println("send success")
			}else{
				fmt.Println("read err",err)
			}
			return
		}
		//写入接口
		conn.Write(buf[:n])
	}
}
