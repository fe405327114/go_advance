package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	start:=time.Now()
	ch:=make(chan string)
	for _,url:=range os.Args[1:]{//  os系统包，Args是[]string类型，遍历返回值为n和value
		go fetch(url,ch)
	}
	for range os.Args{
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f total",time.Since(start).Seconds())
}
func fetch(url string,ch chan string){
	start:=time.Now()
	resp,err:=http.Get(url)
	if err!=nil{
		ch<-fmt.Sprintf("get %s err,%s",url,err)
		return
	}
	defer resp.Body.Close()
	//读取整个响应数据流的缓冲区，放至b[]byte
	b,err1:=ioutil.ReadAll(resp.Body)
	if err1!=nil{
		ch<-fmt.Sprintf("read %s err,%s",url,err1)
		return
	}
	s:=time.Since(start).Seconds()
	ch<-fmt.Sprintf("%.2fs,%s,%s",s,string(b),url)

}