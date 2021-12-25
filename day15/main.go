package main

import (
	"advent-of-code-2021/day15/cave"
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
	c := cave.NewCave(lib.To2DIntSlice(input))
	c.ShortestPath()
	
	// 2
	// checking if the increments are correct
	testCave := cave.NewCave([][]int{[]int{8}})
	testCave.Enlarge()
	fmt.Println(testCave)

	c.Enlarge()
	c.ShortestPath() // took 2 minutes
}
