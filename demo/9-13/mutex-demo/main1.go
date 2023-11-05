package main

import (
	"log"
	"sync"
)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {
	lock := sync.Mutex{}

	//Go 1.18 新增，是一种非阻塞模式的取锁操作。当调用 TryLock() 时，
	//该函数仅简单地返回 true 或者 false，代表是否加锁成功
	//在某些情况下，如果我们希望在获取锁失败时，并不想停止执行，
	//而是可以进入其他的逻辑就可以使用TryLock()
	lock.Lock()
	log.Println("TryLock：", lock.TryLock())
	//已经通过TryLock（）加锁，不能再次加锁

}
