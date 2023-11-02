package channel

import (
	"sync"
)

type Channel struct {
	C    chan any
	once sync.Once
}

func NewChannel() *Channel {
	return &Channel{
		C:    make(chan any),
		once: sync.Once{},
	}
}

func (ch *Channel) Close() {
	ch.once.Do(func() {
		close(ch.C)
	})
}
