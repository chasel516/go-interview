package main

import (
	"github.com/imdario/mergo"
	"log"
	"reflect"
	"time"
)

type timeTransformer struct {
}

func init() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// 自定义转换器
func (t timeTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			if dst.CanSet() {
				addedTime := src.Interface().(time.Time).AddDate(0, 0, 1)
				dst.Set(reflect.ValueOf(addedTime))
			}
			return nil
		}
	}
	return nil
}

type Snapshot struct {
	Time time.Time
	// ...
}

func main() {
	now := time.Now()
	src := Snapshot{now}
	dest := Snapshot{}
	mergo.Merge(&dest, src, mergo.WithTransformers(timeTransformer{}))
	//mergo.Merge(&dest, src)
	log.Println(src)
	log.Println(dest)

}
