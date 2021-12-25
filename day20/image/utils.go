package image

import (
	"strconv"
	"strings"
)

func parseImageAlgo(algoStr string) []bool {
	strLen := len(algoStr)
	algo := make([]bool, strLen)

	for i, char := range algoStr {
		if char == '#' {
			algo[i] = true
		}
	}

	return algo
}

func parseImageString(pixelsStr []string) [][]bool {
	height := len(pixelsStr)
	width := len(pixelsStr[0])

	pixels := make([][]bool, height)

	for i, row := range pixelsStr {
		pixelRow := make([]bool, width)

		for j, pixel := range row {
			if pixel == '#' {
				pixelRow[j] = true
			}
		}

		pixels[i] = pixelRow
	}

	return pixels
}

func (i Image) String() string {
	result := strings.Builder{}

	for _, row := range i.pixels {
		for _, pixel := range row {
			if pixel {
				result.WriteString("#")
			} else {
				result.WriteString(".")
			}
		}
		result.WriteString("\n")
	}

	return result.String()[: result.Len()-1 ]
}

// converts binary []bool to decimal
func bin2dec(binary []bool) int {
	binStr := strings.Builder{}

	for _, binValue := range binary {
		if binValue {
			binStr.WriteString("1")
		} else {
			binStr.WriteString("0")
		}
	}

	result, err := strconv.ParseInt(binStr.String(), 2, 0)
	if err != nil {
		panic(err)
	}

	return int(result)
}

func (i Image) LivePixelCount() int {
	// if all outside pixels are on, then the answer is infinite
	if i.outside { panic("answer is infinite") }

	count := 0

	for _, row := range i.pixels {
		for _, pixel := range row {
			if pixel {
				count++
			}
		}
	}

	return count
}
