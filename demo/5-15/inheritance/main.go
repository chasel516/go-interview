package main

import "fmt"

type People struct{}

func (p *People) Run() {
	fmt.Println("Run")
	p.Eat()
}
func (p *People) Eat() {
	fmt.Println("Eat")
}

type Person struct{}

func (p *Person) Run() {
	fmt.Println("Person Run")
	p.Eat()
}
func (p *Person) Eat() {
	fmt.Println("Person Eat")
}

type Teacher struct {
	//People
	*People
	*Person
}

func (t *Teacher) Speck() {
	fmt.Println("teacher Speck")
}

func (t *Teacher) Run() {
	fmt.Println("teacher Eat")
}

func (t *Teacher) Eat() {
	fmt.Println("teacher Eat")
}

func main() {
	t := Teacher{}
	t.Run()
	t.Eat()
	t.Speck()
}
