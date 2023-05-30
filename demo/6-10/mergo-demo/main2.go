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
		"sex":   "男",
	}
	map2 := map[string]interface{}{
		"phone": "123-456-7890",
		"email": "",
		"sex":   true,
		//"sex":   false,
	}

	// merge maps
	//map2的email值并未覆盖map1,且map2中的新字段sex被增加到map1
	err := mergo.Merge(&map1, map2)
	if err != nil {
		log.Println(err)
	}
	// print merged map
	log.Println(map1)

	//map2的email为默认值值，被map1的email值覆盖, 且map2中的新字段sex在合并过程中依然保留
	err = mergo.Merge(&map2, map1)
	if err != nil {
		log.Println(err)
	}
	// print merged map
	log.Println(map2)

	//map合并总结：
	//1.当原map和目标map中存在相同的key值时，只有当目标map的值为默认值时，才会被原map的key所覆盖；

	//2. 若目标map中存在原map缺失的字段，则合并后的map会新增此字段。

	//所以合并两个map最终的字段数量可以理解为取并集，而当两个map存在相同的key时，
	//依然只有默认值才会被覆盖，否则其值依然以目标map的值为准
}
