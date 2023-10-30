package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer := time.NewTimer(time.Second / 2)
	select {
	case <-timer.C:
	default:
		time.Sleep(time.Second) // 此分支被选中的可能性较大
	}
	timer.Reset(time.Second * 10) // 可能数据竞争
	<-timer.C
	fmt.Println(time.Since(start)) // 大约1s
}
