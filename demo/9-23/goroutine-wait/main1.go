package main

import (
	"log"
	"sync"
)

func main() {
	data := make([]int, 0)
	var wg1, wg2 sync.WaitGroup
	wg2.Add(1)
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		i := i
		go func() {
			wg2.Wait() //等待数据准备好(等待广播通知)
			log.Printf("date[%d]=%d", i, data[i])
			//wg1.Add(-1)
			wg1.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		data = append(data, i*10)
	}
	//数据已经准备好。发送广播通知读取的协程继续执行
	wg2.Done()
	wg1.Wait()
}
