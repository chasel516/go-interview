package main

import (
	"fmt"
	"runtime"
)

func main() {
	const numElements = 10000000
	var beforeGC, afterGC runtime.MemStats
	runtime.ReadMemStats(&beforeGC)
	var data []int

	for i := 0; i < numElements; i++ {
		data = append(data, i)
		processData(data)
	}
	runtime.ReadMemStats(&afterGC)

	// 输出GC统计信息
	fmt.Printf("GC 次数: %d\n", afterGC.NumGC-beforeGC.NumGC)
	fmt.Printf("GC 暂停时间: %.4f ns\n", float64(afterGC.PauseTotalNs)-float64(beforeGC.PauseTotalNs))
}

func processData(data []int) {
	_ = len(data)
}
