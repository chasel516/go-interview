package main

import "fmt"

// 死锁检测
func worker(data chan string) {
	for {
		result := <-data
		fmt.Println(result)
	}
}

func main() {
	data := make(chan string)

	go worker(data)

	data <- "pidancode.com"

	// 如果忘记关闭信道，程序就会一直阻塞
	// close(data)
}
