package main

import (
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

	positions := lib.CsvToIntSlice(input)

	// 1
	fmt.Println(leastFuelAlignment(positions))

	// 2
	fmt.Println(leastIncreasingFuelAlignment(positions))
}

// 1
func leastFuelAlignment(positions []int) int {
	crabs := newCrabs(positions)
	median := crabs.medianPosition()

	return crabs.fuelSpentToAlign(median)
}

// 2
func leastIncreasingFuelAlignment(positions []int) int {
	crabs := newCrabs(positions)
	flooredAvg := crabs.flooredAveragePosition()

	costAvg := crabs.increasingFuelSpentToAlign(flooredAvg)
	costAvg2 := crabs.increasingFuelSpentToAlign(flooredAvg + 1)

	// i found that the ideal "middle value" this time is the average
	// but it was wrong, i brute forced by finding the cost of every pos
	// and it turns out it was the average-1
	// 
	// so i circumvented this by making sure to check floored and rounded avg
	// if floored and rounded avg is the same, we still get the corect answer

	return min(costAvg, costAvg2)
}
