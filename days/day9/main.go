package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"sort"
)

var TESTING = false

func main() {
	var input string
	if TESTING {
		input = lib.ReadTest()
	} else {
		input = lib.ReadInput()
	}

	floorHeights := lib.To2DIntSlice(input)
	heightmap := newHeightmap(floorHeights)

	// 1
	fmt.Println(heightmap.getRisk())

	// 2
	fmt.Println(product3LargestBasins(heightmap))
}

// 2
func product3LargestBasins(m Heightmap) int {
	lowPoints := m.getLowPoints()
	basinSizes := []int{}

	for _, lowPoint := range lowPoints {
		basinSizes = append(basinSizes, m.getBasinSize(lowPoint))
	}

	sort.Ints(basinSizes)
	arrLen := len(basinSizes)
	product := basinSizes[arrLen-1] * basinSizes[arrLen-2] * basinSizes[arrLen-3]

	return product
}
