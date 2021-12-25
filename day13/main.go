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

	var coords, folds []string
	lib.UnpackToStrSlices(input, &coords, &folds)

	// 1
	paper1 := newPaper(coords)
	paper1.fold(folds[0])
	fmt.Println(paper1.dotCount())
	
	// 2
	paper2 := newPaper(coords)
	paper2.foldMultiple(folds)
	fmt.Println(paper2)
}
