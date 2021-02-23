package main

import "fmt"

type miniaction struct {
	name   string //接口名称 热销商品，某一个分类下的商品列表
	router string //路由地址 https://api.mooc.com/hotproduct
	action string //路由对应的方法名 func houtproduct(){}
}

func (mini miniaction) getMiniInfo() {
	fmt.Printf("mini.name=%s,mini.router=%s,mini.action=%s\n", mini.name, mini.router, mini.action)
}

func (mini *miniaction) miniInfo() {
	fmt.Printf("mini.name=%s,mini.router=%s,mini.action=%s\n", mini.name, mini.router, mini.action)
}

type float float32

//获取圆的周长
func (radius float) getCircleAround() float {
	return radius * 3.14 * 2
}

func main() {
	//方法，method，有接收者的函数被称为方法
	//func (var vartype) funName(参数列表)(返回值列表){
	//		方法体
	//		return xxxx
	//}
	mini1 := miniaction{
		name:   "获取商品列表",
		router: "productList",
		action: "productList()",
	}
	mini2 := &miniaction{
		name:   "获取分类下的商品列表",
		router: "productListByCategory",
		action: "productListByCategory()",
	}
	mini1.getMiniInfo()
	//接收者有两种，一种是值类型，一种时指针类型
	//接收者是值类型的时候，调用者可以时值类型，也可以时引用类型
	mini2.getMiniInfo()

	//接收者是一个指针类型，调用者可以是值类型，也可以时指针类型
	mini2.miniInfo()
	(*mini2).miniInfo()
	mini1.miniInfo()

	//方法注意事项：
	//1.定义的方法只能通过指定的类型来调用，不能像函数一样调用
	//2.方法的接收者对应的变量名，习惯使用this，self
	//3.自定义类型也可以作为方法的接收者
	var radius float = 2
	fmt.Println(radius.getCircleAround()) // 获取圆的周长

	//方法和函数的区别
	//定义方式不同
	//作为参数使用，值类型，引用类型
	//方法:调用者.方法名(参数)
	//函数：函数名(参数列表)
	//首字母大小写，遵循包的规则

	//作业：定义一个结构体，为这个结构体定义一个方法，这一个方法实现两个整数的+-×/%
}
