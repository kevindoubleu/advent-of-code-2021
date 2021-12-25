package main

import (
	"advent-of-code-2021/lib"
	"fmt"
)

func main() {
	// input := lib.ReadTest()
	input := lib.ReadInput()
	
	fmt.Println(largerThanPrev(input))
	fmt.Println(largerTripletThanPrev(input))
}

// 1
func largerThanPrev(measurements string) int {
	depths := lib.ToIntSlice(measurements)
	arrLen := len(depths)

	result := 0
	for i := 1; i < arrLen; i++ {
		if depths[i-1] < depths[i] {
			result++
		}
	}

	return result
}

// 2
func largerTripletThanPrev(measurements string) int {
	depths := lib.ToIntSlice(measurements)
	arrLen := len(depths)

	result := 0
	for i := 3; i < arrLen; i++ {
		sumPrev := depths[i-3] + depths[i-2] + depths[i-1]
		sum     := depths[i-2] + depths[i-1] + depths[i]

		if sumPrev < sum {
			result++
		}
	}

	return result
}
