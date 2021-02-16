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
