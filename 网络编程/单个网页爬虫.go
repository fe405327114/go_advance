package main

import (
	"net/http"
	"fmt"
	"os"
)

func main(){
	resp,err:=http.Get("https://www.huxiu.com/")
	if err!=nil{
		fmt.Println("resp err")
		return
	}
	defer resp.Body.Close()
	buf:=make([]byte,4096)
	 var result  string
	 for{
		n,_:=resp.Body.Read(buf)
		if n==0{
			break
		}
		result+=string(buf[:n])
	}
	f,err:=os.Create("E:/学习资料/huxiu.txt")
	if err!=nil{
		fmt.Println("Create err",err)
		return
	}
	f.WriteString(result)
	f.Close()

}

