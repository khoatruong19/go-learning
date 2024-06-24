package main

type IProduct interface {
	setName(value string)
	getName() string
	setPrice(value int32)
	getPrice() int32
}

type Product struct {
	name  string
	price int32
}

func (p *Product) setName(value string) {
	p.name = value
}

func (p *Product) getName() string {
	return p.name
}

func (p *Product) setPrice(value int32) {
	p.price = value
}

func (p *Product) getPrice() int32 {
	return p.price
}
