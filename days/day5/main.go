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

	lines := parseLines(lib.ToStrSlice(input))

	// 1
	fmt.Println(getDangerousPointCount(lines))

	// 2
	fmt.Println(getDangerousPointCountWithDiagonals(lines))
}

// 1
func getDangerousPointCount(lines Lines) int {
	cols, rows := lines.furthestPoint()
	fmap := newFloorMap(rows, cols)

	straightLines := lines.filterStraightLines()
	for _, straightLine := range straightLines {
		fmap.drawLine(straightLine)
	}

	if TESTING {
		fmt.Println(fmap)
	}

	return fmap.getDangerousPointCount()
}

// 2
func getDangerousPointCountWithDiagonals(lines Lines) int {
	cols, rows := lines.furthestPoint()
	fmap := newFloorMap(rows, cols)

	for _, line := range lines {
		fmap.drawLine(line)
	}

	if TESTING {
		fmt.Println(fmap)
	}

	return fmap.getDangerousPointCount()
}
