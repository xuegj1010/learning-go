package action

import "fmt"

//public
var Product int

//private
var product int

func ProductList() {
	fmt.Println("action->ProductList()")
}

func productList() {
	fmt.Println("action->productList()")
}
