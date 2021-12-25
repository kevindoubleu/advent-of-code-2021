package cucumber

type Cucumber struct {
	seamap	[][]byte
}

func NewCucumber(seamap []string) Cucumber {
	height := len(seamap)

	c := Cucumber{
		seamap: make([][]byte, height),
	}

	for i, line := range seamap {
		c.seamap[i] = []byte(line)
	}

	return c
}

// attempt to move a '>'
func (c *Cucumber) moveRight(row, col int) (success bool) {
	// verify that we are going to move a '>'
	if c.seamap[row][col] != '>' { return false }

	// check if its going to wrap back around to 0
	next := col+1
	if col == len(c.seamap[0])-1 {
		next = 0
	}

	// check if its empty
	if c.seamap[row][next] != '.' { return false }

	// actually move the '>'
	c.seamap[row][next] = '>'
	// since all '>' move simultaneously, we mark our previous spot
	// so no one can go to it
	c.seamap[row][col] = 'x'
	return true
}

// attempt to move a 'v'
func (c *Cucumber) moveDown(row, col int) (success bool) {
	// verify that we are going to move a 'v'
	if c.seamap[row][col] != 'v' { return false }

	// check if its going to wrap back around to 0
	next := row+1
	if row == len(c.seamap)-1 {
		next = 0
	}

	// check if its empty
	if c.seamap[next][col] != '.' { return false }

	// actually move the 'v'
	c.seamap[next][col] = 'v'
	// since all 'v' move simultaneously, we mark our previous spot
	// so no one can go to it
	c.seamap[row][col] = 'x'
	return true
}

// cleans up the marks left after stepping
func (c *Cucumber) cleanUp() {
	height := len(c.seamap)
	width := len(c.seamap[0])

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if c.seamap[row][col] == 'x' {
				c.seamap[row][col] = '.'
			}
		}
	}
}

func (c *Cucumber) step() (moved bool) {
	height := len(c.seamap)
	width := len(c.seamap[0])

	moved = false

	// all facing right go first
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			ok := c.moveRight(row, col)
			if ok {
				col++
			}
			moved = moved || ok
		}
	}
	c.cleanUp()

	// then all facing down go
	for col := 0; col < width; col++ {
		for row := 0; row < height; row++ {
			ok := c.moveDown(row, col)
			if ok {
				row++
			}
			moved = moved || ok
		}
	}
	c.cleanUp()

	return moved
}

// returns how many steps are needed until
// all cucumbers are stuck, meaning none will move
func (c *Cucumber) StepsToStuck() int {
	steps := 1

	for c.step() {
		steps++

		// fmt.Println(c)
		// fmt.Fscanln(os.Stdin)
	}

	return steps
}
