package main

import (
	"github.com/imdario/mergo"
	"log"
	"reflect"
	"time"
)

type timeTransformer struct {
}

type timeAddOneDayTransformer struct {
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
				//通过反射值使用MethodByName调用IsZero方法，判断是否是字段的零值
				isZero := dst.MethodByName("IsZero")

				//获取IsZero方法的调用结果
				result := isZero.Call([]reflect.Value{})

				//如果目标段为零值则覆盖
				if result[0].Bool() {
					dst.Set(src)
				}
			}
			return nil
		}
	}
	return nil
}

func (t timeAddOneDayTransformer) Transformer(typ reflect.Type) func(dst, src reflect.Value) error {
	if typ == reflect.TypeOf(time.Time{}) {
		return func(dst, src reflect.Value) error {
			//判断目标值是否是地址类型（只有地址类型才能被覆盖）
			if dst.CanSet() {
				//在原始值的时间上加1天
				addedTime := src.Interface().(time.Time).AddDate(0, 0, 1)
				//将结果写入到目标值中
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
	//time.Now()的返回值虽然是time.Time类型，但和原始的time.Time时间格式并不相同
	//在合并过程中time.Now()的值并不会被合并到目标结构体中，这一点是跟其他类型表现不同的
	log.Println("(time.Time:", time.Time{})
	log.Println("time.Now:", time.Now())
	src := Snapshot{now}
	dest := Snapshot{}
	mergo.Merge(&dest, src)
	//我们可以通过自定义合并规则来控制合并结果
	//mergo.Merge(&dest, src, mergo.WithTransformers(timeTransformer{}))
	mergo.Merge(&dest, src, mergo.WithTransformers(timeAddOneDayTransformer{}))
	log.Println(src)
	log.Println(dest)

}
