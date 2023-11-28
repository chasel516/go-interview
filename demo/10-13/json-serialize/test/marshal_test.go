package test

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

type TestData struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

func BenchmarkStdMarshal(b *testing.B) {
	var testData = TestData{
		Field1: 100,
		Field2: "Hello, World!",
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(testData)
	}
}

func BenchmarkStdUnMarshal(b *testing.B) {
	var testBytes = []byte(`
{"field1":100,"field2":"Hello, World!"}
`)
	var testData = TestData{}
	for i := 0; i < b.N; i++ {
		json.Unmarshal(testBytes, &testData)
	}
}

func BenchmarkJsonIterMarshal(b *testing.B) {
	var testData = TestData{
		Field1: 100,
		Field2: "Hello, World!",
	}
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(&testData)
	}
}
func BenchmarkJsonIterUnMarshal(b *testing.B) {
	var testBytes = []byte(`
{"field1":100,"field2":"Hello, World!"}
`)
	var testData = TestData{}
	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(testBytes, &testData)
	}
}
