package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	num int
}

func (c *Counter) incr(delt int) {
	c.Lock()
	defer c.Unlock()
	c.num += delt
}

func (c *Counter) get() int {
	c.Lock()
	defer c.Unlock()
	return c.num
}

func main() {
	c := Counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.incr(1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.get())
}
