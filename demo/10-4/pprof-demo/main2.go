package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建 Gin 引擎
	router := gin.Default()

	// 注册 pprof 路由
	router.GET("/debug/pprof/*action", gin.WrapF(http.DefaultServeMux.ServeHTTP))

	// 注册其他路由
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	// 启动服务
	router.Run(":8080")
}
