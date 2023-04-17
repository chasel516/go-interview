package main

import (
	"fmt"
	"time"
)

func main() {
	//var c chan int
	//var c2 chan<- int
	//var c3 <-chan int
	//c2 = c
	//c3 = c
	//c = c2  //编译不通过
	//c = c3  //编译不通过
	//c2 = c3 //编译不通过
	//fmt.Println(c2, c3)
	//f1()
	//f2()
	//f3()
	//f4()
	//f4_1()
	//f5()
	//f6()
	//f7()
	//f8()
	//f9()
	//f10()
	//f11()
	//f12()
	//f12_1()
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

// 协程中只写不读
func f2() {
	ch := make(chan int)
	go func() {
		fmt.Println("开始写入ch")
		ch <- 1
		fmt.Println("写入ch完成") //没有机会执行
	}()

	time.Sleep(time.Second * 3)

}

// 协程中只读不写
func f3() {
	ch := make(chan int)
	go func() {
		fmt.Println("开始读取ch")
		<-ch
		fmt.Println("读取ch完成") //没有机会执行
	}()
	time.Sleep(time.Second * 3)

}

// 分别在不同协程中读写
func f4() {
	ch := make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()

	go func() {
		ch <- 1
	}()
	time.Sleep(time.Second)
}

// 新协程中读写
func f4_1() {
	ch := make(chan int)
	go func() {
		fmt.Println("开始读ch")
		fmt.Println(<-ch) //阻塞住随主携程退出
		fmt.Println("不会执行到这里")
		ch <- 1
	}()

	time.Sleep(time.Second)
}

func f5() {
	ch := make(chan int)
	fmt.Println(<-ch)
	go func() {
		ch <- 1
	}()
}

func f6() {
	ch := make(chan int)
	ch <- 1
	<-ch
}

func f7() {
	ch := make(chan int)
	ch <- 1
}

func f8() {
	ch := make(chan int)
	<-ch
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
	go test()
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

// 只接收没有发送导致死锁
func f12_1() {
	ch := make(chan int, 1)
	fmt.Println(<-ch)
}

// 缓冲区为空，接收方阻塞而无法开启新的发送方而导致死锁
func f13() {
	ch := make(chan int, 1)
	<-ch
	ch <- 1
}

// 两个goroutine中有缓冲的channel相互等待而产生死锁
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
