package coregrp

import (
	"fmt"
	"strings"
)

type Cuboid struct {
	state		bool
	XLow, XHigh	int
	YLow, YHigh	int
	ZLow, ZHigh	int
}

func newCuboid(instruction string) Cuboid {
	c := Cuboid{}

	var state string
	fmt.Fscanf(
		strings.NewReader(instruction),
		"%s x=%d..%d,y=%d..%d,z=%d..%d",
		&state, &c.XLow, &c.XHigh, &c.YLow, &c.YHigh, &c.ZLow, &c.ZHigh)
	
	if state == "on" {
		c.state = true
	} else if state == "off" {
		c.state = false
	}

	return c
}
