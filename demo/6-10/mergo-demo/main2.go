package main

import (
	"github.com/imdario/mergo"
	"log"
)

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
	// define two maps
	map1 := map[string]interface{}{
		"name":  "John Doe",
		"email": "tester@example.com",
	}
	map2 := map[string]interface{}{
		"phone": "123-456-7890",
		"email": "",
	}

	// merge maps
	err := mergo.Merge(&map1, map2)
	if err != nil {
		log.Println(err)
	}

	// print merged map
	log.Println(map1)

	err = mergo.Merge(&map2, map1)
	if err != nil {
		log.Println(err)
	}
	// print merged map
	log.Println(map2)

	//map合并总结：
	//默认值字段不会被覆盖

}
