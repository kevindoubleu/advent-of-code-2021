package main

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

func parseCoordinate(str string) Coordinate {
	parts := strings.Split(str, ",")

	x, err := strconv.Atoi(parts[0])
	panicIf(err)

	y, err := strconv.Atoi(parts[1])
	panicIf(err)

	c := Coordinate{
		x: x,
		y: y,
	}
	return c
}

func parseCoordinates(strs []string) []Coordinate {
	coords := []Coordinate{}

	for _, coordStr := range strs {
		coords = append(coords, parseCoordinate(coordStr))
	}

	return coords
}

// returns max x and max y of the coord(s) that are furthest from 0
// not the coord that has the biggest x + y
func getFurthestCoordinate(coords []Coordinate) (x, y int) {
	maxX, maxY := 0, 0

	for _, coord := range coords {
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	return maxX, maxY
}
