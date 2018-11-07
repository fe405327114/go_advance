package main

import (
	"os"
	"fmt"
	"strings"
	"io"
	"strconv"
)

//将指定 C:/a/ 目录下的所有 MP3文件提取，重命名为 1xx.mp3，2xx.mp3，3xx.mp3 ..... 拷贝到 C:/b/ 目录下。
func myerr(err error,info string){
	if err!=nil{
		fmt.Println(err,info)
		os.Exit(1)
	}
}
func main(){
	//打开目录
	path:="C:/a/"
	f1,err:=os.OpenFile(path,os.O_RDONLY,os.ModeDir)
	defer f1.Close()
	myerr(err,"open dir err")
	//读目录信息
	DirInfo,err:=f1.Readdir(-1)
	myerr(err,"read dir err")
	//遍历目录信息，找到MP3文件
	var i int//新名称标志数
	for _,data:=range DirInfo{
		if strings.HasSuffix(data.Name(),".MP3"){
			i++//新名称标志数
			fmt.Printf("正在复制第%d个文件",i)
			//打开源文件
			f2,err:=os.OpenFile(path+data.Name(),os.O_RDWR,6)
			myerr(err,"open mp3 file err")
			//创建新文件
			newname:=strings.Split(data.Name(),".")[0]
			f3,err:=os.Create("C:/b/"+strconv.Itoa(i)+newname+".mp3")
			myerr(err,"create mp3 file err")

			//读取源文件，写入新文件
			buf:=make([]byte,4096)
			for{
				n,err:=f2.Read(buf)
				if n==0{
					break
				}
				if err!=io.EOF{
					myerr(err,"read mp3 err")
					return
				}
				f3.Write(buf[:n])
			}
			fmt.Printf("第%d个文件复制完成",i)
			f2.Close()
			f3.Close()
		}
	}
}
