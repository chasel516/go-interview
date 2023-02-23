package main

import "fmt"

func main() {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	//c := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	if &a == &b {
		fmt.Println("equal")
	} else {
		fmt.Printf("not equal a:%p,b:%p \n", &a, &b)
	}
}
