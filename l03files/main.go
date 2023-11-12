package main

import (
	"fmt"
	"os"

	"frontendmasters.com/gobasics/l03files/fileutils"
)

func main() {

	rootPath, err := os.Getwd()

	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")

	if err != nil {
		fmt.Printf("ERROR PANIC!!: %v\n", err)
	} else {
		fmt.Println(c)
	}
}
