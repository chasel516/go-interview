package main

import "sync"

type Channel2 struct {
	C      chan any
	closed bool
	lock   sync.Mutex
}

func NewChannel2() *Channel2 {
	return &Channel2{
		C:      make(chan any),
		closed: false,
		lock:   sync.Mutex{},
	}
}
func (c *Channel2) close() {
	if c.closed {
		return
	}
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.closed {
		close(c.C)
		c.closed = true
	}
}

func (c *Channel2) iSClose() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.closed
}

func main() {
	ch := NewChannel2()
	ch.close()
	ch.close()
}
