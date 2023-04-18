package main

import "fmt"

type Duck interface {
	Quack()
}

type YellowDuck struct{}

func (yd YellowDuck) Quack() {
	fmt.Println("YellowDuck嘎嘎叫")
}

type NormalDuck struct{}

func (nd NormalDuck) Quack() {
	fmt.Println("NormalDuck嘎嘎叫")
}

func Quack(d Duck) {
	d.Quack()
}
func main() {
	ducks := []Duck{YellowDuck{}, NormalDuck{}}
	for _, duck := range ducks {
		//duck.Quack()
		Quack(duck)
	}
}
