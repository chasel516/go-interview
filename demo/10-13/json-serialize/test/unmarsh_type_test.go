package test

import (
	"encoding/json"
	"testing"
)

// 结构体类型
type TestDataStruct struct {
	Field1 int    `json:"field1"`
	Field2 string `json:"field2"`
}

// map 类型
type TestDataMap map[string]interface{}

var jsonStructData = []byte(`{"field1": 42, "field2": "Hello, World!"}`)
var jsonMapData = []byte(`{"field1": 42, "field2": "Hello, World!"}`)

func BenchmarkJsonStdUnmarshalStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var data TestDataStruct
		if err := json.Unmarshal(jsonStructData, &data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkJsonStdUnmarshalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var data TestDataMap
		if err := json.Unmarshal(jsonMapData, &data); err != nil {
			b.Fatal(err)
		}
	}
}
