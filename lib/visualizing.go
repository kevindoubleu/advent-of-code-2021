package lib

import (
	"fmt"
	"strings"
)

func Print2DIntSlice(slice [][]int) string {
	result := strings.Builder{}

	for _, row := range slice {
		result.WriteString(fmt.Sprint(row))
		result.WriteString("\n")
	}

	return result.String()[: result.Len()-1 ]
}