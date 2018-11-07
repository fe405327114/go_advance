package main

import (
	"net"
	"fmt"
	"time"
	"strings"
)

type Client2 struct {
	C    chan string
	name string
	addr string
}

var Message2 = make(chan string)
var m2 map[string]Client2

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8009")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()

	go Manager1()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Listen err", err)
			return
		}
		go ConnHindle1(conn)
	}
}
func ConnHindle1(conn net.Conn) {
	defer conn.Close()
	//将用户上线的消息发送到Message中
	addr := conn.RemoteAddr().String()
	clnt := Client2{make(chan string), addr, addr}
	m2[clnt.name] = clnt
	//专门给用户发消息的goroutine()写入conn
	go SendtoConn(conn, clnt)
	Message2 <- clnt.addr + clnt.name + ":" + "login"
	//超时处理
	hasData := make(chan bool)
	//用户退出处理
	quit := make(chan bool)
	//专门向message中写数据的goroutine
	go writeinMessage1(conn, clnt, quit, hasData)
	for {
		select {
		case <-quit:
			close(clnt.C)
			delete(m2, clnt.addr)
			Message2 <- clnt.addr + clnt.name + "logout"
			return
		case <-hasData:

		case <-time.After(time.Second * 500):
			delete(m2, clnt.addr)
			Message2 <- clnt.addr + clnt.name + "time out"
			return
		}
	}
}

//从C中取出，放至conn中
func SendtoConn(conn net.Conn, clnt Client2) {
	for data := range clnt.C {
		conn.Write([]byte(data + "\n"))
	}
}
func writeinMessage1(conn net.Conn, clnt Client2, quit chan bool, hasData chan bool) {
	//将用户发送的消息进行广播
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			quit <- true
			return
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		//显示展现用户列表
		msg := string(buf[:n-1])
		if "who" == msg && len(msg) == 3 {
			for _, clnt := range m2 {
				userinfo := clnt.addr + clnt.name + "\n"
				conn.Write([]byte(userinfo))
			}
		} else if len(msg) > 8 && msg[:6] == "rename" {
			newname := msg[7:n-1]
			var flag bool
			for _, clnt := range m2 {
				if clnt.name == newname {
					flag = true
					fmt.Println("the same name")
					break
				}
			}
			if flag == false {
				clnt.name = newname
				m2[clnt.addr] = clnt
				conn.Write([]byte("rename success" + "\n"))
			}
		} else if len(msg) > 3 && msg[:3] == "one" {
			onename := strings.Split(msg,"@")[1]
			for _,clnt:=range m2{
				if clnt.name==onename{
					m2[clnt.addr].C <-msg+"\n"
				}
			}
		} else {
			Message2 <- clnt.addr + clnt.name + ":" + msg
		}
		hasData <- true
	}
}

//专门 将message中取出放置C中的goroutine
func Manager1() {
	m2 = make(map[string]Client2)
	for {
		data := <-Message2
		for _, clnt := range m2 {
			clnt.C <- data
		}
	}
}
