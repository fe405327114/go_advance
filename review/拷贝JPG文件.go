package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"io"
)

func myerr(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		os.Exit(1)
	}
}
func Copyfile(path string) {
	f1, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	myerr(err, "openfile err")
	defer f1.Close()

	info, err := f1.Readdir(-1)

	myerr(err, "readdir err")
	var i int
	for _, data := range info {
		if strings.HasSuffix(data.Name(), ".JPG") {
			i++
			//打开JPG文件
			f2, err := os.Create("E:/学习资料/" + strconv.Itoa(i) + ".JPG")
			myerr(err, "create JPG err")

			f3, err := os.OpenFile(path+"/"+data.Name(), os.O_RDWR, 6)
			myerr(err, "open JPG err")
			buf := make([]byte, 4096)
			for {
				n, err := f3.Read(buf)
				if n == 0 {
					break
				}
				if err != io.EOF {
					myerr(err, "read JPG err")
				}
				f2.Write(buf[:n])
			}
			f2.Close()
			f3.Close()
		}
	}
}
func main() {
	var path string
	fmt.Println("pleasse inputthe path")
	fmt.Scan(&path)
	Copyfile(path)
}
