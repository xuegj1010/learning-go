package model

import "fmt"

//微信支付
//支付宝
//银联
//银行卡

type PaymentArgs struct {
	AppID       string
	MchID       string
	Key         string
	CallbackUrl string
}

func (p *PaymentArgs) Info() {
	fmt.Printf("Info=%v\n", p)
}
