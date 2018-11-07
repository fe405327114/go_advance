package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

// 监听》上线》进入连接接口》获取用户信息存入map》写入全局通道》
//》创建Manager goroutine，将全局通道中的数据读取至用户通道
//》回到连接接口，将用户通道中的数据读取至conn
type Client struct {
	C    chan string
	Name string
	Addr string
}

//声明一个在线用户map
var onlineMap = make(map[string]Client)
//定义一个全局通道处理 聊天室广播消息
var message = make(chan string)

func main() {
	//  监听
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	//创建goroutine，处理消息
	go Manager()

	//判断是否连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err", err)
			continue
		}
		//defer在循环中使用要慎重，因为defer+函数调用可能会导致资源积压
		defer conn.Close()
		//一旦连接成功，调用处理接口的函数
		go HindleConn1(conn)
	}
}
func HindleConn1(conn net.Conn) {
	defer conn.Close()
	//获取conn接口中用户的地址
	netAdd := conn.RemoteAddr().String()
	//给新用户创建结构体
	clnt := Client{make(chan string), netAdd, netAdd}
	//将用户放至map中
	onlineMap[netAdd] = clnt

	//新创建一个goroutine专门给用户发送消息
	//从.C通道写入接口
	go SendMessage(clnt, conn)
	//把新用户上线的信息写入全局通道进行广播
	message <- clnt.Name + clnt.Addr + ":" + "login"

	//将用户退出的消息广播
	isQuit := make(chan bool)
	//超时处理
	hasData := make(chan bool)
	//将用户发的消息广播至所有用户
	go func() {
		buf := make([]byte, 2048)
		// 获取用户发送的消息内容
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("clnt exit", clnt.Name)
				return
			}
			if err != nil {
				fmt.Println("read err")
				return
			}
			//保存用户发来的消息,nc工具默认加"\n"
			msg := string(buf[:n-1])
			//如果用户发送了"who"，就显示所有用户信息
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("user list:\n"))
				for _, clnt := range onlineMap {
					//组织一下语言
					userinfo := clnt.Addr + ":" + clnt.Name + "\n"
					conn.Write([]byte(userinfo))
				}
			} else if len(msg) > 8 && msg[:6] == "rename" { //字符串可以直接使用[:]操作
				newname := strings.Split(msg, "|")[1] //按照"|"分割，reanme为[0],newname为[1]
				//把用户信息中的名字进行更新
				clnt.Name = newname
				//把map中的用户信息进行更新
				onlineMap[netAdd] = clnt
			} else {
				//将获取到的消息发送给每个用户
				//利用临时设计好的专门组织语言的函数比较方便， 此处选择直接拼接
				message <- string(clnt.Addr + clnt.Name + ":" + msg)
			}
			//通道中被写入数据说明用户有书要输入
			hasData <- true
		}
	}()
	//监听用户退出的消息
	for {
		select {
		case <-isQuit:
			delete(onlineMap, netAdd)
			message <- clnt.Addr + clnt.Name + ":" + "logout"
			return
		case <-hasData:

		case <-time.After(60 * time.Second):
			delete(onlineMap, netAdd)
			message <- clnt.Addr + clnt.Name + ":" + "timeout leave"
			return
		}
	}
}
func SendMessage(clnt Client, conn net.Conn) {
	for msg := range clnt.C {
		conn.Write([]byte(msg+"\n"))
	}
}

//从全局通道中读出数据，写入到.C里面，然后SendMessage将.C里面的
//数据写入到conn里面
func Manager() {
	for {
		//把全局通道中的数据读出来
		msg := <-message
		for _, clnt := range onlineMap {
			//把读出来的数据发送给所有用户
			clnt.C <- msg
		}
	}
}
