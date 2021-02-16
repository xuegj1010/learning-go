package main

import "fmt"

func main() {
	/*
		指针：存储的是另一个变量的内存地址
		指针作用：
			1、节省内存空间，提高程序执行效率
			2、间接访问与修改变量的值
	*/
	var intVar1 int = 100
	fmt.Printf("intVar1的值=%d,地址=%v\n", intVar1, &intVar1)

	var pointerVar1 *int = &intVar1
	fmt.Printf("intVar1的值=%v,地址=%v\n", pointerVar1, &pointerVar1)
	/*
		变量名        变量值           地址
		intVar1      100             0xc0000180c0
		pointerVar1  0xc0000180c0    0xc00000e030
	*/
	fmt.Println(intVar1)
	*pointerVar1 = 200
	fmt.Println(intVar1)

	var var2 int = 123
	var pointerVar2 *int = &var2
	fmt.Println(*pointerVar2, var2)

	/*
		值类型：整型，浮点型，bool，array，string，栈中分配
		引用类型：指针，slice，map，chan，interface，堆中分配
	*/
}
