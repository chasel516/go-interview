package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//注意，在使用gin加载模板时，debug模式与release模式有很大的区别
	//debug模式下每个请求都会重新加载模板文件，所以在debug模式下修改模板文件强刷页面后可以看到页面内容会实时变更
	//而release模式下只加载一次模板文件，所以在高并发场景下这两个模式的性能表现会有很大差距。
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	//设置静态资源文件访问路径
	//第一个参数是映射后http访问时相对于站点域名域名或ip的路径
	//第二个参数是静态文件在服务器上的真实路径，需要设置一个文件夹的路径（映射文件夹到路由）
	router.Static("/assets", "./static/dist/assets")
	//参数类型不同
	//router.StaticFS("/assets", http.Dir("./static/dist/assets"))
	//映射文件到路由（需要设置一个文件路径）
	router.StaticFile("favicon.ico", "./static/dist/assets/imgs/favicon.ico")

	//加载html文件,templates有目录，索引需要指定加载层级
	//router.LoadHTMLGlob("./static/dist/templates/**/*")
	//也可以使用LoadHTMLFiles指定多个路径
	htmlFiles := listAllHtmlFiles()
	if htmlFiles != nil {
		router.LoadHTMLFiles(htmlFiles...)
	}

	//这种方式文件夹下不能有文件夹（当所有的html文件都在templates目录下时，就可以用这种方式加载）
	//router.LoadHTMLGlob("./static/dist/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{
			"title": "index",
			"dataM": map[string]string{"k1": "v1", "k2:": "v2"},
			"dateS": []string{"a", "b"},
		})
	})

	if err := router.Run(":8888"); err != nil {
		log.Fatal(err)
	}
}

func listAllHtmlFiles() (files []string) {
	err := filepath.Walk("./static/dist/templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("listAllHtmlFiles error", err)
	}
	return
}
