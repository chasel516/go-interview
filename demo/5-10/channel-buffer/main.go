package main

import (
	"fmt"
	"time"
)

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()
	//f8()
	//f9()
	//f10()
	//f11()
	//f12()
	//f13()
	f14()
}

func f1() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

func f2() {
	ch := make(chan int)
	fmt.Println(<-ch)
	go func() {
		ch <- 1
	}()
}

func f3() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()

	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
}

func f4() {
	ch := make(chan int)
	ch <- 1
	<-ch
}

func f5() {
	ch := make(chan int)
	ch <- 1
}

func f6() {
	ch := make(chan int)
	<-ch
}

// 协程中只写不读
func f7() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	time.Sleep(time.Second * 3)

}

// 协程中只读不写
func f8() {
	ch := make(chan int)
	go func() {
		<-ch
	}()
	time.Sleep(time.Second * 3)

}

func f9() {
	ch := make(chan int)
	ch1 := make(chan int)
	go func() {
		for {
			select {
			case <-ch:
				ch1 <- 1
			}
		}
	}()

	for {
		select {
		case <-ch1:
			ch <- 1
		}
	}

}

func f10() {
	var ch chan int
	for i := 0; i < 3; i++ {
		go test()
	}
	fmt.Println("永久阻塞")
	<-ch
	//ch <- 1
}

func test() {
	for {
		time.Sleep(time.Second)
	}
}

func f11() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
}

// 缓冲区已满，发送方阻塞，没有接收方而发生死锁
func f12() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 2
}

// // 缓冲区为空，接收方阻塞而无法开启新的发送方而导致死锁
func f13() {
	ch := make(chan int, 1)
	<-ch
	ch <- 1
}

//两个goroutine中有缓冲的channel相互等待而产生死锁
func f14() {
	ch := make(chan int, 2)
	ch1 := make(chan int, 2)
	go func() {
		for {
			select {
			case <-ch:
				ch1 <- 1
			}
		}
	}()

	for {
		select {
		case <-ch1:
			ch <- 1
		}
	}

}
