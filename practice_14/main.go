package main

import (
	"fmt"
	"os"
)

var vara = 100

func init() {
	fmt.Println("init...")
}
func main() {
	// 执行时顺序：全局变量-->init()-->main()
	fmt.Println("main...")
	/*
		函数
		func 函数名（参数）（返回值）{
			//函数体
			return xxx
		}
	*/
	var sum int
	var minus int
	sum, minus = actionVar1(1, 2)
	fmt.Println(sum, minus)

	sum, minus = actionVar2(2, 3)
	fmt.Println(sum, minus)
	fmt.Println("hello xueguojun", 3.14)

	actionVar3(1, "hani", 3.14, true, 45, 33)

	var args int    // 默认为0
	testScope(args) // 执行完之后args被内存回收
	fmt.Println(args)

	//递归函数
	//斐波那契数列 特点
	/*
		位置    值
		1		1
		2		1
		3		2
		4		3
		...		...
		n		func(n-1) + func(n-2)
	*/
	// 第20位置对应的值
	fmt.Println(fibonaci(20))

	//作业1
	//要求通过递归的方式来求某个数阶乘，阶乘的特点 5*4*3*2*1
	//传递一个数，计算该数的阶乘

	// 匿名函数和闭包
	// 匿名函数的使用方式一：定义一个匿名函数并调用

	func(var1, var2 int) int { return 1 }(1, 3)

	sumVar := func(var1, var2 int) int {
		return var1 + var2
	}(1, 3)
	fmt.Println(sumVar)

	//匿名函数的使用方式二：
	//将一个匿名函数赋值给一个变量，这个变量存储的时这个匿名函数的地址
	func1 := func(var1, var2 int) int {
		return var1 - var2
	}
	fmt.Printf("func1调用的值=%d,func1=%p\n", func1(1, 2), &func1)

	// 闭包，函数+对函数外部数据的引用
	func2 := clousure()
	fmt.Println(func2(1))
	fmt.Println(func2(2))
	fmt.Println(func2(3))

	fmt.Println()
	//defer
	deferAction2(1, 2)

	// 函数不支持重载

	//函数可以作为一种类型
	sum, minus = funcAsArgs(actionVar4, 1, 2)
	fmt.Println(sum, minus)

	//函数作为值 type
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	evenSlice := judgeSlice(slice, isEven)
	oddSlice := judgeSlice(slice, isOdd)
	fmt.Println("偶数：", evenSlice)
	fmt.Println("奇数：", oddSlice)

	//main init
	// 作业2：以闭包的方式来完成，判断一个字符串是否包含"xxx@imooc.com",如果包含则提交到服务端，如果没有，自动添加再提交到服务端
	//map，slice，map

}

type Boolean func(int) bool

//是否为偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 是否为奇数
func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func judgeSlice(slice []int, boolean Boolean) (result []int) {
	for _, value := range slice {
		if boolean(value) {
			result = append(result, value)
		}
	}
	return result
}

func actionVar4(var1, var2 int) (sum int, minus int) {
	sum = var1 + var2
	minus = var1 - var2
	//return sum, minus
	return
}

func funcAsArgs(funcName func(int, int) (int, int), var1, var2 int) (int, int) {
	return funcName(var1, var2)
}

// defer 先进后出
func deferAction2(var1, var2 int) int {
	defer fmt.Println(var1)
	defer fmt.Println(var2)
	sum := var1 + var2
	defer fmt.Println(sum)
	return sum
}

func deferAction() {
	handle, err := os.Open("hello.txt")
	defer handle.Close()
	//这里是一大堆的文件操作程序。。。
	//fmt.Println(5 / 0)
	err = err

}

func clousure() func(int64) int64 {
	var step int64 = 0
	return func(i int64) int64 {
		step += i
		return step
	}
}

func fibonaci(index int64) int64 {
	if index == 1 || index == 2 {
		return 1
	} else {
		return fibonaci(index-1) + fibonaci(index-2)
	}
}

func actionVar1(var1 int, var2 int) (int, int) {
	return var1 + var2, var1 - var2
}

func actionVar2(var1, var2 int) (sum, minus int) {
	sum = var1 + var2
	minus = var1 - var2
	return
}

func actionVar3(args ...interface{}) {
	for _, value := range args {
		//fmt.Println(key,value)
		switch value.(type) {
		case int:
			fmt.Println(value, "int")
		case string:
			fmt.Println(value, "string")
		case float64:
			fmt.Println(value, "float64")
		case bool:
			fmt.Println(value, "bool")
		default:
			fmt.Println(value, "unknown")
		}
	}
}

func testScope(args int) {
	args = 100
	fmt.Println(args)
}
