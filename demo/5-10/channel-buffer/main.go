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
	f9()
	//f10()
	//var ch chan struct{}
	//
	//for i := 0; i < 3; i++ {
	//	go test()
	//}
	//fmt.Println("永久阻塞")
	//<-ch
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

//协程中只写不读
func f7() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	time.Sleep(time.Second * 3)

}

//协程中只读不写
func f8() {
	ch := make(chan int)
	go func() {
		<-ch
	}()
	time.Sleep(time.Second * 3)

}

//当我们在
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

	go func() {
		for {
			select {
			case <-ch1:
				ch <- 1
			}
		}
	}()

}

func f10() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(<-ch)
}

func test() {
	for {
		time.Sleep(time.Second)
	}
}
