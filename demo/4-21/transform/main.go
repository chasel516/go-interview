package main

func main() {

}
func intFromInterface(selector interface{}) int {
	var value int
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
		value = int(selector.(int64))
	case uint:
		value = int(selector.(uint))
	case uint8:
		value = int(selector.(uint8))
	case uint16:
		value = int(selector.(uint16))
	case uint32:
		value = int(selector.(uint32))
	case uint64:
		value = int(selector.(uint64))

	case string:

	default:
		panic("objx: array access argument is not an integer type (this should never happen)")
	}
	return value
}
