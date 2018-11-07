package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
)

func main() {
	var start, end int
	fmt.Println("请输入起始页面")
	fmt.Scan(&start)
	fmt.Println("请输入结束页面")
	fmt.Scan(&end)
	Dowork(start,end)
}
func Dowork(start, end int) {
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go Spider3(i, ch)
	}
	for j := start; j <= end; j++ {
		fmt.Printf("第%d个爬取完毕\n", <-ch)
	}
}
func Spider3(i int, ch chan int) {
	resp, err := http.Get("http://tieba.baidu.com/f?kw=%E6%B2%81%E9%98%B3&ie=utf-8&pn=" + strconv.Itoa((i-1)*50))
	if err != nil {
		fmt.Println("resp err", err)
		return
	}
	result := DataHindle(resp)
	f, err := os.Create("E:/学习资料/" + strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("creste err")
		return
	}
	f.WriteString(result)
	f.Close()
	ch <- i
}
func DataHindle(resp *http.Response) (result string) {
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		//re:=regexp.MustCompile()//正则表达式
		//re.FindAllStringSubmatch()
		result += string(buf[:n])
	}
	return
}
