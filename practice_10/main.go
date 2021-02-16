package main

import "fmt"

func main() {
	/*
		切片 slice
		切片又被成为动态数组，切片的长度可以改变的
		切片由三部分组成：指向底层数组的指针，切片的元素个数（长度），切片的容量
		引用类型
		遍历方式和数组类似，都是通过for，for range
	*/

	// 切片的定义
	var sliceVar []int
	//定义一个数组
	arrayVar := [...]int{12, 21, 23, 55, 98, 2}
	for i := 0; i < len(arrayVar); i++ {
		fmt.Printf("arrayVar[%d]=%d，地址=%p\n", i, arrayVar[i], &arrayVar[i])
	}
	// 切片去引用数组
	sliceVar = arrayVar[:]
	for i := 0; i < len(sliceVar); i++ {
		fmt.Printf("sliceVar[%d]=%d，地址=%p\n", i, sliceVar[i], &sliceVar[i])
	}

	var sliceVar2 []int
	sliceVar2 = arrayVar[1:3]
	fmt.Println(sliceVar2)

	// 切片指向底层的数组
	for i := 0; i < len(sliceVar2); i++ {
		fmt.Printf("sliceVar2[%d]=%d，地址=%p\n", i, sliceVar2[i], &sliceVar2[i])
	}

	sliceVar2[0] = 100
	fmt.Println(arrayVar)
	fmt.Println(sliceVar2)

	// 切片的第二种定义方式
	var sliceVar3 []int = make([]int, 5, 6)
	fmt.Printf("sliceVar3的长度=%d, 容量=%d,切片指向的底层数组的地址=%p,切片自己的地址%p\n", len(sliceVar3), cap(sliceVar3), sliceVar3, &sliceVar3)

	sliceVar3 = append(sliceVar3, 7)
	fmt.Printf("第一次追加sliceVar3的长度=%d, 容量=%d,切片指向的底层数组的地址=%p,切片自己的地址%p\n", len(sliceVar3), cap(sliceVar3), sliceVar3, &sliceVar3)

	sliceVar3 = append(sliceVar3, 8)
	fmt.Printf("第二次追加sliceVar3的长度=%d, 容量=%d,切片指向的底层数组的地址=%p,切片自己的地址%p\n", len(sliceVar3), cap(sliceVar3), sliceVar3, &sliceVar3)
	fmt.Println(sliceVar3)

	// 切片的copy
	sliceVar4 := []int{1, 2, 3, 4, 5}
	sliceVar5 := make([]int, 10)
	fmt.Println(sliceVar4)
	fmt.Println(sliceVar5)
	copy(sliceVar5, sliceVar4)
	fmt.Println(sliceVar4)
	fmt.Println(sliceVar5)

	sliceVar6 := make([]int, 1)
	copy(sliceVar6, sliceVar4)
	fmt.Println(sliceVar4)
	fmt.Println(sliceVar6)

	// 切片是引用类型
	fmt.Println(sliceVar5)
	changeSlice(sliceVar5)
	fmt.Println(sliceVar5)

}

func changeSlice(slice []int) {
	slice[0] = 100
}
