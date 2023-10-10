package main

import (
	"fmt"
	"sync"
)

func main() {
	//runtime.GOMAXPROCS(1)
	counter1()
}

func counter1() {
	cnt := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt)
}
