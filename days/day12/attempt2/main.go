package main

import (
	"advent-of-code-2021/days/day12/attempt2/cave"
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
	cs := cave.NewCaveSystem(lib.ToStrSlice(input))
	fmt.Println(cs)
	
	// 2

}
