package cave

import "strings"

type Cave struct {
	name		string
	neighbours	map[string]*Cave
}

func newCave(name string) *Cave {
	c := Cave{
		name: name,
		neighbours: make(map[string]*Cave),
	}
	return &c
}

func connect1Way(a, b *Cave) {
	if a == nil || b == nil { return }

	// start may not have nodes pointing to it
	// end may not have nodes pointing out of it
	if a.name == "end" || b.name == "start" { return }

	// dont add duplicates
	if a.neighbours[b.name] != nil { return }

	a.neighbours[b.name] = b
}

func connect2Way(a, b *Cave) {
	connect1Way(a, b)
	connect1Way(b, a)
}

func (c Cave) isBig() bool {
	return len(c.name) == 1 && c.name == strings.ToUpper(c.name)
}

func (c Cave) isSmall() bool {
	return len(c.name) == 1 && c.name == strings.ToLower(c.name)
}
