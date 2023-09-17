package main

import (
	"fmt"
	"log"
	"runtime"
)

const Cnt = 100000

var m = make(map[int]int, Cnt)

func init() {
	log.SetFlags(log.Lshortfile)
}
func main() {

	for i := 0; i < Cnt; i++ {
		m[i] = i
	}
	runtime.GC()
	log.Println(getMemStats())
	//模拟大量，map删除的场景
	for k, _ := range m {
		if k != 1 {
			delete(m, k)
		}

	}

    //注意，这里只是为了演示而手动触发GC，线上环境不推荐这样做
	runtime.GC()
	log.Println(getMemStats())

	//将原map的值拷贝到新map
	tmp := make(map[int]int, len(m))
	for k, v := range m {
		tmp[k] = v
	}
	//将新map置空
	m = nil
	//将临时map赋值给新map
	m = tmp
	//将临时map置空
	tmp = nil
	runtime.GC()
	log.Println(m)
	log.Println(getMemStats())
}

func getMemStats() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("分配的内存 = %vKB, GC的次数 = %v\n", m.Alloc/1024, m.NumGC)
}
