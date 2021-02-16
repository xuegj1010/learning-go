package main

import "fmt"

var (
	varx       int
	slicex     []int
	interfacex interface{}
)

func main() {
	// 零值
	//int int32 int64
	//float32 float64
	//bool
	//string
	//var VarUpper = 19 // 变量名大写可以被其他包引入，小写只能被本包使用
	var var1 int
	var2 := "我是xueguojun" // 简短申明
	//_, x := 22,100
	var var3 = 100
	var var4, var5, var6, var7 = 1, "我是xueguojun", 3.14, true

	var i = 100
	var j = 200
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
	fmt.Println(var1, var2, var3, var4, var5, var6, var7)
}
