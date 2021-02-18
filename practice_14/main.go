package main

import "fmt"

func main() {
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
