package main

import (
	"flag"
	"os"
	"fmt"
	"bufio"
	"io"
	"time"
	"strings"
	"github.com/garyburd/redigo/redis"
)

type cmdParams struct {
	logFilePath string
	routineNum  int
}

//创建日志统计结构体(js上报)
type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data digData
	uid  string
}

type urlNode struct {
	//
}

//用于存储的结构体
type storageBlock struct {
	counterType  string //类型
	storageModel string //存储格式
	unode        urlNode
}

//初始化打日志库
var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main( ) {
	//获取参数
	logFilePath := flag.String("logFilePath", "/*", "log")

	routineNum := flag.Int("routineNum", 5, "consumer num")

	l := flag.String("l", "/tmp/log", "log")
	flag.Parse()
	params := cmdParams{*logFilePath, *routineNum}
	//打日志
	lodFd, err := os.OpenFile(*l, os.O_RDWR, 6)
	if err != nil {
		log.Out = logFd
		defer lodFd.Close()
	}

	//初始化一些channel,用于数据传递
	var logChannel = make(chan string, routineNum)
	var pvChannel = make(chan urlData, routineNum)
	var uvChannel = make(chan urlData, routineNum)
	var storageChannel = make(chan string, routineNum)

	//日志消费者
	go readFileLine(params, logChannel)

	//创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}
	//创建PV  UV统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)

	//创建存储器
	go dataStorage(storageChannel)

}

func readFileLine(params cmdParams, logChannel chan string) {
 fd,err:=os.Open(params.logFilePath)
 defer fd.Close()
 if err!=nil{
 	fmt.Println("open err",err)
 	return

 }

	bufferRead:=bufio.NewReader(fd)
	for{
		line,err:=bufferRead.ReadString('\n')
		logChannel<-line
		if err!=nil{
			if err==io.EOF{
				time.Sleep(3*time.Second)
			}else{
				fmt.Println("read err",err)
			}
		}



	}

}
func logConsumer(logChannel chan string, pvChannel chan urlData, uvChannel chan urlData) {
   for logStr:= range logChannel{
   	data:=cutLogData(logStr)

   }
}
func cutLogData(logStr string)digData{
	//去除空格
	logStr=strings.TrimSpace(logStr)
	post1:=strings.Index(logStr,HANDLE_DIG)
	if post1==-1{
		return digData{}
	}
	post1+=len(HANDLE_DIG)
	post2:=strings.Index(logStr,"HTTP/")
	//截取字符串
	rs:=[]rune(logStr)
	d:=string(rs[post1:post2])

}

func pvCounter(pvChannel chan urlData, storageChannel chan string) {

}
func uvCounter(uvChannel chan urlData, storageChannel chan string) {

}
func dataStorage(storageChannel chan string) {

}
