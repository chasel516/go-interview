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
	f10()
	//f11()
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

// 缓冲区大小为1的channel中的两个goroutine相互等待
func f14() {
	ch := make(chan int, 1)
	ch <- 1 // 发送方向缓冲区中放入数据
	<-ch    // 接收方从缓冲区中取出数据
	ch <- 2 // 发送方尝试向缓冲区中发送数据，但接收方已经退出，导致发送方阻塞
}

//在这个例子中，我们创建了一个大小为1的缓冲区，并向其中发送了一个数据，然后从中接收了一个数据。由于缓冲区大小为1，发送方尝试向缓冲区中发送数据时将被阻塞，而由于接收方已经退出，发送方也无法继续执行，从而导致死锁。
