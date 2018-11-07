package main

import (
	"net/http"
	"fmt"
)

func myHandler(w http.ResponseWriter,r *http.Request){
	//还有两个方法Header Header类型和 WriteHeader（int）
	//参数2为Request结构体类型指针，用于处理客户端请求(读取)
	fmt.Println(r.Method)
	fmt.Println(r.URL)
	fmt.Println(r.Header)
	fmt.Println(r.Body)
	fmt.Println(r.RemoteAddr,"connect success")

	//参数1为ResponseWriter接口类型，用于给客户端回复数据（写入）
	w.Write([]byte("hello http"))

}
func main(){
	// http.HandleFunc，参数1为pattern（模式），类型为字符串
	//参数2为一个函数，此函数的格式的规定有两个参数
 	http.HandleFunc("/hello",myHandler)//注意有个"/"

 	// http.ListenAndServe第一个参数addr是监听地址
 	//第二个参数是Handle类型的handle即服务端处理程序，通常为nil
 	//注意不要加http，客户端请求时候要加http
 	http.ListenAndServe("127.0.0.1:8000",nil)
 }