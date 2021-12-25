package main

import (
	"advent-of-code-2021/day18/attempt2/sfnum"
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

	fmt.Println(input)
	// 1
	n := sfnum.NewSFNumber(input)
	fmt.Println(n)
	
	// 2

}
