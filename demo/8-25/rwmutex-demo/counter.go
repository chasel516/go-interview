package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int64
	lock  sync.RWMutex
}

func (c *Counter) ReadValue() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.count
}
func (c *Counter) ReadValue1() int64 {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.count
}

func (c *Counter) ReadValue2() int64 {
	return c.count
}
func (c *Counter) Incr(delta int64) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.count += delta
}

func main() {
	c := Counter{}
	c.Incr(1)
	const Num = 10000000
	t := time.Now()
	for i := 0; i < Num; i++ {
		go func() {
			c.ReadValue()
		}()
	}
	fmt.Println("ReadValue cost:", time.Since(t).Milliseconds())

	t1 := time.Now()
	for i := 0; i < Num; i++ {
		go func() {
			c.ReadValue1()
		}()
	}
	fmt.Println("ReadValue1 cost:", time.Since(t1).Milliseconds())

	t2 := time.Now()
	for i := 0; i < Num; i++ {
		go func() {
			c.ReadValue2()
		}()
	}
	fmt.Println("ReadValue2 cost:", time.Since(t2).Milliseconds())
}
