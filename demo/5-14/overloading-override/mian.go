package main

import "fmt"

type Calculator struct {
}

func (c Calculator) add(x, y int) int {
	return x + y
}

func (c Calculator) add(x, y, z int) int {
	return x + y + z
}

func main() {
	c := Calculator{}
	fmt.Println(c.add(1, 2))
	fmt.Println(c.add(1, 2, 3))
}
