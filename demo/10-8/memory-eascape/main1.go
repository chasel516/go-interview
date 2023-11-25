package main

var g1 = 10
var g2 *int
var g3 = &g1
var g4 = "global"
var g5 = []int{1, 2, 3}
var g6 = map[int]int{1: 1} //发生了逃逸
var g7 map[int]int         //只定义则不会发生逃逸
var g8 = make(chan int)
var g9 = [3]int{}

type st struct {
	Name string
	Tag  map[int]int
}

var g10 = st{
	Name: "test",
	Tag:  nil,
}

func main() {
	_ = g3
	_ = g6
	m := map[int]int{1: 1} //map作为局部变量没有发生逃逸
	_ = m
}
func f() (int, *int, string, []int, map[int]int, chan int, [3]int, st) {
	g8 <- g1
	go func() { //func逃逸了，但变量g8没有
		<-g8
	}()
	return g1, g2, g4, g5, g7, g8, g9, g10
}
