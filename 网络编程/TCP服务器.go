package 网络编程

import (
	"net"
	"fmt"
)

func main() {
	//net.Listen(协议，地址),返回值为Listener接口
	//Listener接口包括 Accept,Close,Addr 三个方法,其中Accept返回值为Conn接口和err
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	fmt.Println("等待客户端建立连接")
//Conn接口 type Conn interface {
	//	Read(b []byte) (n int, err error)
	//	Write(b []byte) (n int, err error)
	//	Close() error
	//	LocalAddr() Addr
	//	RemoteAddr() Addr
	//	SetDeadline(t time.Time) error
	//	SetReadDeadline(t time.Time) error
	//	SetWriteDeadline(t time.Time) error
	//}
	conn,err:=listener.Accept()
	if err!=nil {
		fmt.Println("conn err",err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功")
	buf:=make([]byte,1024)
	n,err:=conn.Read(buf)
	if err!=nil{
		fmt.Println("read err",err)
	}
	return
	fmt.Println("服务器读到",string(buf[:n]))

}
