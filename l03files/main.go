package main

import (
	"fmt"
	"os"

	"frontendmasters.com/gobasics/l03files/fileutils"
)

func main() {

	rootPath, err := os.Getwd()
	filepath := rootPath + "/data/text.txt"

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	c, err := fileutils.ReadTextFile(filepath)

	if err != nil {
		fmt.Printf("ERROR PANIC!!: %v\n", err)
	} else {
		fmt.Println(c)
		newContent := fmt.Sprintf("Original: %v\nDouble Original:\n%v\n%v\n", c, c, c)

		fileutils.WriteToFile(filepath+".output.txt", newContent)
	}
}
