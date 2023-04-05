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
	yd := YellowDuck{}
	nd := NormalDuck{}
	Quack(yd)
	Quack(nd)

}
