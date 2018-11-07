package main

import (
	"fmt"
	"os"
	"strings"
)
//指定目录检索特定文件：
//		从用户给出的目录中，找出所有的 .jpg 文件。
func main() {
	fmt.Println("请输入目录")
	var path string
	fmt.Scan(&path)
	f1, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err!=nil{
		fmt.Println("OpenFile err")
		return
	}
	defer f1.Close()
	info,_:=f1.Readdir(-1)//和操作文件比不用创建缓冲区
	//返回值为目录信息的切片(接口类型的切片)
	for _,data:=range info{
		if data.IsDir(){
			fmt.Println(data.Name(),"这是目录")
		}else{
			fmt.Println(data.Name(),"这是文件")
		}
		 if strings.HasSuffix(data.Name(), "JPG"){//返回值为布尔类型
			 fmt.Println(data.Name())
		 }
	}
}
