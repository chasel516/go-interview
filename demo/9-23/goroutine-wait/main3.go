package main

import (
	"fmt"
)

func worker1(id int, done chan bool) {
	// 执行协程任务
	fmt.Printf("Worker %d done\n", id)
	done <- true
}

func main() {
	numWorkers := 5
	done := make(chan bool, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker1(i, done)
	}

	for i := 0; i < numWorkers; i++ {
		<-done
	}

	fmt.Println("All workers done")
	close(done)
}
