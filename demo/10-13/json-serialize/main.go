package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type st struct {
	Field1 json.RawMessage `json:"field1"`
	Field2 json.RawMessage `json:"field2"`
}

func main() {
	m := map[string]json.RawMessage{}
	var testBytes = []byte(`
{"field1":100,"field2":"Hello, World!"}
`)
	err := json.Unmarshal(testBytes, &m)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(m)
	s := st{}
	err = json.Unmarshal(testBytes, &s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s)
}
