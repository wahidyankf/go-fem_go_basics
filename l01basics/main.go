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

	// defer is LIFO
	// it will also executed when panic
	defer fmt.Println("Bye!!")
	defer fmt.Println("Good")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

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

	sTax, cTax := calculateTaxWithName(100)
	fmt.Println(sTax, cTax)

	fmt.Println("-----")

	age := 22
	birthday(age)
	birthday(age)
	// 22
	fmt.Println(age)

	birthday2(&age)
	// 23
	fmt.Println(age)

	// The pointer is 0xc0000b4008 and the value is 23
	birthday3(&age)

	// 24
	fmt.Println(age)

	fmt.Println("-----")

	panicking()
}

func panicking() {
	panic("Something bad happened")
}

func birthday3(pointerAge *int) {
	fmt.Printf("The pointer is %v and the value is %v\n", pointerAge, *pointerAge)

	*pointerAge++
}

func birthday2(age *int) {
	*age++
}

func birthday(age int) {
	_ = age + 1
}

func calculateTax(price float32) (float32, float32) {
	return price * 0.09, price * 0.02
}

func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price * 0.09
	cityTax = price * 0.02

	return
}
