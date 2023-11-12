package main

func main() {
	var x int
	var name string

	const y = 2
	var z int = 2

	// var text string
	// text = "Hello!"
	// merged as below (due to the linter)
	var text = "Hello!"

	otherText := "bye!"

	println(x, name, y, z, text, otherText)
}
