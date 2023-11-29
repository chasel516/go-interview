package main

import (
	"log"
	"reflect"
	"strings"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	s := strings.Repeat("字", 10)
	ptr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	log.Println("s :", unsafe.Pointer(ptr.Data))

	//转字节切片(发生拷贝)
	b := []byte(s)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&b))
	log.Println("b:", unsafe.Pointer(ptr.Data))

	//强转(发生拷贝)
	s2 := string(b)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s2))
	log.Println("s2:", unsafe.Pointer(ptr.Data))

	//字节切片到字符串的零拷贝转换
	s3 := bytesToString(b)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&s3))
	//s4和b的地址相同
	log.Println("s3:", unsafe.Pointer(ptr.Data))

	//字符串到字节切片的零拷贝转换
	b1 := stringToBytes(s)
	ptr = (*reflect.StringHeader)(unsafe.Pointer(&b1))
	//b1和s的地址相同
	log.Println("b1:", unsafe.Pointer(ptr.Data))

	log.Println("s3:", s3)
}

func stringToBytes(s string) []byte {
	// 将字符串头信息转换为 reflect.StringHeader 类型，以获取字符串底层数据的地址和长度
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	// 创建一个 reflect.SliceHeader 类型的结构体，用于构建字节切片头信息
	sliceHeader := &reflect.SliceHeader{
		Data: stringHeader.Data, // 字符串底层数据的地址
		Len:  stringHeader.Len,  // 字符串的长度
		Cap:  stringHeader.Len,  // 字节切片的容量，与字符串长度相同
	}
	//利用 unsafe.Pointer 将 reflect.SliceHeader 转换为 []byte 类型的指针，
	//然后再通过 * 运算符解引用得到字节切片
	return *(*[]byte)(unsafe.Pointer(sliceHeader))
}

func bytesToString(b []byte) string {
	//将字节切片头信息转换为 reflect.SliceHeader 类型，以获取字节切片底层数据的地址和长度
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	// 创建一个 reflect.StringHeader 类型的结构体，用于构建字符串头信息
	stringHeader := &reflect.StringHeader{
		Data: sliceHeader.Data, // 字节切片底层数据的地址
		Len:  sliceHeader.Len,  // 字节切片的长度
	}

	// 利用 unsafe.Pointer 将 reflect.StringHeader 转换为 string 类型的指针，
	//然后再通过 * 运算符解引用得到字符串
	return *(*string)(unsafe.Pointer(stringHeader))
}
