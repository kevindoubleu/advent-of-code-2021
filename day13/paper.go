package main

import (
	"fmt"
	"strings"
)

type Paper struct {
	dots	[]Coordinate
}

func newPaper(coordinateStrings []string) Paper {
	p := Paper{
		dots: parseCoordinates(coordinateStrings),
	}
	return p
}

func (p Paper) String() string {
	// prepare the plot
	maxX, maxY := getFurthestCoordinate(p.dots)
	plot := [][]bool{}
	for i := 0; i < maxY+1; i++ {
		plot = append(plot, make([]bool, maxX+1))
	}

	// mark the coords
	for _, dot := range p.dots {
		plot[dot.y][dot.x] = true
	}

	// draw it on a string builder
	result := strings.Builder{}
	for _, line := range plot {
		lineOnPaper := strings.Builder{}

		for _, point := range line {
			if point {
				lineOnPaper.WriteString("#")
			} else {
				lineOnPaper.WriteString(".")
			}
		}

		result.WriteString(lineOnPaper.String())
		result.WriteString("\n")
	}

	return result.String()[: result.Len()-1 ]
}

func (p *Paper) fold(instruction string) {
	axis, value := parseFold(instruction)

	switch axis {
	case 'x':
		fmt.Println("folding at x", value)
		// for all coords at x more than value
		for i := range p.dots {
			coord := &p.dots[i]

			// the problem says there will never be a dot on the fold
			if coord.x > value {
				
				// find their distance to the x fold value
				distance := coord.x - value

				// their new coord is the same distance from the fold
				// but instead of on the right, its now on the left
				coord.x = value - distance
			}
		}

	case 'y':
		fmt.Println("folding at y", value)

		// for all coords at y more than value
		for i := range p.dots {
			coord := &p.dots[i]

			// the problem says there will never be a dot on the fold
			if coord.y > value {
				
				// find their distance to the y fold value
				distance := coord.y - value

				// their new coord is the same distance from the fold
				// but instead of below, its now above the fold
				coord.y = value - distance
			}
		}
	}
}

// 1
func (p Paper) dotCount() int {
	return strings.Count(p.String(), "#")
}

// 2
func (p *Paper) foldMultiple(instructions []string) {
	for _, instruction := range instructions {
		p.fold(instruction)
	}
}
