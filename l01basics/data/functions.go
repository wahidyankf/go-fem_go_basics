package data

import "fmt"

func save() {
	fmt.Println("Saving data")
}

func save2(text string) {
	fmt.Println("Saving data", text)
}

func add(a int, b int) int {
	return a + b
}

func increment(x *int) {
	*x++
}

func main() {
	i := 1
	increment(&i)
}
