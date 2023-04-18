package main

import "fmt"

type Iperson interface {
	Talk()
	Walk()
}
type Person struct {
	Iperson
}

func (p *Person) Talk() {
	fmt.Println("talk")
}

//func (p *Person) Walk() {
//	fmt.Println("walk")
//}

func Talk(p Iperson) {
	p.Talk()
}

type Teacher struct {
}

func (t *Teacher) Talk() {
	fmt.Println("Teacher talk")
}

func (t *Teacher) Walk() {
	fmt.Println("Teacher walk")
}
func main() {
	p := &Person{}
	var i Iperson
	i = p
	fmt.Println(i)
	p.Talk()
	p.Walk()
	t := &Teacher{}
	Talk(p)
	Talk(t)

}
