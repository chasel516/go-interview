package main

import (
	"encoding/json"
	"runtime"
	"sync"
)

func main() {
	// 设置使用多个逻辑处理器
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(1)

	// 第一个goroutine，被锁定在第一个逻辑处理器上
	go func() {
		defer wg.Done()
		for i := 0; i < 10000000000000; i++ {
			json.Marshal(i)
		}
	}()

	// 启动更多goroutine
	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10000000000000; i++ {
				json.Marshal(i)
			}
		}()
	}

}
