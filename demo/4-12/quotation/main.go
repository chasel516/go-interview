package main

import (
	"fmt"
	"reflect"
)

func main() {
	//str1 := "\141\142\143"
	//fmt.Println(str1)
	//str2 := "\141\142\143\42"
	//fmt.Println(str2)
	//str3 := "\141\142\143\47"
	//fmt.Println(str3)

	// 单引号用来定义一个 byte或者rune,默认是rune
	//char1 := 'a'
	//var char2 byte = 'a'
	//char3 := '中'
	//fmt.Println("char1:", char1, reflect.TypeOf(char1))
	//fmt.Println("char2:", char2, reflect.TypeOf(char2))
	//fmt.Println("char3:", char3, reflect.TypeOf(char3))

	//s1 := "123\t45\"a\"6" //\t和\"将被转义
	//s2 := `123\t45\"a\"6` //\t和\"将被原样输出
	//fmt.Println("s1:", s1)
	//fmt.Println("s2:", s2)

	//a,b,c的ASCII码值的十进制分别是97,98,99,对应的8进制为141，142，143
	//对于字符串abc我们可以通过
	s3 := "\141\142\143"
	fmt.Println("s3:", s3, reflect.TypeOf(s3))

	//a,b,c的ASCII码值的十进制分别是97,98,99,对应的16进制为61，62，63
	s4 := "\x61\x62\x63"
	fmt.Println("s4:", s4, reflect.TypeOf(s4))

	//unicode也是通过16进制的码值表示的
	s5 := "\u0061\u0062\u0063"
	fmt.Println("s5:", s5, reflect.TypeOf(s5))
	s6 := "\U00000061\U00000062\U00000063"
	fmt.Println("s6:", s6, reflect.TypeOf(s6))

	//单引号ASCII码值的八进制是47，十六进制是27
	s7 := "'"
	fmt.Println("s6:", s6, reflect.TypeOf(s7))
	s8 := "\'" //编译不通过

	s9 := ""
	" //编译不通过
	s10 := "\"" //需要加上转义符方可成为合法的字符串
	//fmt.Println("s10:", s10, reflect.TypeOf(s10))

	s11 := "\47"     //编译不通过
	s12 := "erer\42" //编译不通过

	s12 := "\x22"
	//fmt.Println("s12:", s12)
	s13 := "\u0027"
	//fmt.Println("s13:", s13)
	s14 := "\U00000027"
	//fmt.Println("s14:", s14)

}
