package filelisting

import (
	"net/http"
	"fmt"
	"os"
	"bufio"
	"strings"
)

const prefix="/list/"

type userError string

func (e userError)Error()string{
	return e.Message()
}

func (e userError)Message()string{
	return string(e)
}

func HindleFuncFile(w  http.ResponseWriter,r *http.Request)error{
	//当每次读取块的大小小于4KB，建议使用bufio.NewReader(f), 大于4KB用bufio.NewReaderSize(f,缓存大小)
	//要读Reader, 图方便用ioutil.ReadAll()
	//一次性读取文件，使用ioutil.ReadFile()
	path:=r.URL.Path
	fmt.Println(path)
	//判断url是否以prefix开头
	if strings.Index(path,prefix)!=0{
		return userError("必须开头是"+prefix)
	}
	file,err:=os.Open(path[len(prefix):])
	if err!=nil{
		return err
	}
	defer file.Close()

	//读取文件
	reader:=bufio.NewReader(file)
	//ioutil.ReadAll(file)
	buf:=make([]byte,4096)
	for {
		n,err:=reader.Read(buf)
		/*if err !=nil {
			if err==io.EOF{
				fmt.Println("send success")
			}else{
				fmt.Println("read err",err)
			}
			return
		}*/
		if n==0{
			fmt.Println("文件读取完毕")
			return err
		}
		if err!=nil{
			fmt.Println("文件读取错误 ",err)
			return err
		}
		w.Write(buf[:n])
	}

}
