package main

import (
	"fmt"
	"strings"
)

// a node in a graph
type Cave struct {
	name		string
	neighbours	[]*Cave
}

func newCave(name string) *Cave {
	c := Cave{
		name: name,
		neighbours: make([]*Cave, 0),
	}
	return &c
}

func (c Cave) String() string {
	return c.name
}

func (c Cave) longString() string {
	result := strings.Builder{}

	result.WriteString(fmt.Sprint(c.name, ": "))
	for _, neighbour := range c.neighbours {
		result.WriteString(fmt.Sprint(neighbour.name, ", "))
	}

	return result.String()[ : result.Len()-2 ]
}

func (c *Cave) connectCave(otherCave *Cave) {
	// undirected so 2-way connections
	c.neighbours = append(c.neighbours, otherCave)
	// start must not have any edges going in to it
	// end must not have any edges going out of it
	if otherCave.name != "end" && c.name != "start" {
		otherCave.neighbours = append(otherCave.neighbours, c)
	}
}

// check if cave is small
func (c Cave) isSmall() bool {
	return c.name == strings.ToLower(c.name)
}
