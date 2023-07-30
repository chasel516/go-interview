package main

import "C"
import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func init() {
	//监控goroutine数量
	Ticker(func() {
		log.Print("current go routine num: ", runtime.NumGoroutine())
	}, time.Second)
}

func main() {
	mutexTest()
	//channelTest()

	time.Sleep(10 * time.Second)
}

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

func processFiles() error {
	for i := 0; i < 100000; i++ {
		file, err := os.Open("file" + strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		defer file.Close()

		// 处理文件内容
		// ...
	}
	return nil
}

// 将文件操作的逻辑封装起来，处理完一个文件就在函数中关闭资源
func processFiles1() error {
	for i := 0; i < 100000; i++ {
		processFile("file" + strconv.Itoa(i) + ".txt")
	}
	return nil
}

func processFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 处理文件内容
	// ...
	return nil
}

func mutexTest() {
	mutex := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			//假设程序读取了大对象
			//未释放锁导致内存泄露
			mutex.Lock()
			// do something
		}()
		time.Sleep(time.Millisecond * 10)
	}
}

func channelTest() {
	//声明未初始化的channel读写都会阻塞
	var c chan int
	for i := 0; i < 1000; i++ {
		//向channel中写数据
		go func() {
			//假设程序读取了大对象
			c <- 1
			fmt.Println("g1 send succeed")
			time.Sleep(1 * time.Second)
		}()
		//从channel中读数据
		go func() {
			//假设程序读取了大对象
			<-c
			fmt.Println("g2 receive succeed")
			time.Sleep(1 * time.Second)
		}()
		time.Sleep(10 * time.Millisecond)
	}
}

//// cgo_example.c
//#include <stdlib.h>
//#include <string.h>
//
//char* get_string() {
//char* s = malloc(100);
//strcpy(s, "hello world");
//return s;
//}
//
//// cgo_example.go
//package main
//
///*
//   #include "cgo_example.c"
//*/
//import "C"
//import "fmt"
//
//func main() {
//	s := C.get_string()
//	fmt.Println(C.GoString(s))
//	// 没有调用 C.free(s) 来释放内存
//}
