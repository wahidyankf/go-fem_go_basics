package fileutils

import "os"

func ReadTextFile(filename string) (string, error) {

	content, err := os.ReadFile(filename)
	if err != nil {
		// we could not read the file

		// first option
		// return ""

		// second option
		// panic("AHAHAHAHAHAHA")

		// third option
		return "", err

	} else {
		// read operation was successful
		return string(content), nil
	}

}
