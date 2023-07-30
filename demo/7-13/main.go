package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func init() {
	goMaxProcs := 2
	cpuCoreNum := runtime.GOMAXPROCS(goMaxProcs)
	Ticker(func() {
		if goMaxProcs < cpuCoreNum {
			goMaxProcs += 1
			runtime.GOMAXPROCS(goMaxProcs)
			fmt.Println("goMaxProcs:", goMaxProcs)
		}
	}, time.Second*5)
}
func main() {

	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			x := 0
			for i := 0; i < math.MaxInt; i++ {
				x++
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

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
