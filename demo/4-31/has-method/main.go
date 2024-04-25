package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
}

func (p person) Say(str string) string {
	return p.Name + str
}

func main() {
	//方法1：通过类型转化判断
	p := &person{}
	fmt.Println(HasMethodSay(p)) //true

	//方法2；通过反射实现
	fmt.Println(HasMethod(p, "Say")) //true

	//对象x转换成反射对象v，v又通过Interface()接口转换成interface对象，interface对象通过.(float64)类型断言获取float64类型的值。
	var x float64 = 3.4
	v := reflect.ValueOf(x) //v is reflect.Value
	// v.Elem().SetFloat(7.1)  //反射对象可修改，value值必须是可设置的
	var y float64 = v.Interface().(float64)
	fmt.Println("value:", y)

}

func HasMethodSay(v interface{}) (has bool) {
	_, has = v.(interface{ Say(str string) string })
	return
}

func HasMethod(obj interface{}, methodName string) bool {
	if methodName == "" {
		return false
	}
	object := reflect.ValueOf(obj)
	// 获取到方法
	method := object.MethodByName(methodName)
	if !method.IsValid() {
		return false
	}

	return true
}

//反射提供一种让程序检查自身结构的能力
/*

reflect.Type 提供一组接口处理interface的类型，即（value, type）中的type
reflect.Value提供一组接口处理interface的值,即(value, type)中的value

*/

// go存在4种类型转换：断言、强制、显示、隐式
/*
断言
var s = x.(T)
如果x不是nil，且x可以转换成T类型，就会断言成功，返回T类型的变量s。
如果T不是接口类型，则要求x的类型就是T
如果T是一个接口，要求x实现了T接口

如果断言类型成立，则表达式返回值就是T类型的x，如果断言失败就会触发panic


go提供了另外一种带返回是否成立的断言语法
s, ok:= x.(T)
该方法和第一种差不多一样，但是ok会返回是否断言成功不会出现panic，ok就表示是否是成功了

强制类型转换
该方法不常见，主要用于 unsafe 包和接口类型检测，需要懂得go变量的知识
unsafe
var f float64
bits = *(*uint64)(unsafe.Pointer(&f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))

var p ptr = nil
unsafe 强制转换是指针的底层操作了，用 c 的朋友就很熟悉这样的指针类型转换，利用内存对齐才能保证转换可靠，例如 int 和 uint 存在符号位差别，利用 unsafe 转换后值可能不同，但是在内存存储二进制一模一样。

接口类型检测
var _ Context = (*ContextBase)(nil)

nil 的类型是 nil 地址值为 0，利用强制类型转换成了 * ContextBase，返回的变量就是类型为 * ContextBase 地址值为 0，然后 Context=xx 赋值如果 xx 实现了 Context 接口就没事，如果没有实现在编译时期就会报错，实现编译期间检测接口是否实现。


显示类型转换
一个显示转换的表达式T(x),其中T是一种类型并且x是可转换为类型的表达式T，例如:uint(666)

x 可以分配成 T 类型。
忽略 struct 标签 x 的类型和 T 具有相同的基础类型。
忽略 struct 标记 x 的类型和 T 是未定义类型的指针类型，并且它们的指针基类型具有相同的基础类型。
x 的类型和 T 都是整数或浮点类型。
x 的类型和 T 都是复数类型。
x 的类型是整数或 [] byte 或 [] rune，并且 T 是字符串类型。
x 的类型是字符串，T 类型是 [] byte 或 [] rune。

int64(222)
[]byte("ssss")

type A int
A(2)


隐式类型转换
隐式类型转换日常使用并不会感觉到，但是运行中确实出现了类型转换，以下列出了两种
组合间的重新断言类型

type Reader interface {
    Read(p []byte) (n int, err error)
}
type ReadCloser interface {
    Reader
    Close() error
}
var rc ReaderClose
r := rc

ReaderClose 接口组合了 Reader 接口，但是 r=rc 的赋值时还是类型转换了，go 使用系统内置的函数执行了类型转换。以前遇到过类似接口组合类型的变量赋值，然后使用 pprof 和 bench 测试发现了这一细节，在接口类型转移时浪费了一些性能。

相同类型间赋值

type Handler func()

func NewHandler() Handler {
    return func() {}
}
虽然 type 定义了 Handler 类型，但是 Handler 和 func () 是两种实际类型，类型不会相等，使用反射和断言均会出现两种类型不同。

package main

import (
    "fmt"
    "reflect"
)

type Handler func()

func a() Handler {
    return func() {}
}

func main() {
    var i interface{} = main
    _, ok := i.(func())
    fmt.Println(ok)
    _, ok = i.(Handler)
    fmt.Println(ok)
    fmt.Println(reflect.TypeOf(main) == reflect.TypeOf((*Handler)(nil)).Elem())
}

// true
// false
// false

*/
