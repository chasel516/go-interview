package main

import (
	"fmt"
)

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}

func main() {
	var slice []int
	slice = append(slice, 1, 2, 3)
	fmt.Println(len(slice), cap(slice)) //3 3

	newSlice := AddElement(slice, 4)
	fmt.Println(len(newSlice), cap(newSlice)) //4 6
	fmt.Println(&slice[0])
	fmt.Println(&newSlice[0])
	fmt.Println(&slice[0] == &newSlice[0]) //false
}

//https://www.jianshu.com/p/db045aeba371
