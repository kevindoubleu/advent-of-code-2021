package main

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	row, col int
}

type Octopi struct {
	energyLvls	[][]int
	flashed		[][]bool
}

func newOctopi(energyLvls [][]int) Octopi {
	o := Octopi{
		energyLvls: energyLvls,
		flashed: make([][]bool, 0),
	}

	colCount := len(energyLvls[0])
	for range energyLvls {
		o.flashed = append(o.flashed, make([]bool, colCount))
	}

	return o
}

func (o Octopi) String() string {
	result := strings.Builder{}

	for _, row := range o.energyLvls {
		for _, cell := range row {
			result.WriteString(fmt.Sprintf("%2d", cell))
		}
		result.WriteString("\n")
	}

	return result.String()
}

func (o *Octopi) steps(n int) int {
	total := 0

	for i := 0; i < n; i++ {
		total += o.step()
		fmt.Println(o)
		fmt.Println("step", i+1, "done")
	}

	return total
}

// returns number of flashes happened at this step
func (o *Octopi) step() int {
	// add 1 to all octopi
	for i, row := range o.energyLvls {
		for j := range row {
			o.energyLvls[i][j]++
		}
	}

	flashes := o.executeFlashProcedure()

	o.cleanUpFlashResidue()

	return flashes
}

func (o *Octopi) executeFlashProcedure() int {
	// keep flashing until there are no new flashes
	totalFlashes := 0
	for {
		flashes := o.flashAllReady()
		totalFlashes += flashes
		if flashes == 0 {
			break
		}
	}

	return totalFlashes
}

func (o *Octopi) flash(from Coordinate) {
	row, col := from.row, from.col

	// straights
	if row > 0 { o.energyLvls[row-1][col]++ }
	if row < 9 { o.energyLvls[row+1][col]++ }
	if col > 0 { o.energyLvls[row][col-1]++ }
	if col < 9 { o.energyLvls[row][col+1]++ }
	// diagonals
	if row > 0 && col > 0 { o.energyLvls[row-1][col-1]++ }
	if row > 0 && col < 9 { o.energyLvls[row-1][col+1]++ }
	if row < 9 && col > 0 { o.energyLvls[row+1][col-1]++ }
	if row < 9 && col < 9 { o.energyLvls[row+1][col+1]++ }

	o.energyLvls[row][col] = 0
	o.flashed[row][col] = true
}

func (o *Octopi) flashAllReady() int {
	flashes := 0

	for i, row := range o.energyLvls {
		for j, cell := range row {
			if cell > 9 {
				o.flash(Coordinate{
					row: i,
					col: j,
				})
				flashes++
			}
		}
	}

	return flashes
}

func (o *Octopi) cleanUpFlashResidue() {
	for i, row := range o.flashed {
		for j, cellFlashed := range row {
			if cellFlashed {
				o.energyLvls[i][j] = 0
				o.flashed[i][j] = false
			}
		}
	}
}

// 2
func (o *Octopi) stepsToSync() int {
	steps := 1
	for {
		if o.step() == 100 {
			return steps
		}
		steps++
	}
}
