package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	cnt := int32(0)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt32(&cnt, 1)
		}()
	}
	fmt.Println(cnt)
	time.Sleep(time.Minute)
}
