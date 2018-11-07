package main

import (
	"net/http"
   "case/filelisting"
    "log"
	"os"
)

type appHanler func(w http.ResponseWriter,r *http.Request)error

func errWrapper(handler appHanler) func(http.ResponseWriter,*http.Request){
return func(w http.ResponseWriter,r *http.Request){

	defer func() {
		if rv:=recover();rv!= nil{
			log.Printf("panic recover ",rv)
			http.Error(w,http.StatusText(http.StatusNotFound),http.StatusNotFound)
		}
	}()
     err:=handler(w,r)
     if err!=nil{
     	log.Printf("err request: %s",err.Error())
//类型断言，判断是否为userError
     	if userError,ok:=err.(userError);ok{
     		http.Error(w,userError.Message(),http.StatusBadRequest)
     		return
		}

     	code:=http.StatusOK
		 switch {
		 case os.IsNotExist(err):
		 	//未找到页面
		 	code=http.StatusNotFound
		 	//statustext包装错误信息
		 	http.Error(w,http.StatusText(code),code)
		 case os.IsPermission(err):
		 	//无权限
			 code=http.StatusForbidden
			 http.Error(w,http.StatusText(code),code)
		 default:
		 	//服务器响应失败
			 code=http.StatusInternalServerError
			 http.Error(w,http.StatusText(code),code)
		 }
	 }
}
}

//自定义错误类型，给用户看
type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/",errWrapper(filelisting.HindleFuncFile))

	err:=http.ListenAndServe("127.0.0.1:8899",nil)
	if err!=nil{
		panic(err)
	}
}
