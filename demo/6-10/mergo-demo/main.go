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
	Hobby string
}

type User2 struct {
	Name  string
	hobby string
	Hobby string
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
	log.Printf("m1: %+v", m1)

	// 构造一个map（注意：结构体和map中的字段名需要一一对应）
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
	log.Printf("p2: %+v", p2)

	//合并结构体
	u1 := User{
		Name:  "tester3",
		Sex:   false,
		hobby: "",
		Hobby: "swimming",
	}

	//合并结构体,将u1中的字段复制到p1
	//失败(src and dst must be of same type)
	err = mergo.Merge(&p1, u1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("p1:%+v", p1)

	//u2 := User2{
	//	Name:  "tester4",
	//	hobby: "ping-pong",
	//}
	u2 := User{
		Name:  "tester4",
		Sex:   true,
		hobby: "ping-pong",
		Hobby: "ping-pong",
	}

	//合并结构体,将u2中的字段复制到u1
	//成功
	err = mergo.Merge(&u1, u2)
	if err != nil {
		log.Println(err)
	}
	//u1中的已有字段Name值未被覆盖
	log.Printf("u1:%+v", u1)

	//将u1的Name值赋值为默认值
	u1.Name = ""
	err = mergo.Merge(&u1, u2)
	if err != nil {
		log.Println(err)
	}
	//u1中的已有字段Name值被覆盖
	//由于u1中的name为默认值，在merge后将会以u2覆盖当前值
	//由于u2中没有Name字段，所以合并后u1的Name字段会丢失
	log.Printf("u1:%+v", u1)

	//合并结构体,将u1中的字段复制到u2
	//失败（src and dst must be of same type）
	err = mergo.Merge(&u2, u1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("u2:%+v", u2)

	//结构体合并总结：
	//注：我们将第一个参数成为目标结构体，第二个参数称为原结构体
	//1. 原结构和目标结构体必须是同一个结构体类型的实例
	//2.当目标结构体的字段为默认值时，将会被原结构体的字段覆盖，但当目标结构体字段值不为默认值时，其字段和值不会被原结构体字段覆盖
	//3.若目标结构体字段时小写开头的非导出字段，则在合并的过程中目标结构体中此字段的值也不会被合并

	//可以通过mergo提供的一些默认配置来控制这些规则：
	//比如：

	u3 := User{
		Name:  "test",
		Sex:   true,
		hobby: "test",
		Hobby: "test",
	}
	u4 := User{
		Name:  "",
		Sex:   false,
		hobby: "",
		Hobby: "",
	}

	//mergo.Merge(&u3, u4)
	//允许覆盖默认值（空值）
	mergo.Merge(&u3, u4, mergo.WithOverwriteWithEmptyValue)
	log.Printf("u3:%+v", u3)
}
