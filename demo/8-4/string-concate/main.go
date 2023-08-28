package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	//连接符 +
	s1 := "ab"
	s2 := "cd"
	fmt.Println(s1 + s2)
	//字符串格式化函数fmt.Sprintf
	s := fmt.Sprintf("%s%s", s1, s2)
	fmt.Println("fmt.Sprintf:", s)

	//bytes.Buffer
	var buf bytes.Buffer
	buf.WriteString(s1)
	buf.WriteString(s2)
	fmt.Println("bytes.Buffer:", buf.String())

	//strings.Builder
	var bu strings.Builder
	bu.WriteString(s1)
	bu.WriteString(s2)
	fmt.Println("strings.Buffer:", bu.String())

	//预分配[]byte
	bt := make([]byte, 0, len(s1)+len(s2))
	bt = append(bt, s1...)
	bt = append(bt, s2...)
	fmt.Println("预分配[]byte:", string(bt))

	//strings.join
	sl := []string{"ab", "cd"}
	str := strings.Join(sl, "")
	fmt.Println("strings.Join:", str)
}
