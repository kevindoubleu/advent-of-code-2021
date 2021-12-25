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

	note := newNote(lib.ToStrSlice(input))

	// 1
	fmt.Println(note.getEasyOutputCount())

	// 2
	fmt.Println(getNoteSum(note))
}

// 2
func getNoteSum(note Note) int {
	sum := 0

	for _, entry := range note.entries {
		sum += entry.parseOutput()
	}

	return sum
}
