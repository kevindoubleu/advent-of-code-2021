package coregrp

type Core struct {
	// store the cuboids that are on
	cuboids []Cuboid
	// we then can get the area of them all for the answer
}

func NewCore() Core {
	c := Core{
		cuboids: make([]Cuboid, 0),
	}
	return c
}

func (c *Core) Set(cuboid Cuboid) {
	c.cuboids = append(c.cuboids, cuboid)
}

func (c *Core) Process(instructions []string) {
	for _, instruction := range instructions {
		cuboid := newCuboid(instruction)
		c.Set(cuboid)
	}
}
