package main

import (
	"advent-of-code-2021/day24/alu"
	"advent-of-code-2021/lib"
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
	a := alu.NewALU(lib.ToStrSlice(input))
	a.BiggestModelNumber()
	
	// 2

}
