package main

import (
	"fmt"
	"sync"
)

type operation int

const (
	opGet operation = iota
	opSet
	opAppend
	opLen
	opRange
)

// 将所有对slice的操作丢到channel中，利用channel的并发安全来避免slice的并发操作
type request struct {
	op    operation
	index int
	value any
	resp  chan any
}

type ConcurrentSlice struct {
	s        []any
	requests chan request
}

func NewConcurrentSlice() *ConcurrentSlice {
	cs := &ConcurrentSlice{
		s:        make([]any, 0),
		requests: make(chan request),
	}
	go cs.run()
	return cs
}

func (s *ConcurrentSlice) run() {
	for req := range s.requests {
		switch req.op {
		case opGet:
			req.resp <- s.s[req.index]
		case opSet:
			s.s[req.index] = req.value
			req.resp <- nil
		case opAppend:
			s.s = append(s.s, req.value)
			req.resp <- nil
		case opLen:
			req.resp <- len(s.s)
		case opRange:
			f := (req.value).(func(index int, value any) bool)
			for k, v := range s.s {
				if !f(k, v) {
					break
				}
			}
			req.resp <- nil
		}
	}
}

func (s *ConcurrentSlice) Get(index int) any {
	resp := make(chan interface{})
	s.requests <- request{
		op:    opGet,
		index: index,
		resp:  resp,
	}
	return <-resp
}

func (s *ConcurrentSlice) Set(index int, value any) {
	resp := make(chan interface{})
	s.requests <- request{
		op:    opSet,
		index: index,
		value: value,
		resp:  resp,
	}
	<-resp
}

func (s *ConcurrentSlice) Append(value any) {
	resp := make(chan interface{})
	s.requests <- request{
		op:    opAppend,
		value: value,
		resp:  resp,
	}
	<-resp
}

func (s *ConcurrentSlice) Len() int {
	resp := make(chan interface{})
	s.requests <- request{
		op:   opLen,
		resp: resp,
	}
	return (<-resp).(int)
}

// 使用匿名函数的入参接收迭代后的值
func (s *ConcurrentSlice) Range(value func(key, value any) bool) {
	resp := make(chan interface{})
	s.requests <- request{
		op:    opRange,
		value: value,
		resp:  resp,
	}
	<-resp
}

func main() {
	s := NewConcurrentSlice()
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			s.Append(n)
		}(i)
	}
	wg.Wait()
	fmt.Println("Length of ConcurrentSlice:", s.Len())
}
