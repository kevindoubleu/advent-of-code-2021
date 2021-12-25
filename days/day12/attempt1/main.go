package main

import (
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

	cs := newCaveSystem(lib.ToStrSlice(input))
	// cs.DFS()

	// 1
	cs.allPaths()
	fmt.Println(len(cs.routes))

	// 2
	cs.allPaths2()
	fmt.Println(len(cs.routes))
}
