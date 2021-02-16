package main

import "fmt"

func main() {
	/*
		字符串：string
	*/

	var strVar1 string
	strVar1 = "hello, 世界\n"
	strVar2 := `
package main

import "fmt"

func main() {
	// 常量
	const constVar1 float64 = 3.1415926
	const constVar2, constVar3 = 100, "xueguojun"
	fmt.Println(constVar1, constVar2, constVar3)

	const (
		iotaVar1 = iota
		iotaVar2 = iota
		iotaVar3 = iota
	)
	const iotaVar4 = iota
	fmt.Println(iotaVar1, iotaVar2, iotaVar3, iotaVar4)

	const (
		iotaVar5 = iota
		iotaVar6
		iotaVar7
	)
	fmt.Println(iotaVar5, iotaVar6, iotaVar7)

	const (
		Monday = iota
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	const (
		iotaVar8, iotaVar9, iotaVar10 = iota, iota, iota
	)
	fmt.Println(iotaVar8, iotaVar9, iotaVar10)

	const (
		iotaVar11 = iota
		iotaVar12 = "Bob"
		iotaVar13 = iota
	)
	fmt.Println(iotaVar11, iotaVar12, iotaVar13)
}
`
	fmt.Printf("strVar1=%T,strVar2=%T\n", strVar1, strVar2)
	fmt.Printf("strVar1=%v,strVar2=%v", strVar1, strVar2)
	fmt.Println(len(strVar2))
	fmt.Println("Hello, " + "世界")

	var strVar3 = "Hello World, 我是你好"
	strVar3Length := len(strVar3)
	for index := 0; index < strVar3Length; index++ {
		fmt.Printf("%s--编码值=%d,值=%c\n", strVar3, strVar3[index], strVar3[index])
	}
	for _, value := range strVar3 {
		// value的类型rune int32
		fmt.Printf("%s--编码值=%d,值=%c,类型=%T\n", strVar3, value, value, value)
	}
}
