package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	// 参数值
	paramValue := 123

	// 创建Goroutine并传递参数
	go func(param unsafe.Pointer) {
		// 类型转换获取参数的具体值
		value := *(*int)(param)
		fmt.Println("Received parameter:", value)
	}(unsafe.Pointer(&paramValue))

	// 手动触发垃圾回收以确保Goroutine执行
	runtime.GC()

	// 等待Goroutine执行完成
	runtime.Goexit()
}
