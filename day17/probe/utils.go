package probe

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	x, y int
}

// check if a coord is within an area
func (c Coordinate) isWithin(area Area) bool {
	return (
		area.XLow <= c.x && c.x <= area.XHigh &&
		area.YLow <= c.y && c.y <= area.YHigh )
}

// check if a coord is below an area
func (c Coordinate) isBelow(area Area) bool {
	return c.y < area.YLow
}

// check if a coord has passed an area
// assuming it came from (0,0)
func (c Coordinate) isAfter(area Area) bool {
	return c.x > area.XHigh
}




type Area struct {
	XLow, XHigh int
	YLow, YHigh int
}

func NewArea(str string) Area {
	a := Area{}

	strReader := strings.NewReader(str)
	fmt.Fscanf(
		strReader,
		"target area: x=%d..%d, y=%d..%d",
		&a.XLow, &a.XHigh, &a.YLow, &a.YHigh)
	
	return a
}




// get the furthest distance a velocity will reach
func GetMaxVelocity(initial int) int {
	total := 0

	// max horizontal distance is: x + (x-1) + (x-2) + ... + 1
	// due to drag reducing x velocity by 1 each step

	// max height is: y + (y-1) + (y-2) + ... + 1
	// due to gravity reducing y velocity by 1 each step
	for i := initial; i > 0; i-- {
		total += i
	}

	return total
}
