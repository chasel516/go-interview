package main

import (
	"fmt"
	"log"
)

func main() {
	//nativeGo()
	//nativeGoFunc()
	goWithPointerData()
	select {}
}

//1. 并发访问共享资源：在使用goroutine和闭包时，需要注意并发访问共享资源，例如map、切片等。在并发场景下，需要使用互斥锁（sync.Mutex）或者原子操作（sync/atomic）来保护共享资源。
//2. 同步与顺序问题：在使用goroutine和闭包时，需要注意同步问题。如果需要等待所有的goroutine执行完毕，可以使用sync.WaitGroup来实现同步。同时，如果需要顺序执行goroutine，可以使用通道（channel）来实现。

// 通过go关键词启动
func nativeGo() {
	for i := 0; i < 10; i++ {
		//变量i和go声明同时估值，i的值跟当前迭代的值同步
		go fmt.Println(i)
	}
}

// 通过匿名函数启动
func nativeGoFunc() {
	for i := 0; i < 10; i++ {
		//x := i
		//由于打印语句在匿名函数内，匿名函数执行时打印语句才会被执行，此时i才会被估值，但随着迭代的进行，i的值大概率已经是迭代结束后的值
		go func() {
			log.Println(i)
		}()
	}
}

type S struct {
	F1 string
	F2 int
	F3 int
}

func goWithPointerData() {
	s1 := &S{
		F1: "test",
		F2: 1,
		F3: 0,
	}
	for i := 0; i < 10; i++ {
		s2 := s1
		s2.F3 = i

		writeData(s2)
	}

}

func writeData(s *S) {
	//s是指针类型，在协程外有修改，内部访问的值不确定
	go func() {
		log.Println("F3:", s.F3)
	}()
}
