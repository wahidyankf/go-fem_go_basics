package main

import (
	"fmt"

	"frontendmasters.com/gobasics/l01basics/data"
)

// global variable
var url = "http://google.com"
var name = "Frontend Masters"

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func main() {
	// function-scoped variables
	message := "Hello from Go"
	price := 34.4

	println(message, price)

	fmt.Println(data.MaxSpeed)

	printData()

	fmt.Println("-----")

	fmt.Println(data.Countries)

	fmt.Println("-----")

	stateTax, _ := calculateTax(100)
	fmt.Println(stateTax)
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}
