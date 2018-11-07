package main

import (
	"fmt"
	"strings"
	"io"
	"bufio"
)

func fibonacci() inGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type inGen func() int
//给函数实现接口
func (g inGen) Read(p []byte) (n int, err error) {
	//取到斐波那契值
	next:=g()
	if next>10000{
		return 0,io.EOF
	}
	//nexts := strconv.Itoa(next)
	//Sprintf将格式化的字符串写入某个字符串缓冲区，可以将其他格式转换成字符串
	//但是有可能发生缓冲区溢出
	nexts:=fmt.Sprintf("%d\n",next)
	//将值读取放置p中
	return strings.NewReader(nexts).Read(p)
}

//定义扫描器,需要传入一个实现了reader接口的参数，上面inGen实现了reader
func printScanner(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan(){
      fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	printScanner(f)


}
