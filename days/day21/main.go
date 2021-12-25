package main

import (
	"advent-of-code-2021/days/day21/dirac"
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
	
	start1, start2 := dirac.ParsePlayerPositions(lib.ToStrSlice(input))

	// 1
	practice := dirac.NewPracticeGame(start1, start2, 1000)
	practice.Play()
	fmt.Println(practice.LoserTimesDiceRolls())
	
	// 2
	
}
