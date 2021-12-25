package main

import (
	"advent-of-code-2021/days/day18/attempt1/sfnum"
	"advent-of-code-2021/lib"
	"fmt"
)

var TESTING = true

func main() {
	var input string
	if TESTING {
		input = lib.ReadTest()
	} else {
		input = lib.ReadInput()
	}

	// 1
	n := sfnum.AddMultiple(lib.ToStrSlice(input))
	fmt.Println(n)

	// 2

}
