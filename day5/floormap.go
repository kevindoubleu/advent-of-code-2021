package main

import (
	"fmt"
	"strings"
)

type FloorMap struct {
	state [][]int
}

func (m FloorMap) String() string {
	result := strings.Builder{}

	for _, row := range m.state {
		for _, cell := range row {
			if cell == 0 {
				result.WriteString(".")
			} else {
				result.WriteString(fmt.Sprint(cell))
			}
		}
		result.WriteString("\n")
	}

	return result.String()[ : result.Len() - 1]
}

// initialize a floor map
func newFloorMap(rows, cols int) FloorMap {
	fm := FloorMap{
		state: make([][]int, rows+1),
	}
	for i := range fm.state {
		fm.state[i] = make([]int, cols+1)
	}

	return fm
}

// increments the values in the cells in the line by 1
func (m *FloorMap) drawLine(line Line) {
	distance := line.distance()
	x := line.x1
	y := line.y1
	
	for i := 0; i <= distance; i++ {
		m.state[y][x]++

		if line.x1 < line.x2 {
			x++
		} else if line.x1 > line.x2 {
			x--
		}
		if line.y1 < line.y2 {
			y++
		} else if line.y1 > line.y2 {
			y--
		}
	}
}

func (m FloorMap) getDangerousPointCount() int {
	count := 0

	for _, row := range m.state {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}

	return count
}