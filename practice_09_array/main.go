package main

import "fmt"

func main() {
	/*
		数组：数组是长度固定的数据类型，元素的类型都是一样的，数组在内存当中时连续的存储空间，可以有效的提升CPU的执行效率

		变量或长说数据在定义时，如果没有赋值，会有默认值，零值
		整型：默认值0
		字符串：默认值 ""
		bool：默认值 false
	*/
	// 数组定义方式1
	var arrayVar1 [10]int
	arrayVar1[0] = 100
	arrayVar1[3] = 200
	fmt.Println(arrayVar1)

	var arrayVar2 [5]int = [5]int{1, 2, 3, 4, 5}
	var length = len(arrayVar2)
	for i := 0; i < length; i++ {
		fmt.Printf("arrayVar2[%d]=%d,地址=%p\n", i, arrayVar2[i], &arrayVar2[i])
	}
	/*
		数组在内存当中时连续的存储空间
		数据的元素（下标）    元素值      元素的地址
			 0              1         0xc000114000
			 1              2         0xc000114008
			 2              3         0xc000114010
			 3              4         0xc000114018
			 4              5         0xc000114020
	*/
	// 数组定义方式2
	// [10]int [5]int 被认为是不同的两种类型
	// 数组的长度是数组的组成部分
	arrayVar3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arrayVar3, len(arrayVar3))

	// 数组定义方式3
	arrayVar4 := [...]int{100: 200, 300: 500} //值200的下标是100,值500的下标是300
	fmt.Println(arrayVar4, len(arrayVar4))

	// 在golang中，数组是值类型
	arrayVar5 := [10]int{1, 2, 3}
	fmt.Println(arrayVar5)
	changeArr(arrayVar5)
	fmt.Println(arrayVar5)

	// 通过传递指针的方式来改变数组的元素
	changeArrByPointer(&arrayVar5)
	fmt.Println(arrayVar5)

	// 将一个数组赋值给另一个数组
	arrayVar6 := [3]string{"hello", "你好", "golang"}
	arrayVar7 := arrayVar6
	// for ... range
	for key, value := range arrayVar6 {
		fmt.Printf("arrayVar6[%d]=%v,地址=%p\n", key, value, &arrayVar6[key])
	}
	for key, value := range arrayVar7 {
		fmt.Printf("arrayVar7[%d]=%v,地址=%p\n", key, value, &arrayVar7[key])
	}

	// 二维数组
	var arrayVar8 [4][2]int
	fmt.Println(arrayVar8)

	arrayVar9 := [4][2]int{{10, 11}, {3, 5}, {2, 3}, {100, 88}}
	fmt.Println(arrayVar9)

	arrayVar10 := [4][2]int{1: {100, 90}, 2: {8, 9}}
	fmt.Println(arrayVar10)

	// 多维数组
	var arrayVar11 [4][3][2]int = [4][3][2]int{}
	fmt.Println(arrayVar11)

}

func changeArrByPointer(arr *[10]int) {
	(*arr)[0] = 100
}

func changeArr(arr [10]int) {
	arr[0] = 100
}
