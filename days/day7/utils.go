package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b int) int {
	if a < b {
	   return b - a
	}
	return a - b
}

// returns the n-th element in the triangle sequence
func triangleSequence(n int) int {
	lastNumber := 0

	for i := 1; i <= n; i++ {
		lastNumber = lastNumber + i
	}

	return lastNumber
}
