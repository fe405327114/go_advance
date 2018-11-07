package main

import (
	"fmt"
	"net/http"
	"strconv"
	"os"
	"io"
	"regexp"
)

func myerr1(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}
func Cdata(resp *http.Response) (result string) {
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err != io.EOF {
			myerr1(err, "read err")
		}
		result += string(buf[:n])
	}
	return
}
func Cresult(result string,ch2  chan [][]string,) {
	re1 := regexp.MustCompile(`<img width="100" alt="(.*?)"`)
	anthor := re1.FindAllStringSubmatch(result, -1)

	re2 := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	score:= re2.FindAllStringSubmatch(result, -1)

	ret3 := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	peopleNum := ret3.FindAllStringSubmatch(result, -1)
	ch2<- anthor
	ch2<-score
	ch2<-peopleNum
}
func Ccreate(i int, anthor, score,peopleNum [][]string){
	f, err := os.Create("E:/学习资料/" + strconv.Itoa(i) + ".txt")
	myerr1(err, "create err")
	n := len(anthor)
	f.WriteString("电影" + "---------" + "评分"+"-------"+"人数"+"\r\n")
	for i := 0; i < n; i++ {
		f.WriteString(anthor[i][1] + "---------" + score[i][1]+"-------"+peopleNum[i][1]+"\r\n")
	}
	f.Close()
}
func Spider5(i int, ch chan int) {
	ch2:=make(chan [][]string)
	resp, err := http.Get("https://movie.douban.com/top250?start="+ strconv.Itoa((i-1)*25) + "&filter=")
	defer resp.Body.Close()
	myerr1(err, "resp err")
	result :=Cdata(resp)

	go Cresult(result,ch2)
	 anthor:=<-ch2
	score:=<-ch2
	peopleNum:=<-ch2
	Ccreate(i,anthor,score,peopleNum)
	ch <- i
}
func startwork(start int, end int) {
	ch := make(chan int)
	for i := start; i <= end; i++ {
		go Spider5(i, ch)
	}
	for j := start; j <= end; j++ {
		fmt.Printf("%d page finish\n", <-ch)
	}
}
func main() {
	var start, end int
	fmt.Println("scan start")
	fmt.Scan(&start)
	fmt.Println("scan end")
	fmt.Scan(&end)
	startwork(start, end)
}
