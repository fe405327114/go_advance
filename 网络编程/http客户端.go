package main

import (
	"net/http"
	"fmt"
)

//http.Get（url string ）
// 返回值为（resp *Response,err error）
//服务器发送的响应包体保存至Response
// 是一个结构体
func main() {
	resp, err := http.Get("http://127.0.0.1:8000/hello")
	if err != nil {
		fmt.Println("Get err")
		return
	}
	//记得将包体关闭
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	//读取服务器响应的数据
	buf:=make([]byte,1024*4)
	var result string
	for{
		n,_:=resp.Body.Read(buf)
		if n==0{
			break
		}
		result+=string(buf[:n])
	}
	fmt.Println(result)
}
