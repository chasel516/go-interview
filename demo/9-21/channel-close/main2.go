package main

import (
	"fmt"
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

func main() {
	c := NewChannel()
	c.Close()
	c.Close()
	_, ok := <-c.C
	fmt.Println(ok)
}
