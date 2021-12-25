package lib

import (
	"os"
	"strings"
)

func ReadInput() string {
	return readFile("input.txt")
}

func ReadTest() string {
	return readFile("test.txt")
}

func readFile(filename string) string {
	// the file path is relative to where the caller file is at
	// not where this lib file is at
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Trim(string(content), " \n")
}
