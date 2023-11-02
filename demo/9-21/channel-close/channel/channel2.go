package channel

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
func (c *Channel2) Close() {
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

func (c *Channel2) ISClose() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.closed
}
