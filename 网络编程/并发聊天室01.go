package main

import (
	"net"
	"fmt"
	"time"
)

type Client1 struct {
	C     chan string
	Name  string
	Addrr string
}

//全局变量不可以用自动推导
var message1 = make(chan string)
var m1 = make(map[string]Client1)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("listen err")
		return
	}
	defer listener.Close()
	//向所有用户广播
	go MessageToC()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("conn err")
			return
		}
		go HindleConn3(conn)
	}
}
func HindleConn3(conn net.Conn) {
	defer conn.Close()
	addrr := conn.RemoteAddr().String()
	clnt := Client1{make(chan string), addrr, addrr}
	m1[addrr] = clnt
	//专门给单个客户发送消息的协程
	go CToConn(conn, clnt)

	message1 <- clnt.Name + clnt.Addrr + ":" + "login"

	//监听是否退出
	quit := make(chan bool)
	//监听是否超时
	Stay := make(chan bool)

	//专门向message中写数据
	go WriteInMessage(conn,clnt,quit,Stay)
	for {
		select {
		case <-quit:
			delete(m1, clnt.Addrr)
			message1 <- clnt.Addrr + clnt.Name + "logout"
			return
		case <-Stay: //让计时器归零

		case <-time.After(60 * time.Second):
			delete(m1, clnt.Addrr)
			message1 <- clnt.Addrr + clnt.Name + ":" + "timeout leave"
			return
		}
	}
}
func CToConn(conn net.Conn, clnt Client1) {
	for data := range clnt.C {
		conn.Write([]byte(data+"\n"))
	}
}
func MessageToC() {
	//将需要发送的信息从message中取出发送至每个用户的通道，
	//那么SendToClient函数会负责将每个用户通道中的信息写入接口
	for {
		data := <-message1
		for _, clnt := range m1 {
			clnt.C <- data
		}
	}
}
func WriteInMessage(conn net.Conn, clnt Client1,quit chan bool,Stay chan bool) {
	buf := make([]byte, 4096)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			quit <- true
			break
		}
		msg := string(buf[:n-1])
		if msg == "who" && len(msg) == 3 {
			conn.Write([]byte("users list:\n"))
			for _, clnt := range m1 {
				info := clnt.Addrr + clnt.Name + "\n"
				conn.Write([]byte(info))
			}
		} else if len(msg) > 8 && msg[:6] == "rename" { //字符串可以直接使用[:]操作
			//newname := strings.Split(msg, "|")[1] //按照"|"分割，reanme为[0],newname为[1]
			newname:=msg[8:]
			//把用户信息中的名字进行更新
			clnt.Name = newname
			//把map中的用户信息进行更新
			m1[clnt.Addrr] = clnt
			conn.Write([]byte("rename success"))
		} else {
			message1 <- clnt.Addrr + clnt.Name + ":" + msg
		}
		//说明用户有数据要输入
		Stay <- true
	}
}