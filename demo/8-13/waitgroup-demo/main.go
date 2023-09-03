package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(uint64(2) << 32)
	//wg := sync.WaitGroup{}
	//for i := 0; i < 10; i++ {
	//	//增加异步任务
	//	wg.Add(1)
	//	go func() {
	//		task(i, wg)
	//	}()
	//}
	////等待全部goroutine完成
	//wg.Wait()
}

func task(i int, wg sync.WaitGroup) {
	fmt.Println(i)
	//标记异步任务完成
	wg.Done()
}
