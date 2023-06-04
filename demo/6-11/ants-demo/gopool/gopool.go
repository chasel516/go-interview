package gopool

import (
	"fmt"
	"log"
	"runtime/debug"
	"sync/atomic"
	"time"
)

// 定义超时错误
var ErrScheduleTimeout = fmt.Errorf("schedule error: timed out")

// 定义协程池结构体
// 包含了两个channel：concurrency和work
// concurrency是用来控制并发度的， 在创建协程池时会传进来并发数，将这个并发数作为concurrency这个channel的容量，
// 在执行任务前向concurrency写入一个零尺寸的空结构体，待任务结束后读出这个空结构体释放channel的缓冲容量来达到控制并发度的功能

// work为需要执行的任务队列
type Pool struct {
	concurrency chan struct{}
	work        chan func()
	running     int32
}

// 根据给定的大小创建一个协程池，并立即启动给定数量的协程
func NewPool(concurrencyNum, queue, workers int) *Pool {
	if workers <= 0 && queue > 0 {
		panic("dead queue configuration detected")
	}
	if workers > concurrencyNum {
		panic("workers > concurrencyNum")
	}
	p := &Pool{
		concurrency: make(chan struct{}, concurrencyNum),
		work:        make(chan func(), queue),
	}
	for i := 0; i < workers; i++ {
		//若concurrency已满会阻塞
		p.concurrency <- struct{}{}
		go p.run(func() {})
	}

	return p
}

// 提交任务
func (p *Pool) Submit(task func()) {
	p.submit(task, nil)
}

// 提交任务，带超时时间
func (p *Pool) SubmitWithTimeout(timeout time.Duration, task func()) error {
	return p.submit(task, time.After(timeout))
}

// 当concurrency容量未满说明仍在并发度限制内，则直接启动任务，否则提交到任务队列
// 当达到指定的超时时间会返回超时错误
func (p *Pool) submit(task func(), timeout <-chan time.Time) error {
	select {
	case <-timeout:
		return ErrScheduleTimeout
	case p.work <- task:
		p.addRunning(1)
		return nil
	case p.concurrency <- struct{}{}:
		go p.run(task)
		return nil
	}
}

// 执行当前的任务和任务队列中的任务
func (p *Pool) run(task func()) {
	defer func() {
		p.addRunning(-1)
		if err := recover(); err != nil {
			//捕获异常，并打印错误堆栈
			log.Println("task panic", err, string(debug.Stack()))
		}
		<-p.concurrency
	}()

	task()

	for task := range p.work {
		task()
		p.addRunning(-1)
	}
}
func (p *Pool) addRunning(delta int) {
	atomic.AddInt32(&p.running, int32(delta))
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}
