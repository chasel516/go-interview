package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var i uint64 = math.MaxUint64
	fmt.Println(intFromInterface(i))

}
func intFromInterface(selector interface{}) (value int, err error) {
	maxInt := 1<<(strconv.IntSize-1) - 1
	minInt := -1 << (strconv.IntSize - 1)

	switch selector.(type) {
	case int:
		value = selector.(int)
	case int8:
		value = int(selector.(int8))
	case int16:
		value = int(selector.(int16))
	case int32:
		value = int(selector.(int32))
	case int64:
		v := selector.(int64)
		//32为操作系统解析出的int类型是int32，这里防止32位整型溢出
		if strconv.IntSize <= 32 && (v > int64(maxInt) || v < int64(minInt)) {
			err = errors.New("convert overflow")
			return
		}
		value = int(v)
	case uint:
		value = int(selector.(uint))
		if value > maxInt {
			err = errors.New("convert overflow")
			return
		}
	case uint8:
		value = int(selector.(uint8))
	case uint16:
		value = int(selector.(uint16))
	case uint32:
		v := selector.(uint32)
		if uint64(v) > uint64(maxInt) {
			err = errors.New("convert overflow")
			return
		}
		value = int(v)
	case uint64:
		v := selector.(uint64)
		if v > uint64(maxInt) {
			err = errors.New("convert overflow")
			return
		}
		value = int(v)

	case string:
		var v int64
		v, err = strconv.ParseInt(selector.(string), 10, 64)
		if err != nil {
			return
		}
		//32为操作系统解析出的int类型是int32，这里防止32位整型溢出
		if v > int64(maxInt) || v < int64(minInt) {
			err = errors.New("convert overflow")
			return 0, err
		}

	default:
		err = errors.New("unexpected type")
	}
	return
}
