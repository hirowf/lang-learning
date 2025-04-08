package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64

	//var n int64 = 100000000000000
	var n int = 100000000000000 // uses arch processor type

	fmt.Println(n)

	var n2 uint64 = 100000 // assigned
	fmt.Println(n2)

	//alias
	//INT32 = RUNE
	var n32bit rune = 123456

	fmt.Println(n32bit)

	// BYTE - UNIT8
	var n4 byte = 123
	fmt.Println(n4)

	// float32, aafloat64

	var realN float32 = 123.45
	fmt.Println(realN)

	realN3 := 1234321312.43
	fmt.Println(realN3)

	// strings

	var str string = "BLABALBALBAL"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	char := '%'
	fmt.Println(char)

	// value 0
	// var text int
	var text string
	fmt.Println(text)

	// bool

	var bool1 bool
	fmt.Println(bool1)

	var err error = errors.New("internal error")

	fmt.Println(err)
}
