package main

import "time"

type Context interface {
	// 当 context 被取消或者到了 deadline，返回一个被关闭的 channel
	//它是一个只读的Channel，也就是说在整个生命周期都不会有写入操作，只有当这个channel被关闭时，
	//才会读取到这个channel对应类型的零值，否则是无法读取到任何值的。
	//正式因为这个机制，当子协程从这个channel中读取到零值后可以做一些收尾工作，让子协程尽快退出。
	Done() <-chan struct{}

	// 在 channel Done 关闭后，返回 context 取消原因，这里只有两种原因，取消或者超时。
	Err() error

	// 返回 context 是否会被取消以及自动取消时间（即 deadline）
	//通过这个时间我们可以判断是否有必要进行接下来的操作，如果剩余时间太短则可以选择不继续执行一些任务，可以节省系统资源。
	Deadline() (deadline time.Time, ok bool)

	// 获取 key 对应的 value
	Value(key interface{}) interface{}
}
