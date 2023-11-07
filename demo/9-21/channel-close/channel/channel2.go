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

	//为了提高性能，第一次判断相当于快速路径，不加锁。如果已经关闭就直接返回
	if c.closed {
		return
	}
	//如果快速路径判断出未关闭还得进行加锁判断，避免判断时被其他协程修改了这个状态
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
