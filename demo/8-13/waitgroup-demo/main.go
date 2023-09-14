package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		//增加异步任务
		wg.Add(1)
		x := i
		go func() {
			task(x, &wg)
		}()
	}
	//等待全部goroutine完成
	wg.Wait()
}

func task(i int, wg *sync.WaitGroup) {
	fmt.Println(i)
	//标记异步任务完成
	wg.Done()
}
