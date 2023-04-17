package main

import (
	"fmt"
	"time"
	"unsafe"
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

}

func f1() {
	ch := make(chan int)
	close(ch)
	go func() {
		fmt.Println("开始接收数据")
		fmt.Println("<-ch:", <-ch)
		fmt.Println("数据接收完成")
	}()

	go func() {
		fmt.Println("开始发送数据")
		ch <- 1
		fmt.Println("发送数据完毕")
	}()
	time.Sleep(time.Second)
}

func f2() {
	ch := make(chan int)
	close(ch)
	for {
		select {
		case c := <-ch:
			fmt.Println("读取<-ch:", c)
			time.Sleep(time.Second)
		}
	}
}

func f3() {
	ch := make(chan int)
	close(ch)
	<-ch          //读取已经关闭的channel ，不会阻死
	println("完成") //这句会输出
}

func f4() {
	ch := make(chan int)
	go func() {
		close(ch)
		println("已关闭ch") //这句会输出
	}()
	<-ch          //这里读取时也不会组塞住,会等到gorutine中关闭
	println("完成") //这句会输出
}

func f5() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}
}

func f6() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		select {
		case x, ok := <-ch:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(x)
		}
	}
}

// 多次关掉 channel 会触发运行时错误
func f7() {
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 1
		close(c)
	}()
	fmt.Println(<-c)
}

// 使用 defer 延迟关闭 channel，并且确保 channel 只释放一次
func f8() {
	c := make(chan int)
	defer close(c)
	go func() {
		c <- 1
	}()

	fmt.Println(<-c)
}

// 关闭channel之后，channel中的数据依然是可以读取到
func f9() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()

	fmt.Println(<-c)
}

func f10() {
	arr := [4095]string{}
	fmt.Println(unsafe.Sizeof(arr))
	ch := make(chan [4095]string)
	go func() {
		ch <- arr
	}()
	time.Sleep(time.Second)
}
