package polygrp

import (
	"fmt"
	"strings"
)

func parseTemplate(template string) map[string]int {
	result := make(map[string]int)
	strLen := len(template)
	if strLen < 2 { return result }
	
	for i := 1; i < strLen; i++ {
		pair := template[i-1 : i+1]
		result[pair]++
	}

	return result
}

func parseRules(rules []string) map[string][]string {
	result := make(map[string][]string)

	for _, rule := range rules {
		// get the string parts
		var from, to string
		ruleReader := strings.NewReader(rule)
		fmt.Fscanf(ruleReader, "%s -> %s", &from, &to)

		// convert the "to" part into the 2 new pairs that will be generated
		newPair1 := fmt.Sprint(string(from[0]), to)
		newPair2 := fmt.Sprint(to, string(from[1]))

		result[from] = []string{newPair1, newPair2}
	}

	return result
}

func pairsToCharMap(pairs map[string]int) map[byte]int {
	charMap := make(map[byte]int)

	// for each pair
	for pair, count := range pairs {
		// separate the pair into its respective chars
		chars := []byte(pair)

		// for each char
		for _, char := range chars {
			// add the count of it to the charmap
			charMap[char] += count
		}
	}

	return charMap
}
