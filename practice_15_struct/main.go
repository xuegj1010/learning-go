package main

import (
	"fmt"
	"unsafe"
)

type userinfo struct {
	name     string
	age      int
	height   float32
	edushool string
	hobby    []string
	moreinfo map[string]interface{}
}
type peopleinfo struct {
	name     string
	age      int
	height   float32
	edushool string
	hobby    []string
	moreinfo map[string]interface{}
}

type Integer int

//后台管理系统中，权限问题，这里涉及来角色，超级管理员，管理员，普通用户
type role struct {
	user          userinfo
	authorization Integer //1=超级管理员，2=普通管理员，3=普通用户
}

func main() {
	//结构体 struct
	//java，php，c++
	//定义自定义类型，两种方式：
	//1.type
	type integer int
	var intVar int
	var integerVar integer
	//intVar = integer // 类型不同不能相互赋值
	fmt.Println(intVar, integerVar)
	//如何赋值？类型转换
	intVar = int(integerVar)

	//2.struct,组合一定数量的字段完成，值类型
	//type 结构体名 struct {
	//	field1 type1
	//	field2 type2
	//	...
	//}

	//结构体的使用方式1
	var bob userinfo
	bob.name = "xueguo"
	bob.age = 18
	bob.height = 181
	bob.edushool = "清华大学"
	bob.hobby = []string{"coding", "运动", "旅游"}
	bob.moreinfo = map[string]interface{}{
		"work": "百度",
		"duty": "产品狗",
	}
	fmt.Printf("name=%s,hobby=%v\n", bob.name, bob.hobby)
	//结构体的使用方式2
	// :=简短声明来实现一个结构变量
	//	a.声明变量
	//	b.初始化
	//var liuge userinfo = userinfo{
	//	name:     "胡歌",
	//	age:      28,
	//	height:   188,
	//	edushool: "清华大学",
	//	hobby:    []string{"拍电影", "唱歌", "旅行"},
	//	moreinfo: map[string]interface{}{
	//		"role":     "演员",
	//		"earnmony": 3000000,
	//	},
	//}
	huge := userinfo{
		name:     "胡歌",
		age:      28,
		height:   188,
		edushool: "清华大学",
		hobby:    []string{"拍电影", "唱歌", "旅行"},
		moreinfo: map[string]interface{}{
			"role":     "演员",
			"earnmony": 3000000,
		},
	}
	fmt.Printf("huge=%v\n", huge)

	//不指定字段名的时候，需要严格的按照定义结构体时候的顺序赋值
	muse := userinfo{
		"胡歌",
		28,
		188,
		"清华大学",
		[]string{"拍电影", "唱歌", "旅行"},
		map[string]interface{}{
			"role":     "演员",
			"earnmony": 3000000,
		},
	}
	fmt.Printf("muse=%v\n", muse)

	xiaoge := userinfo{"小哥", 12, 120, "小学", []string{"学习", "玩", "打游戏"}, map[string]interface{}{"年级": "六年级"}}
	fmt.Printf("xiaoge=%v\n", xiaoge)

	//结构体的使用方式3 new new(int),new(string),new(T)结构体指针
	//var t = *T
	//t = new(T)
	var xiaoming *userinfo
	xiaoming = new(userinfo) // 返回的是一个指针
	(*xiaoming).name = "小明"
	(*xiaoming).age = 18
	(*xiaoming).edushool = "北京大学"
	//xiaoming --> (*xiaoming) go语言在底层自动转换
	xiaoming.hobby = []string{"学习", "玩", "打游戏"}
	fmt.Println(xiaoming)

	//结构体的使用方式4 &取地址符，同样时返回的结构体指针
	var xiaohong *userinfo = &userinfo{"小红", 12, 120, "小学", []string{"学习", "玩", "打游戏"}, map[string]interface{}{"年级": "六年级"}}
	fmt.Println(xiaohong, *xiaohong)
	//结构体的注意事项
	//a.结构体是值类型
	user1 := userinfo{}
	user2 := user1
	fmt.Printf("user1=%p,user2=%p\n", &user1, &user2) // 值类型地址不同
	//b.结构体之间师傅可以相互转换？可以转换，前提条件：具有相同字段（个数，类型，名称）
	user3 := userinfo{}
	user4 := peopleinfo{}
	//user3 = user4 // 类型不同直接转换会报错
	user3 = userinfo(user4) //强制转换
	fmt.Println(user3)

	//结构体可以作为另一个结构体字段的类型
	superadmin := role{user: userinfo{name: "超级管理员"}, authorization: 1}
	admin := role{user: userinfo{name: "管理员"}, authorization: 2}
	fmt.Println(superadmin, admin)

	//结构体在内存中的布局
	fmt.Printf("name=%p,age=%p,height=%p,eduschool=%p,hobby=%p,moreinfo=%p\n",
		&xiaohong.name, &xiaohong.age, &xiaohong.height, &xiaohong.edushool, &xiaohong.hobby, &xiaohong.moreinfo)
	// 字节大小
	fmt.Printf("name=%d,age=%d,height=%d,eduschool=%d,hobby=%d,moreinfo=%d\n",
		unsafe.Sizeof(xiaohong.name), unsafe.Sizeof(xiaohong.age), unsafe.Sizeof(xiaohong.height), unsafe.Sizeof(xiaohong.edushool), unsafe.Sizeof(xiaohong.hobby), unsafe.Sizeof(xiaohong.moreinfo))

	type Hobby struct {
		hobby1 string
		Hobby2 string
		Hobby3 string
	}
	hb := Hobby{"学习", "玩", "打游戏"}
	fmt.Printf("hobby1=%p,hobby2=%p,hobby3=%p\n", &hb.hobby1, &hb.Hobby2, &hb.Hobby3)
}
