package main

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func difference(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}
