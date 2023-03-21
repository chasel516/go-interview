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
		//IntPtr         //具名指针类型
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
