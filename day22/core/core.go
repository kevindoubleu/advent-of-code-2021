package core

type Core struct {
	// lets try simulating individual cubes for pt 1
	cubes map[int]map[int]map[int]bool
}

func NewCore() Core {
	c := Core{
		cubes: make(map[int]map[int]map[int]bool),
	}
	return c
}

func (c *Core) set(cuboid Cuboid) {
	for i := cuboid.XLow; i <= cuboid.XHigh; i++ {
		for j := cuboid.YLow; j <= cuboid.YHigh; j++ {
			for k := cuboid.ZLow; k <= cuboid.ZHigh; k++ {
				// make the nested maps if not exist yet
				if c.cubes[i] == nil {
					c.cubes[i] = make(map[int]map[int]bool)
				}
				if c.cubes[i][j] == nil {
					c.cubes[i][j] = make(map[int]bool)
				}

				// after making sure we are not assigning to nil map
				c.cubes[i][j][k] = cuboid.state
			}
		}
	}
}

func (c *Core) Initialize(steps []string) {
	for _, step := range steps {
		cuboid := newCuboid(step)
		// only process xyz -50..50
		if cuboid.XLow < -50 || cuboid.XHigh > 50 ||
			cuboid.YLow < -50 || cuboid.YHigh > 50 ||
			cuboid.ZLow < -50 || cuboid.ZHigh > 50 {
			continue
		}
		c.set(cuboid)
	}
}

func (c Core) CountOn() int {
	total := 0

	for x, length := range c.cubes {
		for y, width := range length {
			for z := range width {
				if c.cubes[x][y][z] {
					total++
				}
			}
		}
	}

	return total
}
