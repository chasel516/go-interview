package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		x := i
		wg.Add(1)
		go func() {
			m[x] = x
			wg.Done()
		}()
	}
	wg.Wait()
}
