package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//1.连接符 +
	s1 := "ab"
	s2 := "cd"
	//fmt.Println("连接符 +:", s1+s2)
	//
	////2.字符串格式化函数fmt.Sprintf
	//s := fmt.Sprintf("%s%s", s1, s2)
	//fmt.Println("fmt.Sprintf:", s)
	//
	////3.预分配bytes.Buffer
	var buf bytes.Buffer
	buf.Grow(len(s1) + len(s2))
	buf.WriteString(s1)
	buf.WriteString(s2)
	fmt.Println("bytes.Buffer:", buf.String())
	BufferCapResize()

	//4.预分配strings.Builder
	var bu strings.Builder
	bu.Grow(len(s1) + len(s2))
	bu.WriteString(s1)
	bu.WriteString(s2)
	fmt.Println("strings.Buffer:", bu.String())
	BuilderCapResize()
	//
	//5.预分配[]byte
	bt := make([]byte, 0, len(s1)+len(s2))
	bt = append(bt, s1...)
	bt = append(bt, s2...)
	fmt.Println("预分配[]byte:", string(bt))

	//6.strings.join
	sl := []string{"ab", "cd"}
	str := strings.Join(sl, "")
	fmt.Println("strings.Join:", str)
}
func BufferCapResize() {
	var str = "abcd"
	var buf bytes.Buffer
	buf.Grow(4 * 10000)
	cap := 0
	for i := 0; i < 10000; i++ {
		if buf.Cap() != cap {
			fmt.Println(buf.Cap())
			cap = buf.Cap()
		}
		buf.WriteString(str)
	}
}

func BuilderCapResize() {
	var str = "abcd"
	var builder strings.Builder
	cap := 0
	for i := 0; i < 10000; i++ {
		if builder.Cap() != cap {
			fmt.Println(builder.Cap())
			cap = builder.Cap()
		}
		builder.WriteString(str)

	}
}
