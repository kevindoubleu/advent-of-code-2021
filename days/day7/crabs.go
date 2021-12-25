package main

import (
	"sort"
)

type Crabs struct {
	positions []int
}

func newCrabs(positions []int) *Crabs {
	return &Crabs{
		positions: positions,
	}
}

func (c Crabs) flooredAveragePosition() int {
	sum := 0

	for _, pos := range c.positions {
		sum += pos
	}

	return sum / len(c.positions)
}

func (c Crabs) medianPosition() int {
	sort.Ints(c.positions)

	return c.positions[ len(c.positions) / 2 ]
}

func (c Crabs) fuelSpentToAlign(position int) int {
	total := 0

	for _, pos := range c.positions {
		total += diff(pos, position)
	}

	return total
}

// 1 2 3 4 5 6 distance
// 1 3 6 10 15 fuel cost
// we can see from the simulation, the fuel costs creates a pattern
// googling this pattern we find its called the triangle pattern
func (c Crabs) increasingFuelSpentToAlign(position int) int {
	total := 0

	for _, pos := range c.positions {
		total += triangleSequence( diff(pos, position) )
	}

	return total
}
