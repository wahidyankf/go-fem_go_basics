package data

import "fmt"

var Countries [10]string
var Slice []int
var Codes map[int]bool

func init() {
	Countries = [10]string{"USA", "Canada", "UK", "France", "Germany", "Italy", "Japan", "China", "India", "Brazil"}

	_ = len(Countries)

	fmt.Println("Coutries saved")
}
