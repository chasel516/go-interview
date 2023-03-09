package main

import "fmt"

func main() {

	s := []string{"a", "b", "c"}
	fmt.Println("s:origin", s)
	f1(s)
	fmt.Println("s:f1", s)

	f2(s)
	fmt.Println("s:f2", s)

	f3(s)
	fmt.Println("s:f3", s)

}

func f1(s []string) {
	var tmp = []string{"x", "y", "z"}
	s = tmp
}

func f2(s []string) {
	// item只是一个副本，不能改变s中元素的值
	for i, item := range s {
		item = "d"
		fmt.Printf("item=%s;s[%d]=%s", item, i, s[i])
	}
}

func f3(s []string) {
	for i := range s {
		s[i] = "d"
	}
}
