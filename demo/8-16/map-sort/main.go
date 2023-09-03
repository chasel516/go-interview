package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)
	keys := make([]int, 0, 3)
	m[1] = "v1"
	m[2] = "v2"
	m[3] = "v3"
	for k, v := range m {
		keys = append(keys, k)
		fmt.Println("k:", k, "v:", v)
	}

	//对key进行排序
	sort.Ints(keys)

	//根据key的值对map进行遍历
	for _, key := range keys {
		fmt.Println("sorted key :", key, "value:", m[key])
	}
}
