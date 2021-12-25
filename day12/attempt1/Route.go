package main

import "strings"

type Route struct {
	nodes			[]*Cave
	visitedSmall2x	bool
}

func newRoute() Route {
	r := Route{
		nodes: make([]*Cave, 0),
		visitedSmall2x: false,
	}
	return r
}

func (r Route) String() string {
	result := strings.Builder{}

	for _, node := range r.nodes {
		result.WriteString(node.name)
		result.WriteString(",")
	}

	return result.String()[: result.Len()-1 ]
}

func (r *Route) push(node *Cave) {
	r.nodes = append(r.nodes, node)
}

func (r *Route) pop() {
	r.nodes = r.nodes[: len(r.nodes)-1 ]
}
