package main

func main() {
	// var m map[string]int
	// var n map[string]int
	// m1 := m
	// fmt.Println(m == nil) //true
	// fmt.Println(n == nil) //true
	// // 不能通过编译
	// // fmt.Println(m == n) //map can only be compared to nil
	// fmt.Println(reflect.ValueOf(m).Pointer() == reflect.ValueOf(n).Pointer())
	// fmt.Println(reflect.ValueOf(m).Pointer() == reflect.ValueOf(m1).Pointer())
	// fmt.Println(reflect.DeepEqual(m, n))
	// fmt.Println(reflect.DeepEqual(m, m1))

	// m = make(map[string]int, 10)
	// n = make(map[string]int, 100)
	// for s, i := range m {
	// 	fmt.Println(s, i)
	// }
	// m["a"] = 1
	// n["a"] = 1

	// fmt.Println(reflect.DeepEqual(m, n)) //true

	// type person struct {
	// 	name string
	// }
	// p1 := person{"test"}
	// p2 := struct {
	// 	name string
	// }{"test"}
	// // fmt.Println(p1)
	// // fmt.Println(p2)
	// fmt.Println(p1 == p2)                  //true
	// fmt.Println(reflect.DeepEqual(p1, p2)) //false

	// x := new(int)
	// y := new(int)
	// fmt.Println("x==y:", x == y)                             //x==y: false
	// fmt.Println("DeepEqual(x, y):", reflect.DeepEqual(x, y)) //DeepEqual(x, y): true

	// type link struct {
	// 	data interface{}
	// 	next *link
	// }
	// var a, b, c link
	// a.next = &b
	// b.next = &c
	// c.next = &a

	// fmt.Println(a == b)                  //false
	// fmt.Println(reflect.DeepEqual(a, b)) //true

	// x := 3.14
	// y := 2.71
	// if math.Abs(x-y) < 0.0001 {
	// 	fmt.Println("x and y are equal")
	// } else {
	// 	fmt.Println("x and y are not equal")
	// }

	// type Person struct {
	// 	Name string
	// 	Age  int
	// }

	// p1 := Person{Name: "Alice", Age: 25}
	// p2 := Person{Name: "Alice", Age: 25}

	// equal := p1 == p2
	// if equal {
	// 	fmt.Println("两个对象相等")
	// } else {
	// 	fmt.Println("两个对象不相等")
	// }

}

//基本类型、切片、map、结构体

//如果是不同的类型，即使是底层类型相同，相应的值也相同，那么两者也不是“深度”相等。

//其实== 运算符对于指针类型的比较并不比较它们的内容，而是比较它们的引用地址。
