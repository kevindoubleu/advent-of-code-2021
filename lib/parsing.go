package lib

import (
	"strconv"
	"strings"
)

// splits newline separated string into strings
func ToStrSlice(fileContent string) []string {
	return strings.Split(fileContent, "\n")
}

// splits newline separated string into ints
func ToIntSlice(fileContent string) []int {
	lines := strings.Split(fileContent, "\n")
	nums := []int{}

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	return nums
}

func CsvToIntSlice(fileContent string) []int {
	return ToIntSlice(strings.ReplaceAll(fileContent, ",", "\n"))
}

// single digits only
func To2DIntSlice(fileContent string) [][]int {
	lines := ToStrSlice(fileContent)
	slice := [][]int{}

	for i, line := range lines {
		slice = append(slice, []int{})
		
		for _, numStr := range line {
			if num, err := strconv.Atoi(string(numStr)); err != nil {
				panic(err)
			} else {
				slice[i] = append(slice[i], num)
			}
		}
	}

	return slice
}
