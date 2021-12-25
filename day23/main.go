package main

import (
	"advent-of-code-2021/day23/burrow"
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
	b := burrow.NewBurrow(input)
	fmt.Println(b)
	b.AllScenarios(0)
	// answer is 14467 from manual simulation on paper
	
	// 2
	// couldnt get algo to work
	// lets try manually first
	// 
	// emptying 4th room first
	// 55405 too high
	// 
	// emptying middle 2 rooms
	// both A to the left
	// 48799 wrong
	// the 2 As are split left and right
	// 49499 wrong
	// 48759 correct
}
