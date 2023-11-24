package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
	"time"
)

const (
	cpuprofile   = "cpu.prof"
	memprofile   = "mem.prof"
	mutexprofile = "mutex.prof"
	blockprofile = "block.prof"
)

func main() {
	cpuf, err := os.Create(cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer cpuf.Close()
	if err := pprof.StartCPUProfile(cpuf); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	memf, err := os.Create(memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer memf.Close()
	if err := pprof.WriteHeapProfile(memf); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}

	//开启mutexprofile指标
	//每 rate 次互斥锁争用事件中，会报告一个事件（传值越高采集频率越低）
	//返回之间设置的值
	//传负数只会返回之前的设置不会修改
	//传0关闭采集
	runtime.SetMutexProfileFraction(1)
	defer runtime.SetMutexProfileFraction(0)
	mutexf, err := os.Create(mutexprofile)
	if err != nil {
		log.Fatal("could not create mutex profile: ", err)
	}
	defer mutexf.Close()

	if mp := pprof.Lookup("mutex"); mp != nil {
		mp.WriteTo(mutexf, 0)
	}

	runtime.SetBlockProfileRate(1)
	defer runtime.SetBlockProfileRate(0)
	blockf, err := os.Create(blockprofile)
	if err != nil {
		log.Fatal("could not create block profile: ", err)
	}
	defer blockf.Close()

	if mp := pprof.Lookup("block"); mp != nil {
		mp.WriteTo(blockf, 0)
	}

	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		for {
			select {
			case <-c:
				wg.Done()
				return
			default:
				s1 := [1024]byte{}
				s2 := [2048]byte{}
				_ = string(s1[:]) + string(s2[:])
			}
			time.Sleep(20 * time.Millisecond)
		}
	}()
	wg.Wait()
	fmt.Println("main exit")
}
