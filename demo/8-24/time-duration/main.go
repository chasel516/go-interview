package main

import (
	"fmt"
	"time"
)

func main() {
	//获取500ms后的执行戳，单位ms

	start := time.Now().UnixMilli()
	fmt.Println(start)
	end := start + int64(time.Millisecond*500)
	end1 := time.Duration(start) + time.Millisecond*500
	fmt.Println(end)
	fmt.Println(int64(end1), end1)

}
