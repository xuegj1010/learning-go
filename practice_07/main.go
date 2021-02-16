package main

import "fmt"

func main() {
	/*
		bool true false
	*/
	var boolVar11 bool
	boolVar11 = true
	boolVar12 := true == false

	fmt.Println(boolVar11, boolVar12)

	if 3 == 4 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
	fmt.Println(!true, !false, !!true)
}
