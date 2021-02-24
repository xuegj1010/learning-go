package model

import "fmt"

type Alipay struct {
	PaymentArgs               // 匿名结构体
	PaymentOther PaymentOther //有名结构体
	AlipayOpenID string
}

func (a *Alipay) Info() {
	fmt.Printf("alipay=%v\n", a)
}
