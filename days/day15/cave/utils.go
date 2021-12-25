package cave

import (
	"advent-of-code-2021/lib"
	"math"
)

// all cells are false
func tableOfBool(size int) [][]bool {
	result := make([][]bool, size)

	for i := range result {
		result[i] = make([]bool, size)
	}

	return result
}

// returned table will be a square
func tableOfInfinites(size int) [][]int {
	result := make([][]int, size)

	for i := range result {
		row := make([]int, size)
		
		for j := range row {
			// input is 100x100, so max is 9*10000 = 90000
			row[j] = math.MaxInt32
		}

		result[i] = row
	}

	return result
}

// this is first candidate for time complexity improvement
// this function iterates through whole table
// and it is called for each cell in the table
func minimumUnvisited(distances [][]int, visited [][]bool) (minX, minY int) {
	min := math.MaxInt32

	for i, row := range visited {
		for j := range row {
			if !visited[i][j] && distances[i][j] < min {
				min = distances[i][j]
				minX = j
				minY = i
			}
		}
	}

	return minX, minY
}

// safely returns neighbours of coordinate in cave
func (c Cave) neighboursOf(coord Coordinate) []Coordinate {
	caveSize := len(c.weights)
	neighbours := []Coordinate{}

	// up
	if coord.y > 0 {
		neighbours = append(neighbours, Coordinate{
			x: coord.x,
			y: coord.y-1,
		})
	}
	// down
	if coord.y < caveSize-1 {
		neighbours = append(neighbours, Coordinate{
			x: coord.x,
			y: coord.y+1,
		})
	}
	// left
	if coord.x > 0 {
		neighbours = append(neighbours, Coordinate{
			x: coord.x-1,
			y: coord.y,
		})
	}
	// right
	if coord.x < caveSize-1 {
		neighbours = append(neighbours, Coordinate{
			x: coord.x+1,
			y: coord.y,
		})
	}

	return neighbours
}

func (c Cave) String() string {
	return lib.Print2DIntSlice(c.weights)
}
