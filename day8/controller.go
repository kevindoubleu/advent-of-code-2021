package main

type Controller struct {
	signalToSegment	map[byte]int
	display			map[int]bool
}

func newController() Controller {
	return Controller{
		signalToSegment: make(map[byte]int),
		display: make(map[int]bool),
	}
}

func (c *Controller) assignSignalToSegment(signal byte, segment int) {
	c.signalToSegment[signal] = segment
}

// reads a signal (e.g. "abcd") and displays it based on these segment labels
//  111
// 2   3
//  444
// 5   6
//  777
func (c *Controller) displaySignal(signal string) {
	for _, char := range signal {
		c.display[c.signalToSegment[byte(char)]] = true
	}
}

func (c *Controller) resetDisplay() {
	c.display = make(map[int]bool)
}

func (c Controller) readDisplay() int {
	// go from most segments to least segments

	if c.display[1] &&
		c.display[2] &&
		c.display[3] &&
		c.display[4] &&
		c.display[5] &&
		c.display[6] &&
		c.display[7] {
		return 8
	}

	if c.display[1] &&
		c.display[2] &&
		c.display[3] &&
		c.display[4] &&
		c.display[6] &&
		c.display[7] {
		return 9
	}

	if c.display[1] &&
		c.display[2] &&
		c.display[4] &&
		c.display[5] &&
		c.display[6] &&
		c.display[7] {
		return 6
	}

	if c.display[1] &&
		c.display[2] &&
		c.display[3] &&
		c.display[5] &&
		c.display[6] &&
		c.display[7] {
		return 0
	}

	if c.display[1] &&
		c.display[3] &&
		c.display[4] &&
		c.display[5] &&
		c.display[7] {
		return 2
	}

	if c.display[1] &&
		c.display[3] &&
		c.display[4] &&
		c.display[6] &&
		c.display[7] {
		return 3
	}

	if c.display[1] &&
		c.display[2] &&
		c.display[4] &&
		c.display[6] &&
		c.display[7] {
		return 5
	}

	if c.display[2] &&
		c.display[3] &&
		c.display[4] &&
		c.display[6] {
		return 4
	}

	if c.display[1] &&
		c.display[3] &&
		c.display[6] {
		return 7
	}

	if c.display[3] &&
		c.display[6] {
		return 1
	}

	return -1
}
