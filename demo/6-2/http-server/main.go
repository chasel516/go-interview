package main

import (
	"fmt"
	"net/http"
)

func main() {

	//1.创建http server对象
	mux := http.NewServeMux()

	//2.注册路由和中间件
	mux.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(resp, "root")
	})
	mux.HandleFunc("/test", func(resp http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(resp, "test")
	})
	//3.启动http服务器，传入监听地址和多路复用器
	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
	}

	//http.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(resp, "root")
	//})
	//
	//http.HandleFunc("/test", func(resp http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(resp, "test")
	//})
	//
	//err := http.ListenAndServe(":8888", nil)
	//if err != nil {
	//	fmt.Println(err)
	//}

}
