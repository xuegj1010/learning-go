package model

type Product struct {
	productName  string
	productPrice float32
}

//getter setter

func (p *Product) SetProductName(_productName string) {
	p.productName = _productName
}

func (p *Product) GetProductName() string {
	return p.productName
}

func (p *Product) SetProductPrice(_productPrice float32) {
	p.productPrice = _productPrice
}

func (p *Product) GetProductPrice() float32 {
	return p.productPrice
}
