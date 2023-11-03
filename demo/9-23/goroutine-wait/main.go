package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 执行协程任务
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
