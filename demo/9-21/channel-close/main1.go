package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
		//close(ch)
	}()
	for v := range ch {
		log.Println(v)
	}
	fmt.Println("退出for range")

	/**
		for {
			select {
			case x, ok := <-ch:
				log.Println(x, ok)
				//if !ok {
				//	return
				//} else {
				//	log.Println(x)
				//}
			}
		}
	**/

	log.Println("exit")
}
