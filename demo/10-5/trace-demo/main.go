package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
	"time"
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
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(time.Millisecond)
		}()
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println(<-c)
}
