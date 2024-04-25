package main

func main() {
	type m map[string]int
	type person struct {
		name string
		m
		int
	}

	p := person{}
	p.name = "lee"
	p.int = 1
	p.m = nil

	type IPerson interface{ say(words string) }
	type Person struct {
		name string
		age  int
	}
	type PersonPtr = *struct {
		name string
		age  int
	}

	type Int int
	type IntPtr *int
	type AliasIntPtr *IntPtr

	type s1 struct {
		// 可以被内嵌的类型
		//m
		//*m
		//Int
		//*Int
		//IPerson
		//Person
		//*Person
		//PersonPtr
		//int
		//*int

		//不能被内嵌的类型：
		// IntPtr         //具名指针类型
		//*IntPtr        //基类型IntPtr为指针类型
		//AliasIntPtr    //基类型AliasIntPtr为指针类型
		//*PersonPtr     //基类型PersonPtr为指针类型
		//*IPerson       //基类型IPerson为接口类型
		//[]int          //无名非指针类型
		//map[string]int //无名非指针类型
		//func()         //无名非指针类型
	}

	//Worker继承了Person,可以调用Person的方法
	w := Worker{}
	w.setName()

}

//一个内嵌字段必须被声明为形式T或者一个基类型为非接口类型的指针类型*T，其中T为一个类型名但是T不能表示一个指针类型。

// 一个类型名T只有在它既不表示一个具名指针类型也不表示一个基类型为指针类型或者接口类型的指针类型的情况下才可以被用做内嵌字段。
// 一个指针类型*T只有在T为一个类型名并且T既不表示一个指针类型也不表示一个接口类型的时候才能被用做内嵌字段。
// 一个结构体类型不能内嵌（无论间接还是直接）它自己。
type Person struct {
	name string
	age  int
}

func (p *Person) setName() {
	p.name = "lee"
}

type Worker struct {
	Person
}
