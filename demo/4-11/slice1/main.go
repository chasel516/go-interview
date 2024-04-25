package main

import "fmt"

func main() {
	var array [10]int

	var slice = array[7:8] //分配的cap基于array的len-区间开始

	fmt.Println("lenth of slice:", len(slice))
	fmt.Println("capacity of slice:", cap(slice))
	fmt.Println(&slice[0] == &array[7])
}
