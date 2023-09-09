package main

import (
	"fmt"
	"time"
)

func main() {
	afterUnixMilli1()
	//afterUnixMilli2()
}

// 获取500ms后的时间戳，单位ms
func afterUnixMilli1() {
	nowUnixMilli := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).UnixMilli()
	fmt.Println("afterUnixMilli1: now", nowUnixMilli)
	//延迟后的时间戳 = 指定的时间戳+延时时间
	end := nowUnixMilli + int64(time.Millisecond*500)
	end1 := time.Duration(nowUnixMilli) + time.Millisecond*500
	fmt.Println(end)
	fmt.Println(int64(end1), end1)
}

// 获取500ms后的执行戳，单位ms
func afterUnixMilli2() {
	nowUnixMilli := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).UnixMilli()
	end := nowUnixMilli + 500
	delayduration := 500 * time.Millisecond
	end1 := nowUnixMilli + delayduration.Milliseconds()
	fmt.Println(end)
	fmt.Println(end1)
}
