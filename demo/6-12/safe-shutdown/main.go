package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	router := gin.Default()
	listenAddr := fmt.Sprintf(":%d", 8888)
	server := &http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("http server start error", err)
		}
	}()

	//优雅关闭，原生版
	signals := make(chan os.Signal, 0)
	//监听退出信号
	//SIGHUP 1  终端控制进程结束(终端连接断开)
	//SIGINT 2 用户发送INTR字符(Ctrl+C)触发
	//SIGTERM 15 结束程序(可以被捕获、阻塞或忽略)
	//SIGQUIT 3 用户发送QUIT字符(Ctrl+/)触发
	//SIGKILL 9 | 无条件结束程序(不能被捕获、阻塞或忽略)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-signals
	log.Println(" receive system signal:", s)
	//依次关闭应用中的所有服务，（注意：需要根据服务的依赖顺序关闭）
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Println("http server error", err)
	}
}
