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

	// 1
	fishes := newFishes(lib.CsvToIntSlice(input))
	fishes.age(80)
	fmt.Println(fishes.count())

	// 2
	fishgroup := newFishGroup(lib.CsvToIntSlice(input))
	fishgroup.age(256)
	fmt.Println(fishgroup.count())
}
