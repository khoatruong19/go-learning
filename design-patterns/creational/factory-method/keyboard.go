package main

type Keyboard struct {
	Product
}

func newKeyboard() IProduct {
	return &Keyboard{
		Product: Product{
			name:  "keyboard",
			price: 100,
		},
	}
}
