package main

import (
	"fmt"
	"learning-go/practice_13_pkg/model"
)

func main() {
	//面向对象
	//类型组合，在一个类型当中嵌入一个或多个类型来实现面向对象
	//go class struct
	//抽象的实体（类）与对象
	//结构体与结构变量
	//接口来增加实体的方法
	//面向对象的三大特征
	//封装，对外暴露公开的接口，增强安全，简化编程
	//继承，子类继承父类，自动拥有父类的属性和方法
	//多态，通过接口来实现
	//继承，重写，父类引用指向子类对象

	//boge := model.UserInfo{
	//	Name:      "guojun",
	//	Age:       20,
	//	Height:    178,
	//	EduSchool: "清华大学",
	//	Hobby:     []string{"coding", "运动"},
	//	MoreInfo:  nil,
	//}
	guojun := model.NewUserInfo("guojun", 20, 178, "清华大学", []string{"coding", "运动"}, nil)
	fmt.Printf("guojun的信息=%v\n", guojun)

	product := model.Product{}
	product.SetProductName("go语言高级")
	fmt.Println(product.GetProductName())

}
