package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	fmt.Println("start")
	//f1()
	//f2()
	//f3()
	//f4()
	f5()

	fmt.Println("exit")

}
func f1() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()
	for {
	}
}

func f2() {
	ch := make(chan struct{})
	go func() {
		for {
			fmt.Println("后台协程正在运行...")
			time.Sleep(time.Second)
		}
	}()
	//<-ch // 阻塞主协程
	ch <- struct{}{} // 阻塞主协程
}

func f3() {
	var ch chan struct{}
	go func() {
		for {
			fmt.Println("后台协程正在运行...")
			time.Sleep(time.Second)
		}
	}()
	//<-ch // 阻塞主协程
	ch <- struct{}{} // 阻塞主协程
}
func f4() {
	go func() {
		for {
			fmt.Println("后台协程正在运行...")
			time.Sleep(time.Second)
		}
	}()

	// 让主协程永久阻塞
	select {}
}

func f5() {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println(<-sig)
}
