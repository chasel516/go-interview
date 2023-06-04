package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"log"
	"runtime"
	"sync/atomic"
	"time"
)

func init() {
	Ticker(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, 2*time.Second)
}
func main() {
	//nativeGo()
	//nativeGoFunc()
	goWithPointerData()
	select {}
	return
	// 创建一个Ants池，最多允许5个goroutine同时执行
	p, _ := ants.NewPool(5)
	defer p.Release()

	// 定义一个计数器，用于统计处理的任务数
	var counter int64

	// 向Ants池提交100个任务
	for i := 0; i < 100; i++ {
		x := i
		err := p.Submit(func() {
			// 模拟任务处理时间
			time.Sleep(time.Second)
			log.Println("执行任务", x)
			// 增加计数器的值
			atomic.AddInt64(&counter, 1)
		})
		if err != nil {
			log.Println("Submit error", err)
		}
	}

	log.Println("正在等待执行的任务数：", p.Waiting())

	// 输出计数器的值
	log.Println("counter:", counter)

	//当Running的任务数为0说明任务全部执行完
	for p.Running() > 0 {
		time.Sleep(100 * time.Millisecond)
	}
	select {}
}

// 监控系统当前的协程数量：
// 启动一个定时器
func Ticker(f func(), d time.Duration) {
	go func() {
		ticker := time.NewTicker(d)
		for {
			select {
			case <-ticker.C:
				go f()
			}
		}
	}()

}

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
