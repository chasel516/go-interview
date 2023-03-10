package main

import "fmt"

func main() {
	//布尔类型
	//ok := true
	//fmt.Printf("%s,%t \n", ok, ok)

	//整数类型
	fmt.Printf("%T, %d \n", 123456789, 123456789)   //int, 123456789
	fmt.Printf("%T, %5d \n", 123456789, 123456789)  //int, 123456789
	fmt.Printf("%T, %05d \n", 123456789, 123456789) //int, 123456789
	fmt.Printf("%T, %b \n", 123456789, 123456789)   //int, 111010110111100110100010101
	fmt.Printf("%T, %o \n", 123456789, 123456789)   //int, 726746425
	fmt.Printf("%T, %c \n", 66, 66)                 //int, B
	fmt.Printf("%T, %q \n", 66, 66)                 //int, 'B'
	fmt.Printf("%T, %x \n", 123456789, 123456789)
	fmt.Printf("%T, %X \n", 123456789, 123456789)
	fmt.Printf("%T, %U \n", '中', '中')

	//浮点型
	//fmt.Printf("%b \n", 1000.123456789) //8797178959608267p-43
	//fmt.Printf("%f \n", 1000.1)
	//fmt.Printf("%f\n", float64(1000)) //1000.000000
	//fmt.Printf("%.2f \n", 1000.123456789)
	//fmt.Printf("%e\n", 1000.1234567898)   //1.000123e+03
	//fmt.Printf("%.5e\n", 1000.1234567898) //1.00012e+03
	//fmt.Printf("%.E\n", 1000.1234567898)  //1E+03
	//fmt.Printf("%.5E\n", 1000.1234567898) //1.00012E+03
	//fmt.Printf("%.5E\n", float64(1000))   //1.00000E+03
	//fmt.Printf("%F \n", 1000.123456789)
	//fmt.Printf("%g \n", 1000.123456789)
	//fmt.Printf("%G \n", 1000.123456789)
}
