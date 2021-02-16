package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//浮点类型 float默认float64
	//复数：
	//complex64 32位实数+32位虚数，虚数的单位i
	//complex128 64位实数+64位虚数，虚数的单位i
	/*
		float32 4个字节
		float64 8个字节
	*/

	// 1.以10进制的方式展示
	var floatVar1 float32 = 3.1415926
	fmt.Printf("floatvar1的类型=%T，占用的字节大小=%d\n", floatVar1, unsafe.Sizeof(floatVar1))
	floatVar2 := .1415926 // 0.1415926
	fmt.Printf("floatvar2的类型=%T，占用的字节大小=%d\n", floatVar2, unsafe.Sizeof(floatVar2))

	// 2.以科学计数法来展示
	floatVar3 := 3.1415926e2  //3.1415926乘以10的2次方
	floatVar4 := 3.1415926e-2 //3.1415926除以10的2次方
	fmt.Println(floatVar3, floatVar4)

	var floatVar5 float32 = 3.14
	var floatVar6 float64 = 3.14
	floatVar6 = float64(floatVar5)
	floatVar6 = floatVar6

	// 负数  实数 + 虚数i

	var complexVar1 complex64
	complexVar1 = 3.14 + 12i
	complexVar2 := complex(3.14, 12)

	fmt.Printf("complexVar1的类型=%T，值=%v\n", complexVar1, complexVar1)
	fmt.Printf("complexVar2的类型=%T，值=%v\n", complexVar2, complexVar2)
	// 分别打印实部和虚部
	fmt.Println(real(complexVar1), imag(complexVar1))
}
