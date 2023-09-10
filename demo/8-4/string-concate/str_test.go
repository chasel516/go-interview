package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

// 使用 + 拼接字符串
func BenchmarkStringPlus(b *testing.B) {
	result := ""
	s := "Hello, World!"
	for i := 0; i < b.N; i++ {
		result += s
	}
	_ = result
}

func BenchmarkStringPlus2(b *testing.B) {
	result := ""
	const s = "Hello, World!"
	for i := 0; i < b.N; i++ {
		result += s
	}
	_ = result
}

func BenchmarkStringPlus3(b *testing.B) {
	result := ""
	const s1 = "Hello, "
	const s2 = "World!"
	for i := 0; i < b.N; i++ {
		result = s1 + s2
	}
	_ = result
}

func BenchmarkStringPlus4(b *testing.B) {
	result := ""
	s1 := "Hello, "
	s2 := "World!"
	for i := 0; i < b.N; i++ {
		result = s1 + s2
	}
	_ = result
}

// 使用 fmt.Sprintf 拼接字符串
func BenchmarkStringFmtSprintf(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%s%s", result, "Hello, World!")
	}
	_ = result
}

// 使用 buffer.WriteString 拼接字符串
func BenchmarkStringBufferWrite(b *testing.B) {
	var buffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		buffer.WriteString("Hello, World!")
	}
	result := buffer.String()
	_ = result
}

func BenchmarkStringBufferWriteWithGrow(b *testing.B) {
	var buffer bytes.Buffer
	buffer.Grow(b.N * len("Hello, World!"))
	for i := 0; i < b.N; i++ {
		buffer.WriteString("Hello, World!")
	}
	result := buffer.String()
	_ = result
}

// 使用 strings.Builder 拼接字符串
func BenchmarkStringsBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, World!")
	}
	result := builder.String()
	_ = result
}

func BenchmarkStringsBuilderWithGrow(b *testing.B) {
	var builder strings.Builder
	builder.Grow(b.N * len("Hello, World!"))
	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, World!")
	}
	result := builder.String()
	_ = result
}

// 使用 []byte 拼接字符串
func BenchmarkStringByte(b *testing.B) {
	bt := make([]byte, 0, b.N*len("Hello, World!"))
	for i := 0; i < b.N; i++ {
		bt = append(bt, "Hello, World!"...)
	}
	result := string(bt)
	_ = result
}

func BenchmarkStringByteWithPre(b *testing.B) {
	bt := make([]byte, 0, b.N*len("Hello, World!"))
	for i := 0; i < b.N; i++ {
		bt = append(bt, "Hello, World!"...)
	}
	result := string(bt)
	_ = result
}

func BenchmarkStringByteWithPreAndZeroCopy(b *testing.B) {
	bt := make([]byte, 0, b.N*len("Hello, World!"))
	for i := 0; i < b.N; i++ {
		bt = append(bt, "Hello, World!"...)
	}
	//string(bt)
	result := *(*string)(unsafe.Pointer(&bt))
	_ = result
}

// 使用 strings.Join 拼接字符串
func BenchmarkStringsJoin(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		parts := []string{"Hello,", "World!"}
		result = strings.Join(parts, " ")
	}
	_ = result
}
