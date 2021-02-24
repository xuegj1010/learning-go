package main

import (
	"fmt"
	"learning-go/practice_13_pkg/model"
	model2 "learning-go/practice_17_oop/model"
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
	//继承，子类继承父类，子类自动拥有父类的属性和方法
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

	alipay := &model2.Alipay{
		PaymentArgs: model2.PaymentArgs{
			AppID:       "alipay123",
			MchID:       "alipaymchid",
			Key:         "alipayfjkadsfjlsjlsjfl",
			CallbackUrl: "https://api.imooc.com/alipay",
		},
		AlipayOpenID: "alipayopenid",
	}
	weixinpay := &model2.WeixinPay{
		PaymentArgs: model2.PaymentArgs{
			AppID:       "weixinpay123",
			MchID:       "weixinpaymchid",
			Key:         "weixinpayfjkadsfjlsjlsjfl",
			CallbackUrl: "https://api.imooc.com/weixinpay",
		},
		WeixinOpenID: "weixinpayopenid",
	}
	fmt.Println(alipay.PaymentArgs.AppID)
	fmt.Println(weixinpay.WeixinOpenID)

	//继承的优点
	// 1.提高代码的复用率
	// 2.提高代码的扩展性与维护性
	//fmt.Println(alipay.AppID)
	alipay.PaymentOther.AppID = "alipay_paymentother"
	fmt.Println(alipay.PaymentOther.AppID)
	fmt.Println(alipay.PaymentArgs.AppID)
	//访问流程
	//1.先判断AppID是否属于Alipay，如果有就访问
	//2.如果没有，继续去找他继承的结构体PaymentArgs，如果有就访问
	//3.如果没有继续去找PaymentArgs这个结构体，如果没有就访问，没有就报错
	//4.如果一个结构体继承来多个结构体，而这些多个结构体当中有相同的字段，那么我们就要用结构体名来访问
	//5.一个内置类型可以作为结构体的匿名字段，这种方式只能在本包访问
	ss := StringStruct{"hello"}
	fmt.Println(ss.string)

	//方法支持重载
	paymentargs := model2.PaymentArgs{
		AppID:       "superAppID",
		MchID:       "superMchID",
		Key:         "superKey",
		CallbackUrl: "https://api.imooc.com/super",
	}

	paymentargs.Info()
	alipay.Info()

	//多态 通过接口来实现
	//type 接口名 interface {
	// 方法名1(参数列表)(返回值)
	// 方法名2(参数列表)(返回值)
	// 方法名3(参数列表)(返回值)
	//。。。
	//}
	//方法接收者
	//type T struct {
	//
	//}
	//type TT string
	//func (t *TT)方法名1() (返回值){}
	//func (t *TT)方法名2() (返回值){}
	//func (t *TT)方法名3() (返回值){}

	//接口的说明：
	//	定义接口时，方法不能实现
	//一个类型如果实现来接口中的所有方法，我们就说这种类型实现接口，不仅结构体，也可以时内置类型
	// 自定义类型可以实现多个接口
	_payment := &payment{paymentmethod: "alipay"}
	_payment.info()
	_payment.topay()

	var _pay pay
	_pay = _payment //一个变量实现来接口当中所有方法，接口就可以指向这个变量
	_pay.info()
	_pay.topay()

	var _readwrite readwrite
	_readwrite.echo()

	//多重继承接口，所有的方法都要实现
	//InterfaceAA
	//	InterfaceA
	//	InterfaceB

	//多态
	//父类引用指向子类对象
	//一个变量实现来接口当中所有方法，接口就可以指向这个变量
	log := &Log{
		name:    "微信小程序支付日志",
		content: "微信小程序支付日志内容",
		addtime: 0,
	}
	var _iowrite *iowrite
	var _netwrite *netwrite
	log.writelog(_iowrite)
	log.writelog(_netwrite)

	//多态的最佳实践
	netlog := &NetLog{Log{
		name:    "微信小程序网络支付日志",
		content: "微信小程序网络支付日志内容",
		addtime: 0,
	}}
	netlog.writelog(_netwrite)
	filelog := &IOLog{Log{
		name:    "微信小程序文件支付日志",
		content: "微信小程序文件支付日志内容",
		addtime: 0,
	}}
	filelog.writelog(_iowrite)

}

type Log struct {
	name    string
	content string
	addtime int64
}

func (l *Log) writelog(_write write) {
	fmt.Println(l.name + "---" + l.content)
	_write.echo()
	_write.out()
}

type NetLog struct {
	Log
}

type IOLog struct {
	Log
}

type write interface {
	echo()
	out()
}

type iowrite string

func (i *iowrite) echo() {
	fmt.Println("iowrite:echo()")
}

func (i *iowrite) out() {
	fmt.Println("iowrite:out()")
}

type netwrite string

func (n *netwrite) echo() {
	fmt.Println("netwrite:echo()")
}

func (n *netwrite) out() {
	fmt.Println("netwrite:out()")
}

//自定义类型，基于内置类型
type readwrite string

//自定义类型可以实现多个接口
type read interface {
	scan()
	input()
}

//作业：完成这里的多继承接口的实现
//在多重继承当中，一个类要实现所有接口的方法
type readerwriter interface {
	write
	read
}

type Integer int

func (r *readwrite) echo() {
	fmt.Println("readwrite:echo")
}

func (r *readwrite) out() {
	fmt.Println("readwrite:out")
}

func (r *readwrite) scan() {
	fmt.Println("readwrite:scan()")
}
func (r *readwrite) input() {
	fmt.Println("readwrite:input()")
}

type pay interface {
	topay()
	info()
}

type payment struct {
	paymentmethod string
}

func (p *payment) topay() {
	fmt.Println("today:", p.paymentmethod)
}

func (p *payment) info() {
	fmt.Println("info:", p.paymentmethod)
}

type StringStruct struct {
	string
}
