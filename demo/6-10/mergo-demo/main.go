package main

import (
	"github.com/imdario/mergo"
	"log"
)

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type Person struct {
	//注意，小写字母开头的字段无法在转化过程中会丢失
	name    string
	Age     int
	Address Address
}

type Address struct {
	County string
	City   string
}

type User struct {
	Name  string
	Sex   bool
	hobby string
}

type User2 struct {
	Name  string
	hobby string
}

func main() {
	// 构造一个结构体
	p1 := Person{
		name: "tester1",
		Age:  18,
		Address: Address{
			County: "cn",
			City:   "beijing",
		},
	}

	// 将结构体转换为map
	m1 := map[string]interface{}{}
	//传进来的map必须使用地址
	err := mergo.Map(&m1, p1)
	if err != nil {
		panic(err)
	}
	log.Println("m1: %v", m1)

	// 构造一个map
	m2 := map[string]interface{}{
		"name": "tester2",
		"Age":  30,
		"Address": map[string]interface{}{
			"County": "CA",
			"City":   "San Francisco",
		},
	}

	// 将map转换为结构体
	p2 := Person{}
	//传入的结构体必须是地址
	err = mergo.Map(&p2, m2)
	if err != nil {
		panic(err)
	}
	log.Println("p2: %v", p2)

	//合并结构体
	u1 := User{
		Name:  "tester3",
		Sex:   false,
		hobby: "swimming",
	}

	//合并结构体,将u1中的字段复制到p1
	//失败(src and dst must be of same type)
	err = mergo.Merge(&p1, u1)
	if err != nil {
		log.Println(err)
	}
	log.Println("p1", p1)

	u2 := User2{
		Name:  "tester4",
		hobby: "ping-pong",
	}

	//合并结构体,将u2中的字段复制到u1
	//成功
	err = mergo.Merge(&u1, u2)
	if err != nil {
		log.Println(err)
	}
	//u1中的已有字段Name值未被覆盖
	log.Println("u1", u1)

	//将u1的Name值赋值为默认值
	u1.Name = ""
	err = mergo.Merge(&u1, u2)
	if err != nil {
		log.Println(err)
	}
	//u1中的已有字段Name值被覆盖
	//由于u1中的name为默认值，在merge后将会以u2覆盖当前值
	//由于u2中没有Name字段，所以合并后u1的Name字段会丢失
	log.Println("u1", u1)

	//合并结构体,将u1中的字段复制到u2
	//失败（src and dst must be of same type）
	err = mergo.Merge(&u2, u1)
	if err != nil {
		log.Println(err)
	}
	log.Println("u2", u2)

	//总结：
	//1. 第二个参数中的字段必须是第一个字段的子集
	//比如上面将u2复制到u1时，u2的字段是u1的子集，所以能够复制，但不能将u2复制到u1
}
