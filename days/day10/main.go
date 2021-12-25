package main

import (
	"advent-of-code-2021/lib"
	"fmt"
)

var TESTING = false

func main() {
	var input string
	if TESTING {
		input = lib.ReadTest()
	} else {
		input = lib.ReadInput()
	}

	s := newSyntaxChecker(lib.ToStrSlice(input))

	// 1
	fmt.Println(s.totalCorruptedScore())

	// 2
	fmt.Println(s.totalIncompleteScore())
}
