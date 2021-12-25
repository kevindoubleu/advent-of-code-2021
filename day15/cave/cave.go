package cave

import (
	"advent-of-code-2021/lib"
	"fmt"
)

type Coordinate struct {
	x, y int
}

type Cave struct {
	// for the weights table in dijkstra
	weights	[][]int
}

func NewCave(risks [][]int) Cave {
	c := Cave{
		weights: risks,
	}
	return c
}

// prints slice of shortest path to all nodes using dijkstra
func (c Cave) ShortestPath() {
	tableSize := len(c.weights)
	distances := tableOfInfinites(tableSize)
	visiteds := tableOfBool(tableSize)

	// we start at (0,0), it has 0 distance to itself
	distances[0][0] = 0

	// for all vertex
	vertexCount := tableSize * tableSize
	for i := 0; i < vertexCount; i++ {
		// find minimum of unvisited nodes, the currently selected node
		currX, currY := minimumUnvisited(distances, visiteds)

		// visit it
		visiteds[currY][currX] = true

		// relax all neighbours of curr node
		neighbours := c.neighboursOf(Coordinate{currX, currY})
		for _, n := range neighbours {
			
			// only unvisited neighbours
			if !visiteds[n.y][n.x] &&
				// and only if new distance is smaller than old distance
				distances[currY][currX] + c.weights[n.y][n.x] < distances[n.y][n.x] {
				
				// update with new distance
				distances[n.y][n.x] = distances[currY][currX] + c.weights[n.y][n.x]
			}
		}
	}

	fmt.Println(lib.Print2DIntSlice(distances))
}

// 2
func (c *Cave) Enlarge() {
	oriSize := len(c.weights)
	multiplier := 5
	newSize := oriSize * multiplier

	newWeights := make([][]int, newSize)
	for i := range newWeights {
		newWeights[i] = make([]int, newSize)
	}

	for i := 0; i < newSize; i++ {
		for j := 0; j < newSize; j++ {
			//                     repeating the original values     inc to the right  inc downward    loop back to 1 if 10
			newWeights[i][j] = ((c.weights[i % oriSize][j % oriSize] + (j / oriSize) + (i / oriSize)) - 1) % 9 + 1
		}
	}

	c.weights = newWeights
}
