package main

import (
	"fmt"
	"log"
	"time"
)

var cache = make(map[int]int)

func init() {
	Ticker(refreshCache, time.Second)
}
func main() {
	for i := 0; i < 100000000; i++ {
		//模拟非并发写入
		cache[i] = i
	}
	time.Sleep(time.Second)
	fmt.Println(cache)
}

//全量更新缓存
func refreshCache() {
	log.Println("开始更新缓存")
	data := GetDataFromDB()
	if len(data) == 0 {
		return
	}

	tmp := make(map[int]int, len(data))
	for i, d := range data {
		tmp[i] = d
	}
	//全量更新cache
	cache = tmp
	tmp = nil
}

func GetDataFromDB() []int {
	return []int{1, 2, 3}
}

// 启动一个定时器
func Ticker(f func(), d time.Duration) {
	go func() {
		ticker := time.NewTicker(d)
		for {
			select {
			case <-ticker.C:
				go f()
			}
		}
	}()
}
