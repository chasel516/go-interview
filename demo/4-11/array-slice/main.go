package main

import "fmt"

func main() {
	//数组切片和映射在go语言中也叫容器类型
	//数组和切片键值的合法取值范围为左闭右开区间[0, 此数组或切片的长度)

	//一个数组或者切片的所有元素紧挨着存放在一块连续的内存中。一个数组中的所有元素均存放在
	//此数组值的直接部分，一个切片中的所有元素均存放在此切片值的间接部分。 在官方标准编译器
	//和运行时中，映射是使用哈希表算法来实现的。所以一个映射中的所有元素也均存放在一块连续
	//的内存中，但是映射中的元素并不一定紧挨着存放。 另外一种常用的映射实现算法是二叉树算
	//法。无论使用何种算法，一个映射中的所有元素的键值也存放在此映射值（的间接部分）中

	//对于数组切片和映射这三种数据类型，通过元素的键值来访问此元素的时间复杂度均为O(1)
	//但是一般来说，映射元素访问消耗的时长要数倍于数组和切片元素访问消耗的时长
	//是映射相对于数组和切片有两个优点：
	//映射的键值类型可以是任何可比较类型。
	//对于大多数元素为零值的情况，使用映射可以节省大量的内存。

	//数组的零值和切片的零值不同
	//一个数组的零值是数组的所有元素均为对应数组元素类型的零值
	var f byte
	fmt.Println("yyyyyy", f)
	arr1 := [10]int{}
	fmt.Println("arr1:", arr1)

	var arr2 [10]int
	fmt.Println("arr2:", arr2)

	//有零值
	arr3 := new([10]int)
	fmt.Println("arr3:", arr3)

	//为nil,无零值，直接使用会报错
	var arr4 *[10]int
	fmt.Println("arr4:", arr4)
	fmt.Println(arr3 == arr4)
	arr1[0] = 1
	arr2[0] = 1
	//a4[0] = 1  //空指针异常
	var s1 = []int{}
	fmt.Println("s1:", s1)
	var s2 = new([]int)
	fmt.Println("s2:", s2) //为nil,但可以直接使用
	fmt.Println("s2 is nil?:", s2 == nil)
	var s3 = make([]int, 1)
	fmt.Println("s3:", s3)

	var s4 []int
	fmt.Println("s4:", s4)
	fmt.Println("s4 is nil?:", s4 == nil)
	s4 = append(s4, 1)
	var s5 *[]int //为nil，直接使用会报空指针
	tmp := []int{}
	s5 = &tmp
	tmp1 := append(*s5, 1)
	s5 = &tmp1
	fmt.Println("s5:", s5)

	//总结：
	//1. 非指针类型的数组通过字面量或者new声明后，数组的每个元素都会被赋上元素类型的零值，也就是说数组的元素所占的内存空间在声明的时候就会被开辟出来
	//2. 非指针切片通过字面量或者new声明后，切片元素不会被赋上元素类型的零值，需要使用make指定size后才会被赋上元素类型的零值
	//3. 指针类型的数组和切片直接用类型声明后是nil，不能直接使用

	//切片的字面量有下面几种表示方式：
	//s1 := []string{"a", "b", "c"}
	//fmt.Println("s1:", s1)
	//s2 := []string{0: "a", 1: "b", 2: "c"}
	//fmt.Println("s2:", s2)
	//s3 := []string{2: "c", 1: "b", 0: "a"}
	//fmt.Println("s3:", s3)
	//s4 := []string{2: "c", 0: "a", "b"}
	//fmt.Println("s4:", s4)

	//数组和切片的传值方式不同
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1
	//slice1和slice2指向了相同的地址
	fmt.Printf("slice1 ptr:%p;slice2 ptr:%p\n", slice1, slice2)

	array1 := [4]int{1, 2, 3, 4}
	array2 := array1
	//array1和array2指向了不同的地址
	fmt.Printf("array1 ptr:%p;array2 ptr:%p\n", &array1, &array2)

	//1.通过make申请
	var s6 = make([]int, 100)

	//2.初始化第100个元素
	var s7 = []int{99: 0}

	//3.先创建一个容量为100的数组取其地址后再截取全部元素作为切片
	var s8 = (&[100]int{})[:]

	//4.通过new创建一个容量为100的数组，再取全部元素
	var s9 = new([100]int)[:]
	// 100 100 100 100
	println(len(s6), len(s7), len(s8), len(s9)) //100 100 100 100

	//数组的字面量有下面几种表示方式：
	a1 := [4]string{"a", "b", "c", "d"}
	fmt.Println("a1:", a1)

	a2 := [4]string{0: "a", 1: "b", 2: "c", 3: "d"}
	fmt.Println("a2:", a2)

	a3 := [4]string{1: "b", "a", "c"}
	fmt.Println("a3:", a3)

	//如果不指定下标，将会从前一个下标开始递增
	//a4 := [4]string{1: "b", "a", "c", "d"} //编译错误

	//这里元素c的下标也是1，而数组和切片的下标是不能重复的，即元素键值不能重复，但元素值可以重复
	//a5 := [4]string{1: "b", 0: "a", "c", "d"}//编译错误

	a6 := [...]string{"a", "b", "c", "d"}

	fmt.Println("a6:", a6)
	a7 := [...]string{3: "d", 1: "b", "a"}

	fmt.Println("a7:", a7)
	//编译错误
	//a8 := [...]string{3: "d", 1: "b", "a", "c"}//编译错误

	//大于等于0的常量整型可以作为数组和切片的小标，但变量不行
	const x uint = 1
	//var y uint = 1
	a8 := []string{0: "a", x: "b", 2: "c"}
	fmt.Println("a8:", a8)
	//a9 := []string{0: "a", y: "b", 2: "c"} 	//编译错误

	//数组的比较
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	//c := [3]int{5, 6}

	//类型相同，值相同
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	//类型相同，值相同，但地址不同
	if &a == &b {
		fmt.Println("equal")
	} else {
		fmt.Printf("not equal a:%p,b:%p \n", &a, &b)
	}

	//无法比较，编译不过。下标作为数组类型的一部分
	//if a == c {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
}
