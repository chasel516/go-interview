package pkg1

import "fmt"

var TestVer1 = "TestVer1"
var tesVer2 = "tesVer2"

const (
	TestConst1 = "TestConst1"
	testConst2 = "testConst2"
)

type TestStruct1 struct {
	Field1 string
	field2 string
}

type testStruct2 struct {
	Field1 string
	field2 string
}

func (ts1 TestStruct1) Test1() {
	fmt.Println(TestConst1)
}

func (ts1 TestStruct1) test2() {
	fmt.Println(testConst2)
}

func (ts2 testStruct2) Test21() {
	fmt.Println(testConst2)
}

func (ts2 testStruct2) test21() {
	fmt.Println(testConst2)
}

func f() {
	fmt.Println(testConst2, tesVer2)
	t := testStruct2{
		Field1: "we",
		field2: "lee",
	}
	t.test21()
}
