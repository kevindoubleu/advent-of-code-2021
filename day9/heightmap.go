package main

type Coordinate struct {
	row, col	int
}

type Heightmap struct {
	heights	[][]int
}

func newHeightmap(floorData [][]int) Heightmap {
	return Heightmap{
		heights: floorData,
	}
}

// safely determine if coord (col,row) is a low point
// by looking at the height of its 4 adjacent cells, if it exists
func (m Heightmap) isLowPoint(point Coordinate) bool {
	row, col  := point.row, point.col
	here      := m.heights[row][col]
	lowerThan := 0

	if row <= 0                   || m.heights[row-1][col] > here { lowerThan++ }
	if row >= len(m.heights)-1    || m.heights[row+1][col] > here { lowerThan++ }
	if col <= 0                   || m.heights[row][col-1] > here { lowerThan++ }
	if col >= len(m.heights[0])-1 || m.heights[row][col+1] > here { lowerThan++ }

	return lowerThan == 4
}

func (m Heightmap) getLowPoints() []Coordinate {
	lowPoints := []Coordinate{}

	for i, row := range m.heights {
		for j := range row {
			point := Coordinate{
				row: i,
				col: j,
			}

			if m.isLowPoint(point) {
				lowPoints = append(lowPoints, point)
			}
		}
	}

	return lowPoints
}

// 1
func (m Heightmap) getRisk() int {
	lowPoints := m.getLowPoints()
	risk := 0

	for _, lowpoint := range lowPoints {
		risk += m.heights[lowpoint.row][lowpoint.col] + 1
	}

	return risk
}

// 2
func (m Heightmap) getBasinSize(start Coordinate) int {
	visited := make(map[Coordinate]bool)
	
	m.expandBasin(start, &visited)

	return len(visited)
}

func (m Heightmap) expandBasin(curr Coordinate, visited *map[Coordinate]bool) {
	if (*visited)[curr] {
		return
	}
	(*visited)[curr] = true

	// go up
	if curr.row > 0 && m.heights[curr.row-1][curr.col] != 9 {
		m.expandBasin(Coordinate{
			row: curr.row-1,
			col: curr.col,
		}, visited)
	}

	// go down
	rowCount := len(m.heights)
	if curr.row < rowCount-1 && m.heights[curr.row+1][curr.col] != 9 {
		m.expandBasin(Coordinate{
			row: curr.row+1,
			col: curr.col,
		}, visited)
	}

	// go left
	if curr.col > 0 && m.heights[curr.row][curr.col-1] != 9 {
		m.expandBasin(Coordinate{
			row: curr.row,
			col: curr.col-1,
		}, visited)
	}

	// go right
	colCount := len(m.heights[0])
	if curr.col < colCount-1 && m.heights[curr.row][curr.col+1] != 9 {
		m.expandBasin(Coordinate{
			row: curr.row,
			col: curr.col+1,
		}, visited)
	}
}
