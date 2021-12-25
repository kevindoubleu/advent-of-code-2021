package main

import (
	"fmt"
	"strings"
)

type Line struct {
	x1, y1 int
	x2, y2 int
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.x1, l.y1, l.x2, l.y2)
}

func parseLine(lineString string) Line {
	line := Line{}
	LineStringReader := strings.NewReader(lineString)
	fmt.Fscanf(LineStringReader, "%d,%d -> %d,%d", &line.x1, &line.y1, &line.x2, &line.y2)
	return line
}

func parseLines(lineStrings []string) Lines {
	lines := make(Lines, 0)

	for _, lineString := range lineStrings {
		lines = append(lines, parseLine(lineString))
	}

	return lines
}

func (l Line) isStraight() bool {
	return l.x1 == l.x2 || l.y1 == l.y2
}

func (l Line) distance() int {
	return max( difference(l.x1, l.x2), difference(l.y1, l.y2) )
}

type Lines []Line

func (ls Lines) String() string {
	result := strings.Builder{}

	for _, l := range ls {
		result.WriteString(l.String())
		result.WriteString("\n")
	}

	return result.String()[ : result.Len() - 1]
}

// get furthest point among some lines
func (ls Lines) furthestPoint() (int, int) {
	x, y := 0, 0

	for _, l := range ls {
		x = max( max(x, l.x1), max(x, l.x2) )
		y = max( max(y, l.y1), max(y, l.y2) )
	}

	return x, y
}

// returns straight lines from a mix of straight and non-straight
func (ls Lines) filterStraightLines() Lines {
	filtered := make(Lines, 0)

	for _, line := range ls {
		if line.isStraight() {
			filtered = append(filtered, line)
		}
	}

	return filtered
}
