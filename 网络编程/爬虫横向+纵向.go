package main

import (
	"fmt"
	"os"
	"net/http"
	"strconv"
	"io"
	"regexp"
	"strings"
)

func myerr3(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}
func spiderdata1(resp *http.Response) (result string) {
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err != io.EOF {
			myerr3(err, "resp err")
		}
		result += string(buf[:n])
	}
	return
}
func spiderdata2(resp *http.Response) (result2 string) {
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err != io.EOF {
			myerr3(err, "read 2 err")
		}
		result2 += string(buf[:n])
	}
	return
}
func HindleResult2(result2 string)(title,content string) {
	//pix:=`<div class='bdsharebuttonbox clearfix social_group' title="(?s:(.*?))">`
	//pixw:=`<div class='bdsharebuttonbox clearfix social_group' title="(?s:(.*?))" humorId`
	re1 := regexp.MustCompile(`<h1>(.*?)</h1>`)//`<h1>(.*?)</h1>`
	alls1:= re1.FindAllStringSubmatch(result2, -1)
	//fmt.Println(alls1)
	for _, TargetTitle := range alls1{
		title=TargetTitle[1]
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	alls2 := re2.FindAllStringSubmatch(result2, 1)
	for _, TargetContent := range alls2 {
		content=TargetContent[1]
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "&nbsp; ", "", -1)
		break
	}
	return
}
func ColSpider(TargetURL string)(title,content string) {
	fmt.Println(TargetURL)
	resp, err := http.Get(TargetURL)
	myerr3(err, "resp err")
	result2 := spiderdata2(resp)
	title,content=HindleResult2(result2)
	return
}
var f1 =make([]string,0)
var f2 =make([]string,0)
func HindleResult1(result string)  {
	re1 := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	TemURL := re1.FindAllStringSubmatch(result, -1)
	for _, TargetURL := range TemURL {
		title,content:=ColSpider(TargetURL[1])
		f1=append(f1,title)
		f2=append(f2,content)
	}
	return
}

func CreateResult1( i int,title,content []string) {
	f, err := os.Create("E:/学习资料/" + strconv.Itoa(i) + ".txt")
	myerr3(err, "create err")
	for j:=0;j<len(f2);j++{
		f.WriteString(f1[j]+"\r\n"+f2[j]+"\r\n")
		f.WriteString("-----------------------------------\r\n")
	}

	f.Close()

}


func RowSpider(i int, ch chan int) {
	resp, err := http.Get("https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html")
	myerr3(err, "resp err")
	result1 := spiderdata1(resp)
	HindleResult1(result1)

	if len(f1)==10{
		CreateResult1(i,f1,f2)
		ch <- i
	}
}
func workstart(start, end int) {
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go RowSpider(i, ch)
}
	for j := start; j <= end; j++ {
		fmt.Printf("%d  finished\n", <-ch)
	}
}
func main() {
	var start, end int
	fmt.Println("scan start")
	fmt.Scan(&start)
	fmt.Println("scan end")
	fmt.Scan(&end)
	workstart(start, end)
}
