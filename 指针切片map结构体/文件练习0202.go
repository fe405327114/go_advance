package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io"
)

//统计指定目录内单词出现次数：
//统计指定目录下，所有.txt文件中，“Love”这个单词 出现的次数。
//1,打开目录，找到.txt
//2,按行读取，累加次数

func LoveCount(path string)int {
	//打开txt文件
	ftxt, err := os.OpenFile(path, os.O_RDWR, 6)
	if err != nil {
		fmt.Println("Open txt err")
		return 0
	}
	defer ftxt.Close()
	//创建reader,按行读取
	reader := bufio.NewReader(ftxt)
	var sum int
	for {
		buf, err := reader.ReadBytes('\n')
		// 此处需要将buf传给另外一个函数，不然EOF会跳出整个循环，包括后面的程序
		sum += LoveCount1(buf)
		//按行读取得先进行操作，再判断是否读到EOF
		//不然会少读一行
		if err == io.EOF {
			break
		}
		fmt.Println("0000000000")

	}
	return sum
}
	//这里返回值为buf(字符切片),err
	//注意与直接创建缓冲区的区分
	// buf:=make([]byte,1024)
	//n,err:=ftxt.read(buf) 将ftxt读取至buf中

	func LoveCount1(buf []byte) int{
		s1 := strings.Fields(string(buf[:]))
		m1 := make(map[string]int)
		for i := 0; i < len(s1); i++{
		m1[s1[i]]++
		fmt.Println("=========")
	}
		for k, _ := range m1{
		if k == "fd"{
		return m1[k]
	}
	}
	return 0
}
func main() {
	fmt.Println("请输入路径")
	var path string
	fmt.Scan(&path)
	//打开指定目录
	f1, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err")
		return
	}
	defer f1.Close()
	//遍历目录，获取接口类型的切片
	info, _ := f1.Readdir(-1)
	// 遍历切片，获取.txt文件
	count := 0
	for _, data := range info {
		if strings.HasSuffix(data.Name(), "txt") {
			count+=LoveCount(path+"/"+data.Name())
		}
	}
	fmt.Println(count)
}
