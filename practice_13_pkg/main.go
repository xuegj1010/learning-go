package main

import (
	"fmt"
	"learning-go/practice_13_pkg/action"
	"learning-go/practice_13_pkg/model"
)

func main() {
	// 包，package
	// 包名一般与文件夹名一致
	// 同一个包下，所有文件的包名是一致的

	// 大写开头的变量或函数可以被导出，可以 被其他包访问
	// 小写开头的变量或函数不可以被导出，不可以被其他包访问
	// 同一包中，不允许定义相同的变量和函数
	fmt.Println(action.Product)
	action.ProductList()

	bobo := model.UserInfo{
		Name:      "伯格",
		Age:       18,
		Height:    182.43,
		EduSchool: "清华大学",
		Hobby:     []string{"coding", "cooking", "旅游", "读书", "运动"},
		MoreInfo: map[string]interface{}{
			"work": "百度",
			"duty": "产品狗",
		},
	}
	fmt.Println(bobo.MoreInfo["duty"])

	//导入外部包，一个或多个外部的包，go get 支持github，bitbucket，google code
	//git, svn
	// go get 远程包名
	//1.下载远程的源码包
	//2.go install

	//新建一个包，在这个包中我们定义一个功能，只是简单的打印，
	//function 包
	//	function.go
	//		fmt.Println(xxxx)
}
