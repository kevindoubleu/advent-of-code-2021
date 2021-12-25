package main

import (
	"advent-of-code-2021/days/day22/core"
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

	instructions := lib.ToStrSlice(input)
	// 1
	c := core.NewCore()
	c.Initialize(instructions)
	fmt.Println(c.CountOn())
	
	// 2

}
