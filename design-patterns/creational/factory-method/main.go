package main

func main() {
	phone, _ := getProduct("phone")
	keyboard, _ := getProduct("keyboard")

	phone.getName()
	phone.getPrice()
	keyboard.getName()
	keyboard.getPrice()
}
