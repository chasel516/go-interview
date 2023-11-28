package main

func main() {}

func fn() {
	x := 1
	y := 2
	go func() {
		x += 1
	}()
	y += 1
}
