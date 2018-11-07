package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"io"
	"time"
)

func myerr(err error,info string){
if err!=nil && err!=io.EOF{
	fmt.Println(info,err)
	os.Exit(1)
}
}
func dataresult(resp *http.Response)(result string){
	 defer resp.Body.Close()
	buf:=make([]byte,4096)
	for{
		n,err:=resp.Body.Read(buf)
		if n==0{
			return
		}
		myerr(err,"read err")
		result+=string(buf[:n])
	}
	return
}
func Spider4(i int,ch chan int) {
	time.Sleep(time.Second)
	resp, err := http.Get("http://tieba.baidu.com/f?kw=%E7%94%B5%E5%BD%B1&ie=utf-8&pn=50" + strconv.Itoa((i-1)*50))
	myerr(err,"resp err")
	result:=dataresult(resp)
	f,err:=os.Create("E:/学习资料/"+strconv.Itoa(i)+".html")
	myerr(err,"create err")
	f.WriteString(result)
	f.Close()
	ch<-i
}
func gostart(start int, end int) {
	ch:=make(chan int)
	for i := start; i <= end; i++ {
		go Spider4(i,ch)
	}
	for j:=start;j<=end;j++{
		fmt.Printf("spidering %d page finish\n",<-ch)
	}
}
func main() {
	var start, end int
	fmt.Println("scan start")
	fmt.Scan(&start)
	fmt.Println("scan end")
	fmt.Scan(&end)
	gostart(start, end)
}
