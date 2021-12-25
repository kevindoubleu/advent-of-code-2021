package main

import "math"

func difference(a, b int) int {
	return int(math.Abs(float64(a) - float64(b)))
}