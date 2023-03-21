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
