package main

import "fmt"

var (
	appName string
	version string
)

func main() {
	//布尔类型
	//ok := true
	//fmt.Printf("%T,%t \n", ok, ok)

	//var r rune = 65
	////整数类型
	//fmt.Printf("%T, %d \n", 123456789, 123456789)   //int, 123456789
	//fmt.Printf("%T, %5d \n", 123456789, 123456789)  //int, 123456789
	//fmt.Printf("%T, %5d \n", 12, 12)                //int,    12
	//fmt.Printf("%T, %05d \n", 123456789, 123456789) //int, 123456789
	//fmt.Printf("%T, %05d \n", 12, 12)               //int, 00012
	//fmt.Printf("%T, %b \n", 123456789, 123456789)   //int, 111010110111100110100010101
	//fmt.Printf("%T, %o \n", 123456789, 123456789)   //int, 726746425
	//fmt.Printf("%T, %c \n", 66, 66)                 //int, B
	//fmt.Printf("%T, %q \n", 66, 66)                 //int, 'B'
	//fmt.Printf("%T, %x \n", 123456789, 123456789)   //int, 75bcd15
	//fmt.Printf("%T, %X \n", 123456789, 123456789)   //int, 75BCD15
	//fmt.Printf("%T, %U \n", '中', '中')               //int32, U+4E2D //字符的字面量是rune类型
	//fmt.Printf("%T, %v ,%s \n", r, r, string(r))    //int32, 65 ,A
	//fmt.Printf("%T, %c \n", r, r)                   //int32 ,A

	//var s = "面试"
	//fmt.Println(s[0])
	//for i, item := range s {
	//	//0 : A
	//	//1 : B
	//	fmt.Println(i, ":", item)
	//	fmt.Printf("i:%c\n", item)
	//}

	//浮点型
	//fmt.Printf("%b \n", 1000.123456789)   //8797178959608267p-43
	//fmt.Printf("%f \n", 1000.123456789)   //1000.123457
	//fmt.Printf("%f\n", 1000.0)            //1000.000000
	//fmt.Printf("%.2f \n", 1000.123456789) //1000.12
	//fmt.Printf("%.2f \n", 1000.125)       //1000.12
	//fmt.Printf("%.2f \n", 1000.126)       //1000.13
	//fmt.Printf("%e\n", 1000.1234567898)   //1.000123e+03
	//fmt.Printf("%.5e\n", 1000.1234567898) //1.00012e+03
	//fmt.Printf("%E\n", 1000.1234567898)   //1E+03
	//fmt.Printf("%.5E\n", 1000.1234567898) //1.00012E+03
	//fmt.Printf("%F \n", 1000.123456789)   //1000.123457
	//fmt.Printf("%g \n", 1000.123456789)   //1000.123456789
	//fmt.Printf("%G \n", 1000.123456789)   //1000.123456789

	//字符串
	//arr := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//arr1 := []byte{103, 111, 108, 97, 110, 103}
	//arr2 := []uint8{103, 111, 108, 97, 110, 103}
	//arr3 := []rune{'g', 'o', 'l', 'a', 'n', 'g'}
	//arr4 := []int32{'g', 'o', 'l', 'a', 'n', 'g'}
	//
	//fmt.Printf("%s \n", "go面试专题")       //go面试专题
	//fmt.Printf("%q \n", "go面试专题")       //"go面试专题"
	//fmt.Printf("%x \n", "go面试专题")       //676fe99da2e8af95e4b893e9a298
	//fmt.Printf("%X \n", "go面试专题")       //676FE99DA2E8AF95E4B893E9A298
	//fmt.Printf("%T, %s \n", arr, arr)   //[]uint8, golang
	//fmt.Printf("%T, %q \n", arr, arr)   //[]uint8, "golang"
	//fmt.Printf("%T, %x \n", arr, arr)   //[]uint8, 676f6c616e67
	//fmt.Printf("%T, %X \n", arr, arr)   //[]uint8, 676F6C616E67
	//fmt.Printf("%T, %s \n", arr1, arr1) //[]uint8, golang
	//fmt.Printf("%T, %s \n", arr2, arr2) //[]uint8, golang
	//fmt.Printf("%T, %c \n", arr3, arr3) //[]int32, [g o l a n g]
	//fmt.Printf("%T, %c \n", arr4, arr4) //[]int32, [g o l a n g]

	//指针
	//var name *string
	//tmp := "go面试"
	//name = &tmp
	//fmt.Printf("%p \n", name) //0xc00009e220

	//Scan
	//var (
	//	appName string
	//	version int
	//)
	//fmt.Println("请输入name:")
	//fmt.Scan(&appName)
	//fmt.Println("请输入version")
	//fmt.Scan(&version)
	//fmt.Printf("fmt.Scan appName:%s version:%d \n", appName, version)

	//Scanf
	//_, err := fmt.Scanf("name=%s ver=%s", &appName, &version)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("fmt.Scanf appName:%s version:%s \n", appName, version)

	//Scanln
	//fmt.Println("请输入name")
	//fmt.Scanln(&appName)
	//fmt.Println("请输入version")
	//fmt.Scanln(&version)
	//fmt.Printf("fmt.Scan appName:%s version:%s \n", appName, version)

	//reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	//fmt.Println("请输入：")
	//text, _ := reader.ReadString('\n') // 读到换行
	////text, _ := reader.ReadString(' ') // 读到换行
	//fmt.Println(text)

	//go run -ldflags "-X 'main.appName=test' -X 'main.version=1'" main_test.go
	//go build -ldflags "-X 'main.appName=test' -X 'main.version=1'" main_test.go
	fmt.Printf("appName:%s version:%s \n", appName, version)
}
