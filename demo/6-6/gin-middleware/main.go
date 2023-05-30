package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	//包含默认中间件（logger和recovery）的路由
	r1 := gin.Default()

	// 没有任何默认中间件的路由
	r2 := gin.New()

	// 全局中间件
	//即使GIN_MODE 设置为 release， Logger 中间件也会将日志写入 gin.DefaultWriter
	// 默认值 gin.DefaultWriter = os.Stdout
	//r2.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic，一旦请求过程中出现panic，http code会响应500
	//r2.Use(gin.Recovery())

	r1.GET("/test", func(context *gin.Context) {
		context.JSON(200, context.Request.RequestURI)
		panic("test1")
	})

	r2.GET("/test", func(context *gin.Context) {
		context.JSON(200, context.Request.RequestURI)
		panic("test2")
	})

	r2.GET("/test2", func(context *gin.Context) {
		context.JSON(200, context.Request.RequestURI)
	})
	//路由组中间件
	g1 := r1.Group("/api").Use(timeCost)
	g1.GET("/test", func(context *gin.Context) {
		time.Sleep(time.Second)
		context.JSON(200, context.Request.RequestURI)
	})
	//单个路由中间件
	//引入方式1 (注意：中间件函数需要放在前面)

	//r1.GET("test3", timeCost, func(context *gin.Context) {
	//	time.Sleep(time.Second)
	//	context.JSON(200, context.Request.RequestURI)
	//})

	//参数位置错误
	//c.Next控制下一个中间件的执行，这里将timeCost中间件放在第二个参数的位置，无法统计到第一个handler的耗时
	r1.GET("test3", func(context *gin.Context) {
		time.Sleep(time.Second)
		context.JSON(200, context.Request.RequestURI)
	}, timeCost)

	//引入方式2(注意：中间件函数需要放在前面)
	r1.Use(timeCost).GET("test4", func(context *gin.Context) {
		time.Sleep(time.Second)
		context.JSON(200, context.Request.RequestURI)
	})
	//参数位置错误
	r1.Use(timeCost).GET("test4", func(context *gin.Context) {
		time.Sleep(time.Second)
		context.JSON(200, context.Request.RequestURI)
	}).Use(timeCost)

	r1.Use(auth).GET("test5", func(context *gin.Context) {
		context.JSON(200, context.Request.RequestURI)
	})

	// 监听指定端口启动服务
	err := r1.Run(":8888")
	if err != nil {
		log.Println(err)
	}

	//err := r2.Run(":8889")
	//if err != nil {
	//	log.Println(err)
	//}
}

// 自定义中间件
func timeCost(c *gin.Context) {
	//请求处理前获取当前时间
	start := time.Now()
	//执行后续的handler
	c.Next()
	// 统计时间
	cost := time.Since(start).Milliseconds()
	fmt.Println("请求耗时：", cost)
}

func auth(c *gin.Context) {
	authStatus := false
	//TO DO auth
	if authStatus {
		c.Next()
		return
	}
	//验证不通过直接中断
	//c.AbortWithStatus(403)
	c.Abort()

}
