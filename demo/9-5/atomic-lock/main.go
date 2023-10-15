package main

import (
	"log"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func main() {
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)
	go func() {
		//Wait()方法的调用必须加锁
		cond.L.Lock()
		log.Println("Wait1 start")
		//开始等待，需要调用Signal()或者Broadcast()方法唤醒
		cond.Wait()
		log.Println("Wait1 end")
		cond.L.Unlock()
	}()
	go func() {
		//cond.L.Lock()  //cond.L跟mu是同一个锁
		mu.Lock()
		log.Println("Wait2 start")
		cond.Wait()
		log.Println("Wait2 end")
		mu.Unlock()
		//cond.L.Unlock()
	}()
	time.Sleep(time.Second * 2)
	//调用Signal方法只能让上面第一个Wait()得到释放
	//cond.Signal()
	//调用Broadcast()方法能释放全部Wait()
	cond.Broadcast()
	log.Println("Signal end")
	time.Sleep(time.Second * 10)
}
