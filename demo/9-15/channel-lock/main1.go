package main

import (
	"log"
	"sync"
	"time"
)

// channel实现计数信号量
const jobNum = 9

var active = make(chan struct{}, 3)
var jobs = make(chan int, jobNum)

func main() {
	go func() {
		produceJobs(jobNum)
	}()
	var wg sync.WaitGroup
	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}

func produceJobs(n int) {
	for i := 0; i < n; i++ {
		jobs <- (i + 1)
	}
	close(jobs)
}
