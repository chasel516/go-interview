package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 参数值
	paramValue := 123

	// 创建Goroutine并传递参数
	go func(param int) {
		fmt.Println("Received parameter:", param)
	}(paramValue)

	// 手动触发垃圾回收以确保Goroutine执行
	runtime.GC()

	// 等待Goroutine执行完成
	runtime.Goexit()
}
