package main

import (
	"log"
	"viper-demo/config"
)

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig()
}
func main() {
	log.Printf("%+v", config.Cfg)
	select {}

}
