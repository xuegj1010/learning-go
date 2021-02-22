package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//整型，数值类型
	/*
		有符号位
		int8  1  -2^7~2^7-1
		int16 2  -2^15~2^15-1
		int32 4  -2^31~2^31-1
		int64 8  -2^63~2^63-1
		无分符号位
		uint8  1  0～2^8-1
		uint16 2  0~2^15-1
		uint32 4  0~2^32-1
		uint64 8  0~2^64-1
	*/

	var intVar1 = 100 // 默认对应int类型
	intVar2 := 200    // int

	var intVar3 int32
	intVar4 := 126

	intVar3 = int32(intVar4)

	fmt.Printf("intVar1=%T, intVar2=%T,intVar3=%T", intVar1, intVar2, intVar3)

	var intVar5 int64 = 123456789
	fmt.Println()
	fmt.Println(unsafe.Sizeof(intVar5))
}
