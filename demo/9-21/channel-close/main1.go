package main

import (
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func main() {
	ch := make(chan int)
	//close(ch)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
		close(ch)
	}()
	//log.Println("开始for range")
	//for v := range ch {
	//	log.Println(v)
	//}
	//log.Println("退出for range")

	for {
		select {
		case x, ok := <-ch:
			time.Sleep(time.Second)
			//log.Println(x, ok)
			if !ok {
				return
			} else {
				log.Println(x)
			}
		}
	}

	log.Println("exit")
}
