package main

import "fmt"

func main() {

	//错误写法
	//slice := []int{0, 1, 2, 3}
	//m := make(map[int]*int)
	//for key, val := range slice {
	//	m[key] = &val
	//}

	//0 -> 3
	//1 -> 3
	//2 -> 3
	//3 -> 3
	//for k, v := range m {
	//	fmt.Println(k, "->", *v)
	//}

	////正确写法
	//slice := []int{0, 1, 2, 3}
	//m := make(map[int]*int)
	//for key, val := range slice {
	//	value := val
	//	m[key] = &value
	//}
	//
	////0 -> 0
	////1 -> 1
	////2 -> 2
	////3 -> 3
	//for k, v := range m {
	//	fmt.Println(k, "->", *v)
	//}

	//slice := []int{0, 1, 2, 3}
	//for _, val := range slice {
	//	val++
	//}
	//fmt.Println(slice)
	//
	//arr := [3]int{1, 2, 3}
	//for _, val := range arr {
	//	val++
	//}
	//fmt.Println(arr)

	//slice := []int{0, 1, 2, 3}
	//m := make(map[int]*int)
	//for key, _ := range slice {
	//	m[key] = &slice[key]
	//}
	//////0 -> 0
	//////1 -> 1
	//////2 -> 2
	//////3 -> 3
	//for k, v := range m {
	//	fmt.Println(k, "->", *v)
	//}

	var ap *[3]int

	////会导致panic
	for i, p := range ap {
		fmt.Println(i, p)
	}

	////舍弃for range的第二个值
	////不会导致panic
	//for i, _ := range ap {
	//	fmt.Println(i)
	//}
}
