package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		流程控制
		条件控制：if else else if
		选择控制：switch case select fallthrough
		循环控制：for for range
		跳转控制 goto break continue

	*/

	// 条件控制
	if true {
		fmt.Println("true")
	} else if false {
		fmt.Println("false")
	} else {
		fmt.Println("other")
	}

	//if 的另一种使用场景
	//if value,ok := courser["go"];ok{
	//	fmt.Println(value)
	//}

	// 打开某一个文件不存在或者由于权限的问题无法操作
	if fileHandle, err := os.Open("hello.txt"); err != nil {
		fmt.Println(fileHandle)
		fmt.Println(err.Error())
	} else {
		fmt.Println("获取文件成功")
	}

	if true && false {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}

	// 选择控制 switch case select（后妈讲解通道的时候讲）
	// break不需要
	// fallthrough穿透一级向下一级执行
	switch switchVar := 100; {
	case switchVar > 90:
		fmt.Println("成绩优秀")
		fallthrough
	case switchVar > 80:
		fmt.Println("成绩良好")
		fallthrough
	case switchVar >= 60:
		fmt.Println("成绩合格")
		fallthrough
	default:
		fmt.Println("成绩不及格")
	}
	// if  else
	var switchVar2 int = 100
	//if switchVar2 > 100 {
	//
	//} else {
	//
	//}
	switch {
	case switchVar2 > 90:
		fmt.Println("成绩优秀")
	case switchVar2 > 80:
		fmt.Println("成绩良好")
	case switchVar2 >= 60:
		fmt.Println("成绩合格")
	default:
		fmt.Println("成绩不合格")
	}

	// switch判断类型
	var interfaceVar interface{}
	var floatVar = 3.1415926
	interfaceVar = floatVar
	switch typeVar := interfaceVar.(type) {
	case float32:
		fmt.Println("float32", typeVar)
	case float64:
		fmt.Println("float32", typeVar)
	case int8:
		fmt.Println("int8", typeVar)
	default:
		fmt.Println("未知的类型", typeVar)
	}

	// 循环控制 for for range
	// for 初始化条件;循环的条件;变化的条件
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	initVar := 1
	for initVar < 10 {
		fmt.Print(initVar)
		initVar++
	}

	//for {
	//	fmt.Print(1)
	//}

	var stringVar = "go语言体系课"
	for key, value := range stringVar {
		fmt.Println(key, value)
		fmt.Printf("key=%d,value=%c,类型=%T\n", key, value, value)
	}

	// 跳转控制
	//gotoFunc()
	for i := 0; i < 10; i++ {
		if i > 6 {
			break
		}
		fmt.Println(i)
	}
	// break的结果
	//for outerIndex := 0; outerIndex < 10; outerIndex++ {
	//	for innerIndex := 0; innerIndex < 6; innerIndex++ {
	//		fmt.Println(outerIndex, innerIndex)
	//		break
	//	}
	//}
	fmt.Println("xueguojun")
	// break+退出标签
FORBREAK:
	for outerIndex := 0; outerIndex < 10; outerIndex++ {
		for innerIndex := 0; innerIndex < 6; innerIndex++ {
			fmt.Println(outerIndex, innerIndex)
			break FORBREAK // 通过标签直接退出到最外层循环
		}
	}
	fmt.Println("xueguojun")
	//continue的结果
	for outerIndex := 0; outerIndex < 3; outerIndex++ {
		for innerIndex := 0; innerIndex < 2; innerIndex++ {
			fmt.Println(outerIndex, innerIndex)
			continue
			fmt.Println("这里是内层循环剩余的代码")
		}
	}
	fmt.Println("xueguojun")

	// continue+标签执行的结果
FORCONTINUE:
	for outerIndex := 0; outerIndex < 3; outerIndex++ {
		for innerIndex := 0; innerIndex < 2; innerIndex++ {
			fmt.Println(outerIndex, innerIndex)
			continue FORCONTINUE
		}
	}

}

func gotoFunc() {
	var1 := 0
GOTO:
	fmt.Println(var1)
	var1++
	goto GOTO
}
