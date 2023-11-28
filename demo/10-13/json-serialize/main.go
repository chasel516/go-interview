package main

import (
	"encoding/json"
	"fmt"
	"log"
)

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
}
