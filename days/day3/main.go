package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := lib.ReadTest()
	input := lib.ReadInput()

	fmt.Println(powerConsumption(input))
	fmt.Println(lifeSupport(input))
}

// 1
func binaryToDecimal(binary string) int {
	if decimal, err := strconv.ParseInt(binary, 2, 64); err != nil {
		panic(err)
	} else {
		return int(decimal)
	}
}

func invertBinaryString(binary string) string {
	inverted := strings.Builder{}

	for _, char := range binary {
		if char == '0' {
			inverted.WriteString("1")
		} else {
			inverted.WriteString("0")
		}
	}

	return inverted.String()
}

func mostCommonBits(bitCounts [][]int) string {
	numLen := len(bitCounts[0])
	result := strings.Builder{}

	for i := 0; i < numLen; i++ {
		if bitCounts[0][i] > bitCounts[1][i] {
			result.WriteString("0")
		} else {
			result.WriteString("1")
		}
	}

	return result.String()
}

func powerConsumption(diagnosis string) int {
	nums := lib.ToStrSlice(diagnosis)
	numLen := len(nums[0])

	// create a 2xLEN array, 1st row is for 0s, 2nd row is for 1s
	bitCounts := make([][]int, 2)
	for i := range bitCounts {
		bitCounts[i] = make([]int, numLen)
	}

	// populate the 2xLEN array of bit counts per column
	for _, num := range nums {
		for i, bit := range num {
			if bit == '0' {
				bitCounts[0][i]++
			} else if bit == '1' {
				bitCounts[1][i]++
			}
		}
	}

	gamma := mostCommonBits(bitCounts)
	epsilon := invertBinaryString(gamma)

	gammaRate := binaryToDecimal(gamma)
	epsilonRate := binaryToDecimal(epsilon)

	return gammaRate * epsilonRate
}

// 2
// priority is 1 or 0
// when count of 1 == 0, priority will be chosen
func mostCommonBitInColumn(column int, numbers []string, priority rune) rune {
	ones, zeroes := 0, 0

	for _, number := range numbers {
		if number[column] == '0' {
			zeroes++
		} else if number[column] == '1' {
			ones++
		}
	}

	if ones > zeroes {
		return '1'
	} else if zeroes > ones {
		return '0'
	} else {
		return priority
	}
}

func filterDiagnosis(column int, bit rune, diagnosis []string) []string {
	newDiagnosis := []string{}

	for _, number := range diagnosis {
		if number[column] == byte(bit) {
			newDiagnosis = append(newDiagnosis, number)
		}
	}

	return newDiagnosis
}

func lifeSupport(diagnosis string) int {
	nums := lib.ToStrSlice(diagnosis)

	// find oxygen
	i := 0
	buffer := nums
	for len(buffer) > 1 {
		mostCommonBit := mostCommonBitInColumn(i, buffer, '1')
		buffer = filterDiagnosis(i, mostCommonBit, buffer)
		i++
	}
	oxygen := buffer[0]

	// find co2
	i = 0
	buffer = nums
	for len(buffer) > 1 {
		mostCommonBit := mostCommonBitInColumn(i, buffer, '1')
		
		var leastCommonBit rune
		if mostCommonBit == '0' {
			leastCommonBit = '1'
		} else {
			leastCommonBit = '0'
		}

		buffer = filterDiagnosis(i, leastCommonBit, buffer)
		i++
	}
	co2 := buffer[0]

	oxygenRate := binaryToDecimal(oxygen)
	co2Rate := binaryToDecimal(co2)

	return oxygenRate * co2Rate
}
