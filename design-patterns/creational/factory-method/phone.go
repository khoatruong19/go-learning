package main

type Phone struct {
	Product
}

func newPhone() IProduct {
	return &Phone{
		Product: Product{
			name:  "iphone",
			price: 100,
		},
	}
}
