package main

import "fmt"

func main() {

	//错误写法
	// slice := []int{0, 1, 2, 3}
	// m := make(map[int]*int)
	// for key, val := range slice {
	// 	m[key] = &val
	// }

	// //0 -> 3
	// //1 -> 3
	// //2 -> 3
	// //3 -> 3
	// for k, v := range m {
	// 	fmt.Println(k, "->", *v)
	// }

	////正确写法
	// slice := []int{0, 1, 2, 3}
	// m := make(map[int]*int)
	// for key, val := range slice {
	// 	value := val  //显示绑定，每次循环都会生成一个新的变量
	// 	m[key] = &value
	// }

	// //0 -> 0
	// //1 -> 1
	// //2 -> 2
	// //3 -> 3
	// for k, v := range m {
	// 	fmt.Println(k, "->", *v)
	// }

	// slice := []int{0, 1, 2, 3}
	// for _, val := range slice {
	// 	val++
	// }
	// fmt.Println(slice)

	// arr := [3]int{1, 2, 3}
	// for _, val := range arr {
	// 	val++
	// }
	// fmt.Println(arr)

	// slice := []int{0, 1, 2, 3}
	// m := make(map[int]*int)
	// for key, _ := range slice {
	// 	m[key] = &slice[key]
	// }
	// //////0 -> 0
	// //////1 -> 1
	// //////2 -> 2
	// //////3 -> 3
	// for k, v := range m {
	// 	fmt.Println(k, "->", *v)
	// }

	// var ap *[3]int

	// ////会导致panic
	// for i, p := range ap {
	// 	fmt.Println(i, p)
	// }

	//舍弃for range的第二个值
	//不会导致panic
	// for i, _ := range ap {
	// 	fmt.Println(i)
	// }

	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	for _, j := range v {
		fmt.Println(j)
	}
}

/*
如果循环体中会启动协程（并且协程会使用循环变量），就需要格外注意了，因为很可能循环结束后协程才开始执行， 此时，所有协程使用的循环变量有可能已被改写。（是否会改写取决于引用循环变量的方式）
*/

/**
如果循环体没有并发出现，则引用循环变量一般不会出现问题；
如果循环体有并发，则根据引用循环变量的位置不同而有所区别
通过参数完成绑定，则一般没有问题；
函数体中引用，则需要显式地绑定


range 是迭代遍历手段，可以操作的类型有数组、切片、map、channel等

优化：
slice
遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值。
map
函数中for-range语句中只获取key值，然后根据key值获取value值，虽然看似减少了一次赋值，但通过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行。

channel遍历是依次从channel中读取数据,读取前是不知道里面有多少个元素的。如果channel中没有元素，则会阻塞等待，如果channel已被关闭，则会解除阻塞并退出循环。

注意：
遍历过程中可以视情况放弃接收index或value，可以一定程度上提升性能
遍历channel时，如果channel中没有数据，可能会阻塞
尽量避免遍历过程中修改原数据
*/
