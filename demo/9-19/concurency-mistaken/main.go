package main

var counter int

func increase() {
	add()
}

func add() {
	plus()
}

func plus() {
	counter++
}

func main() {
	for i := 0; i < 10; i++ {
		go increase()
	}
}
