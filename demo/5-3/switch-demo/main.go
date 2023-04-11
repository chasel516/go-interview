package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//switch false {
	//case false:
	//	fmt.Println(1)
	//case false:
	//	fmt.Println(2)
	//}

	//switch n := rand.Intn(10); n {
	//case 1, 2, 3:
	//	fmt.Println(1, 2, 3)
	//case 3, 4, 5:
	//	fmt.Println(3, 4, 5)
	//
	//}

	//switch n := rand.Intn(10); n {
	//case n > 1:
	//	fmt.Println(1, 2, 3)
	//}
	//
	//switch n := rand.Intn(10); n {
	//case int32(1):
	//	fmt.Println(1, 2, 3)
	//case int64(2):
	//
	//}

	//rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要
	//switch n := rand.Intn(100) % 5; n {
	//case 0, 1, 2, 3, 4:
	//	fmt.Println("n =", n)
	//	fallthrough // 跳到下个case分支
	//case 5, 6, 7, 8:
	//	fmt.Println("n =", n) // 99
	//	fallthrough           // 跳到下个case分支
	//default:
	//	fmt.Println("n =", n)
	//}

	//rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要
	switch n := rand.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		fallthrough // 跳到下个case分支
	case 5, 6, 7, 8:
		// 一个新声明的n，它只在当前分支代码块内可见。
		n := 999
		fmt.Println("n =", n) // 999
		fallthrough           //跳到下个case分支
	default:
		// 这里的n和第一个分支中的n是同一个变量，它们均为switch表达式"n"。
		fmt.Println("n =", n)
	}

	switch n := rand.Intn(4); n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		// 此整个if代码块为当前分支中的最后一条语句
		if true {
			// 编译失败: 不是当前分支中的最后一条语句
			fallthrough
		}
	case 5, 6, 7, 8:
		// 编译失败: 不是当前分支中的最后一条语句
		fallthrough
		fmt.Println(n)
	default:
		fmt.Println(n)
		fallthrough // error: 不能出现在最后一个分支中
	}

}
