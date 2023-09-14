package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	//m := make(map[int]int)
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		x := i
		wg.Add(1)
		go func() {
			//m[x] = x
			m.Store(x, x)
			wg.Done()
		}()
	}

	wg.Wait()

	//按key读取
	value, ok := m.Load(1)
	if ok {
		fmt.Println(value)
	}

	//遍历
	m.Range(func(key, value any) bool {
		fmt.Println("key:", key, "value:", value)
		if key == 3 {
			// 终止遍历
			return false
		}
		return true
	})

	//支持不同类型的k/v
	m.Store("k1", "v1")
	v1, ok := m.Load("k1")
	if ok {
		fmt.Println(v1)
	}

	//存在则返回，否则写入
	actual, loaded := m.LoadOrStore("k1", "v2")
	fmt.Println("actual:", actual, "loaded:", loaded)

	//存在则删除并返回
	v, existed := m.LoadAndDelete("k1")
	fmt.Println("v:", v, "existed:", existed)
	if _, ok := m.Load("k1"); !ok {
		fmt.Println("k1 not found")
	}
}
