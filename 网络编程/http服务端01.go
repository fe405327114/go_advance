package main

import (
	"net/http"
	"fmt"
	"os"
)


func myhindle(w http.ResponseWriter,r *http.Request){

	f,err:=os.Open("E:/学习资料"+r.URL.String())
	if err!=nil{
		w.Write([]byte("not fount"))
		return
	}
	defer f.Close()
	buf:=make([]byte,4096)
	for{
		n,_:=f.Read(buf)
		if n==0{
			break
		}
		w.Write([]byte(buf[:n]))
	}
	fmt.Println(r.Header)
	fmt.Println(r.Method)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Host)
	fmt.Println(r.URL)
	fmt.Println(r.Body)
}
func main (){
	http.HandleFunc("/",myhindle)
	http.ListenAndServe("127.0.0.1:8000",nil)
}
