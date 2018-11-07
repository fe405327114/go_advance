package main

import (
	"fmt"
	"os"
	"strings"

	"io"
)

//指定目录拷贝特定文件：
//从用户给出的目录中，拷贝 .mp3文件到指定目录中。
func Copyy(path string) {
	f2, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("打开错误")
		return
	}
	defer f2.Close()
	f3, _ := os.Create("E:/学习资料/abc.JPG")
	defer f3.Close()

	info, _ := f2.Readdir(-1)
	for _, data := range info {
		if strings.HasSuffix(data.Name(), "jpg") {
			fmt.Println(data.Name())
			//f2打开的是目录指针，不可以直接读取复制
			//要再打开一次JPG文件
			f4,err:=os.OpenFile(path+"/"+data.Name(),os.O_RDWR,6)
			if err!=nil{
				fmt.Println("Open JPG err")
				return
			}
			for {
				buf := make([]byte, 1024)
				n, err := f4.Read(buf)
				if err == io.EOF {
					fmt.Println("读取完毕")
					return
				}
				f3.Write(buf[:n])
			}
			f4.Close()
		}
	}
}
func main() {
	fmt.Println("请输入目录")
	var path string
	fmt.Scan(&path)
	Copyy(path)
}
