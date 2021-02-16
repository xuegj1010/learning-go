package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map，集合，无序的键值对这种数据类型
		键值对
		map[key] = value
		1 3 5 6
		3 5 6 2
		通过散列函数对我们的map
		0xc0001
		0xc0002
		0xc0003

		map key value
	*/
	// map的定义
	// 1.申明
	var mapVar1 map[string]string
	// 2.make
	mapVar1 = make(map[string]string)
	mapVar1["Monday"] = "周一"
	mapVar1["Tuesday"] = "周二"
	fmt.Println(mapVar1)
	// map这里的key是唯一的，不能重复
	mapVar1["Monday"] = "大周一"
	fmt.Println(mapVar1)
	//map len
	fmt.Println(len(mapVar1))

	// 申明的时候同时赋值
	var mapVar2 = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
	}
	fmt.Println(mapVar2)

	// 通过 :=
	mapVar3 := make(map[string]int)
	mapVar3["Monday"] = 1
	mapVar3["Tuesday"] = 2
	mapVar3["Wednesday"] = 3
	fmt.Println(mapVar3)

	//切片，函数不能作为map的key
	//mapVar4 := map[[]string]int{}

	mapVar5 := map[int][]string{}
	var sliceString []string
	sliceString = make([]string, 3)
	sliceString[0] = "xueguojun"
	sliceString[1] = "golang"
	sliceString[2] = "世界"

	mapVar5[0] = sliceString
	fmt.Println(mapVar5)

	// 只申明没有make，报错
	//var mapVar6 map[string]string
	//mapVar6["hello"] = "shijie"

	// map是一个引用类型
	fmt.Printf("mapVar3['Monday']在调用函数之前的值:%v\n", mapVar3["Monday"])
	changeMap(mapVar3)
	fmt.Printf("mapVar3['Monday']在调用函数之后的值:%v\n", mapVar3["Monday"])

	// 如何来对map的key进行顺序的输出， for range
	//mapVar3["Monday"] = 1
	//mapVar3["Tuesday"] = 2
	//mapVar3["Wednesday"] = 3

	var sortKeys []string
	for key := range mapVar3 {
		sortKeys = append(sortKeys, key)
	}
	fmt.Printf("排序前sortKeys的值=%v\n", sortKeys)

	// sort
	sort.Strings(sortKeys)
	fmt.Printf("排序后sortKeys的值=%v\n", sortKeys)

	for i := 0; i < len(sortKeys); i++ {
		fmt.Printf("mapVar3[%s]=%d\n", sortKeys[i], mapVar3[sortKeys[i]])
	}

	// 结构体和c类似
	type course struct {
		courseName    string  // 课程名称
		courseTime    float32 // 课程时长 单位分钟
		courseTeacher string  // 课程讲师
	}

	courser1 := course{
		courseName:    "go语言体系课",
		courseTime:    300.3,
		courseTeacher: "波哥",
	}
	courser2 := course{
		courseTeacher: "波哥",
		courseTime:    100.2,
		courseName:    "如何变帅",
	}

	courser := make(map[string]course)
	courser["go"] = courser1
	courser["美容"] = courser2
	fmt.Println(courser)

	value, ok := courser["go"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no value")
	}

	if value, ok := courser["go"]; ok {
		fmt.Println(value)
	}

	// map切片 map + 切片
	var mapVar6 []map[string]interface{}
	mapVar6 = make([]map[string]interface{}, 2)
	mapVar6[0] = make(map[string]interface{}, 2)
	mapVar6[0]["name"] = "波哥"
	mapVar6[0]["age"] = 18

	mapVar6[1] = make(map[string]interface{}, 2)
	mapVar6[1]["name"] = "波哥"
	mapVar6[1]["age"] = 16

	fmt.Println(mapVar6)
}

func changeMap(mapVar map[string]int) {
	mapVar["Monday"] = 888
}
