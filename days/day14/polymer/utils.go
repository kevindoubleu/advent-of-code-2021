package polymer

import (
	"fmt"
	"strings"
)

func parseRule(rule string) (from string, to byte) {
	ruleReader := strings.NewReader(rule)
	fmt.Fscanf(ruleReader, "%s -> %c", &from, &to)

	return from, to
}

func toCharMap(str string) map[byte]int {
	result := make(map[byte]int)

	for _, char := range str {
		result[byte(char)]++
	}

	return result
}
