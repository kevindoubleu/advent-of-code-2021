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

	// 1
	octopi := newOctopi(lib.To2DIntSlice(input))
	fmt.Println(octopi.steps(100))

	// 2
	octopi = newOctopi(lib.To2DIntSlice(input))
	fmt.Println(octopi.stepsToSync())
}
