package main

import "fmt"

type Animal struct {
}

func (a *Animal) eat() {
	fmt.Println("Animal is eating")
}

// Cat继承Animal
type Cat struct {
	Animal
}

// Cat子类也可以有eat方法，且实现可以跟父类Animal不同
func (c *Cat) eat() {
	fmt.Println("Cat is eating fish")
}

func main() {
	a := &Animal{}
	c := &Cat{}
	a.eat()
	c.eat()
}
