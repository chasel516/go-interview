package main

import "sync"

func main() {
	wg := sync.WaitGroup{}
	//m := make(map[int]int)
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		x := i
		wg.Add(1)
		go func() {
			//m[x] = x
			m.Store(x, x)
			wg.Done()
		}()
	}
	wg.Wait()
}
