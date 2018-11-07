package main

import (
	"net/http"
	"fmt"
	"strconv"
	"os"
)

func main() {
	var start, end int

	fmt.Println("请输入要爬取的起始页面")
	fmt.Scan(&start)
	fmt.Println("请输入要爬取的结束页面")
	fmt.Scan(&end)
	Spider(start, end)

}
func Spider(start, end int) {
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go Spidergo(i, ch)
	}
	for j := start; j <= end; j++ {
		fmt.Printf("第%d个页面完成\n", <-ch)
	}
}
func Spidergo(i int, ch chan int) {
	fmt.Printf("正在爬取第%d个页面\n", i)
	resp, err := http.Get("http://tieba.baidu.com/f?kw=%E6%B2%81%E9%98%B3&ie=utf-8&pn=" + strconv.Itoa((i-1)*50))
	if err != nil {
		fmt.Println("Resp err")
		return
	}
	defer resp.Body.Close()

	result:=HindleData(resp)
	f, err := os.Create(strconv.Itoa(i) + ".html")
	if err != nil {
		fmt.Println("Create err")
		return
	}
	f.WriteString(result)
	f.Close()
	ch <- i
}
func HindleData(resp *http.Response) (result string) {
	buf := make([]byte, 4096)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}
