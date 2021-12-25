package main

import (
	"errors"
	"fmt"
)

func getStringWithlen(strings []string, strLen int) string {
	for _, str := range strings {
		if len(str) == strLen {
			return str
		}
	}

	panic(errors.New(fmt.Sprint("multiple strings of len", strLen, "found")))
}

func _toCharMap(s string) map[byte]int {
	charmap := make(map[byte]int)

	for _, c := range s {
		charmap[byte(c)]++
	}

	return charmap
}

func stringDifference(a, b string) byte {
	mapA := _toCharMap(a)
	mapB := _toCharMap(b)

	for char, count := range mapA {
		if mapB[char] != count {
			return char
		}
	}

	return 0
}

// combine numbers like appending strings
func combineNumbers(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result *= 10
		result += number
	}

	return result
}
