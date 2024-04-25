package main

import (
	"log"
	"time"
)

func main() {
	// timer := time.NewTimer(5 * time.Second)
	// timer.Reset(1 * time.Second)
	// select {
	// case <-timer.C:
	// 	log.Println("Delayed 5s, start to do something.")
	// }

	// log.Println(time.Now())
	// <-time.After(1 * time.Second)
	// log.Println(time.Now())

	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})

	time.Sleep(2 * time.Second) // 等待协程退出

}
