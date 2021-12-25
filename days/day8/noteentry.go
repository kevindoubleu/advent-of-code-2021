package main

import (
	"strings"
)

type NoteEntry struct {
	digits		[]string
	outputs		[]string
	controller	Controller
}

func newNoteEntry(entry string) NoteEntry {
	parts := strings.Split(entry, " | ")
	
	digits := strings.Split(parts[0], " ")
	outputs := strings.Split(parts[1], " ")

	return NoteEntry{
		digits: digits,
		outputs: outputs,
		controller: newController(),
	}
}

func (e NoteEntry) _getCharCounts() map[byte]int {
	charCounts := make(map[byte]int)

	for _, digit := range e.digits {
		for _, char := range digit {
			charCounts[byte(char)]++
		}
	}

	return charCounts
}

// map each letter to each segment
func (e *NoteEntry) deduce() {
	// to ease the deduction process
	// the 7 segments are labeled as such
	//  111
	// 2   3
	//  444
	// 5   6
	//  777

	// i found an interesting point when looking at how many times
	// each segment gets lit up when we go through each of the number 0-9
	// 
	// using the frequency of each letter, we can deduce some segments
	// segment 1 appears in 8 numbers: 0 2 3 5 6 7 8 9
	// segment 2 appears in 6 numbers: 0 4 5 6 8 9
	// segment 3 appears in 8 numbers: 0 1 2 3 4 7 8 9
	// segment 4 appears in 7 numbers: 2 3 4 5 6 8 9
	// segment 5 appears in 4 numbers: 0 2 6 8
	// segment 6 appears in 9 numbers: 0 1 3 4 5 6 7 8 9
	// segment 7 appears in 7 numbers: 0 2 3 5 6 8 9
	// from here we can immediately find
	// the character that corresponds to segments 2,5,6
	// because they each have unique frequency of their respective letter

	charCounts := e._getCharCounts()
	for char, count := range charCounts {
		switch count {
		case 6:
			e.controller.assignSignalToSegment(char, 2)
		case 4:
			e.controller.assignSignalToSegment(char, 5)
		case 9:
			e.controller.assignSignalToSegment(char, 6)
		}
	}

	// now we have segments 2,5,6
	
	// segment 1 can be deduced by
	// finding the character that exists in the only 3 char string (7)
	// but doesnt exist in the only 2 char string

	number7string := getStringWithlen(e.digits, 3)
	number1string := getStringWithlen(e.digits, 2)
	diff1and7 := stringDifference(number7string, number1string)
	e.controller.assignSignalToSegment(diff1and7, 1)

	// now we have segments 1,2,5,6

	// segment 3 can be deduced by
	// finding the only character we havent deduced in the only 3 char string (7)

	for _, char := range number7string {
		if e.controller.signalToSegment[byte(char)] == 0 {
			e.controller.assignSignalToSegment(byte(char), 3)
			break
		}
	}

	// now we have segments 1,2,3,5,6

	// segment 4 can be deduced by
	// finding the only character in the number 4 (the only 4 char string)
	// that we dont know yet

	number4string := getStringWithlen(e.digits, 4)
	for _, char := range number4string {
		if e.controller.signalToSegment[byte(char)] == 0 {
			e.controller.assignSignalToSegment(byte(char), 4)
			break
		}
	}
		
	// now we have segments 1,2,3,4,5,6

	// we know the only letter we dont know is segment 7
	// and we know what letter we dont know because the letters are always a-g
	// or
	// we find the only character in the number 8 (the only 7 char string)
	// that we dont know yet

	number8string := getStringWithlen(e.digits, 7)
	for _, char := range number8string {
		if e.controller.signalToSegment[byte(char)] == 0 {
			e.controller.assignSignalToSegment(byte(char), 7)
			break
		}
	}
}

func (e NoteEntry) parseOutput() int {
	e.deduce()

	outputs := []int{}

	for _, outputSignal := range e.outputs {
		e.controller.displaySignal(outputSignal)
		outputs = append(outputs, e.controller.readDisplay())
		e.controller.resetDisplay()
	}

	return combineNumbers(outputs)
}
