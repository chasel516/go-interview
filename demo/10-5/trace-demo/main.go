package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	c := make(chan int)
	for i := 0; i < 100000000; i++ {
		go func(index int) {
			if index%2 == 0 {
				<-c
			} else {
				c <- index
			}
		}(i)
	}
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-sig)
}
