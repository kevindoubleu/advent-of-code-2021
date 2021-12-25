package cave

import (
	"fmt"
	"strings"
)

type CaveSystem struct {
	// map of cave names to their address
	caves	map[string]*Cave
}

func NewCaveSystem(connections []string) CaveSystem {
	cs := CaveSystem{
		caves: make(map[string]*Cave),
	}

	for _, connection := range connections {
		name1, name2 := parseCaveNames(connection)

		// if cave hasnt existed then make it
		if cs.caves[name1] == nil {
			cs.caves[name1] = newCave(name1)
		}
		if cs.caves[name2] == nil {
			cs.caves[name2] = newCave(name2)
		}

		// then connect them up
		connect2Way(cs.caves[name1], cs.caves[name2])
	}

	return cs
}

func (cs CaveSystem) String() string {
	result := strings.Builder{}

	for name, cave := range cs.caves {
		result.WriteString(fmt.Sprint(name, ": "))

		for _, neighbour := range cave.neighbours {
			result.WriteString(fmt.Sprint(neighbour.name, ", "))
		}

		result.WriteString("\n")
	}

	return result.String()[: result.Len()-1 ]
}

func (cs CaveSystem) AllPaths() {
	
}
