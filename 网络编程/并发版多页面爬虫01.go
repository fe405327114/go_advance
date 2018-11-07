package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
)

func main() {
	var start, end int
	fmt.Println("请输入要爬取的起始页面")
	fmt.Scan(&start)
	fmt.Println("请输入要爬取的终止页面")
	fmt.Scan(&end)
	Spider1(start, end)
}
func Spider1(start, end int) {
	cha:=make(chan int)
	for i := start; i <= end; i++ {
	go Spider2(i,cha)
	}
	for j:=start;j<=end;j++{
		fmt.Printf("第%d个页面爬取完成\n",<-cha)
	}
}
func Spider2(i int,cha chan int){
	resp, err := http.Get("http://tieba.baidu.com/f?kw=%E6%B2%81%E9%98%B3&ie=utf-8&pn=" + strconv.Itoa((i-1)*50))
	if err!=nil{
		fmt.Println("resp err",err)
		return
	}
	defer resp.Body.Close()
	result:=HindleData1(resp)
	f,err:=os.Create("E:/学习资料/"+strconv.Itoa(i)+".html")
	if err!=nil{
		fmt.Println("Create err")
		return
	}
	f.WriteString(result)
	f.Close()
	cha<-i
}
func HindleData1(resp *http.Response)(result string){
	buf:=make([]byte,4096)
	for{
		n,_:=resp.Body.Read(buf)
		if n==0{
			break
		}
		result+=string(buf[:n])
	}
	return
}
