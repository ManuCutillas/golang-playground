package main

import (
	"reflect"
	"fmt"
)

func main() {
	numerics()
}

func numerics() {
	var _bool bool = !("hello" == "hello")
	fmt.Println(reflect.DeepEqual(_bool, false))
	fmt.Println(reflect.TypeOf(_bool))
	fmt.Println(reflect.ValueOf(_bool))

	var _string string = "Hello"
	fmt.Println(reflect.DeepEqual(_string, "Hello"))
	fmt.Println(reflect.TypeOf(_string))
	fmt.Println(reflect.ValueOf(_string))

	var _int int = 0
	fmt.Println(reflect.DeepEqual(_int, 0))
	fmt.Println(reflect.TypeOf(_int))
	fmt.Println(reflect.ValueOf(_int))

	var _int8 int8 = -128 // -128 to 127
	fmt.Println(reflect.DeepEqual(_int8, -128))
	fmt.Println(reflect.TypeOf(_int8))
	fmt.Println(reflect.ValueOf(_int8))

	var _int16 int16 = -32768 // -32768 to 32767
	fmt.Println(reflect.DeepEqual(_int16, -32768))
	fmt.Println(reflect.TypeOf(_int16))
	fmt.Println(reflect.ValueOf(_int16))

	var _int32 int32 = -2147483648 // -2147483648 to 2147483647
	fmt.Println(reflect.DeepEqual(_int32, -2147483648))
	fmt.Println(reflect.TypeOf(_int32))
	fmt.Println(reflect.ValueOf(_int32))

	var _int64 int64 = -9223372036854775808 // -9223372036854775808 to 9223372036854775807
	fmt.Println(reflect.DeepEqual(_int64, -9223372036854775808))
	fmt.Println(reflect.TypeOf(_int64))
	fmt.Println(reflect.ValueOf(_int64))

	var _uint8 uint8 = 255 // 0 to 255
	fmt.Println(reflect.DeepEqual(_uint8, 255))
	fmt.Println(reflect.TypeOf(_uint8))
	fmt.Println(reflect.ValueOf(_uint8))

	var _uint16 uint16 = 65535 // 0 to 65535
	fmt.Println(reflect.DeepEqual(_uint16, 65535))
	fmt.Println(reflect.TypeOf(_uint16))
	fmt.Println(reflect.ValueOf(_uint16))

	var _uint32 uint32 = 4294967295 // 0 to 4294967295
	fmt.Println(reflect.DeepEqual(_uint32, 4294967295))
	fmt.Println(reflect.TypeOf(_uint32))
	fmt.Println(reflect.ValueOf(_uint32))

	var _uint64 uint64 = 9223372036854775807 // 0 to 9223372036854775807
	fmt.Println(reflect.DeepEqual(_uint64, 9223372036854775807))
	fmt.Println(reflect.TypeOf(_uint64))
	fmt.Println(reflect.ValueOf(_uint64))

	var _float32 float32 = 0.345
	fmt.Println(reflect.DeepEqual(_float32, 0.345))
	fmt.Println(reflect.TypeOf(_float32))
	fmt.Println(reflect.ValueOf(_float32))

	var _float64 float64 = 0.345
	fmt.Println(reflect.DeepEqual(_float64, 0.345))
	fmt.Println(reflect.TypeOf(_float64))
	fmt.Println(reflect.ValueOf(_float64))

	var _complex64 complex64 = 0.345
	fmt.Println(reflect.DeepEqual(_complex64, 0.345))
	fmt.Println(reflect.TypeOf(_complex64))
	fmt.Println(reflect.ValueOf(_complex64))

	var _complex128 complex128 = 0.345
	fmt.Println(reflect.DeepEqual(_complex128, 0.345))
	fmt.Println(reflect.TypeOf(_complex128))
	fmt.Println(reflect.ValueOf(_complex128))

	var _byte byte = 255 // alias to uint8
	fmt.Println(reflect.DeepEqual(_byte, 255))
	fmt.Println(reflect.TypeOf(_byte))
	fmt.Println(reflect.ValueOf(_byte))

	var _rune rune = 2147483647 // alias to int32
	fmt.Println(reflect.DeepEqual(_rune, 2147483647))
	fmt.Println(reflect.TypeOf(_rune))
	fmt.Println(reflect.ValueOf(_rune))
}


