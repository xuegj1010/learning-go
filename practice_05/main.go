package main

import "fmt"

func main() {
	/*
		字符:
		byte(uint8) ascii的一个字符
		rune Unicode int32 utf-8

		字符存储过程：字符-》》ascii码值 -》》二进制
		字符读取过程：二进制-》》ascii码值-》》字符
	*/
	var charVar1 byte = '0'
	charVar2 := '薛'
	fmt.Printf("charVar1 = %d, charVar2 = %d\n", charVar1, charVar2)
	fmt.Printf("charVar2=%c,charVar2=%T\n", charVar2, charVar2)

	var1 := 'a' //对应的编码制
	fmt.Printf("加和=%d，a的编码制=%d\n", 100+var1, var1)
}
