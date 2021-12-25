package main

import (
	"advent-of-code-2021/day25/cucumber"
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
	c := cucumber.NewCucumber(lib.ToStrSlice(input))
	fmt.Println(c.StepsToStuck())
	
	// 2
	
}
