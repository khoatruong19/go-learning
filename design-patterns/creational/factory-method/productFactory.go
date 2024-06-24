package main

import "fmt"

func getProduct(productType string) (IProduct, error) {
	if productType == "phone" {
		return newPhone(), nil
	}
	if productType == "keyboard" {
		return newKeyboard(), nil
	}
	return nil, fmt.Errorf("wrong product type passed")
}
