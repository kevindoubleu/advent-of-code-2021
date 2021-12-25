package main

import (
	"fmt"
	"strings"
)

// basically a graph
type CaveSystem struct {
	nodes	map[string]*Cave
	routes	[]*Route
}

func newCaveSystem(connections []string) CaveSystem {
	cs := CaveSystem{
		nodes: make(map[string]*Cave),
	}

	for _, connection := range connections {
		// parse the cave names
		caveNames := strings.Split(connection, "-")
		cave1Name := caveNames[0]
		cave2Name := caveNames[1]
		
		// if cave is already registered in the cavesystem
		// then just connect them together
		// else make then first
		cave1, ok := cs.nodes[cave1Name]
		if !ok {
			cave1 = newCave(cave1Name)
			cs.nodes[cave1Name] = cave1
		}
		cave2, ok := cs.nodes[cave2Name]
		if !ok {
			cave2 = newCave(cave2Name)
			cs.nodes[cave2Name] = cave2
		}

		cave1.connectCave(cave2)
	}

	return cs
}

func (cs CaveSystem) String() string {
	result := strings.Builder{}

	for _, node := range cs.nodes {
		result.WriteString(node.longString())
		result.WriteString("\n")
	}

	return result.String()
}

// regular DFS
func (cs CaveSystem) DFS() {
	visited := make(map[string]bool) // string is cave name
	cs._DFS(cs.nodes["start"], visited)
}

func (cs CaveSystem) _DFS(current *Cave, visited map[string]bool) {
	// mark it as visited (only small caves become visited)
	visited[current.name] = true
	fmt.Println("processing", current)

	// add all unvisited neighbours
	for _, neighbour := range current.neighbours {
		if !visited[neighbour.name] {
			cs._DFS(neighbour, visited)
		}
	}
}

// 1
func (cs *CaveSystem) allPaths() {
	// initialize routes
	cs.routes = make([]*Route, 0)

	visited := make(map[string]bool)
	route := newRoute()

	cs._allPaths(cs.nodes["start"], visited, route)
}

func (cs *CaveSystem) _allPaths(current *Cave, visited map[string]bool, currentPath Route) {
	// mark current as visited and part of current path
	if current.isSmall() {
		visited[current.name] = true
	}
	currentPath.push(current)

	// check if we're at destination
	if current.name == "end" {
		// fmt.Println(currentPath)
		cs.routes = append(cs.routes, &currentPath)
	}

	// visit all unvisited neighbours
	for _, neighbour := range current.neighbours {
		if !visited[neighbour.name] {
			cs._allPaths(neighbour, visited, currentPath)
		}
	}

	// no more unvisited neighbours
	// so we go back
	// fmt.Println("going back from", currentPath)
	currentPath.pop()
	visited[current.name] = false
}

// 2
func (cs *CaveSystem) allPaths2() {
	// initialize routes
	cs.routes = make([]*Route, 0)

	visited := make(map[string]bool)
	route := newRoute()

	cs._allPaths2(cs.nodes["start"], visited, route)

	cs.removeDupliateRoutes()
}

func (cs *CaveSystem) _allPaths2(current *Cave, visited map[string]bool, currentPath Route) {
	// mark current as visited and part of current path
	if current.isSmall() {
		visited[current.name] = true
	}
	currentPath.push(current)

	// check if we're at destination
	if current.name == "end" {
		cs.routes = append(cs.routes, &currentPath)
	}

	// visit all unvisited neighbours
	for _, neighbour := range current.neighbours {
		if !visited[neighbour.name] {
			// dont mark current (if small) for visiting 2x
			cs._allPaths2(neighbour, visited, currentPath)

			// mark current (if small) for visiting 2x
			if current.isSmall() && current.name != "start" && !currentPath.visitedSmall2x {
				currentPath.visitedSmall2x = true
				visited[current.name] = false
				cs._allPaths2(neighbour, visited, currentPath)
			}
		}
	}

	// no more unvisited neighbours
	// so we go back
	// fmt.Println("going back from", currentPath)
	currentPath.pop()
	visited[current.name] = false
}

func (cs *CaveSystem) removeDupliateRoutes() {
	seen := make(map[string]bool)
	newRoutes := []*Route{}

	for _, route := range cs.routes {
		if seen[route.String()] {
			continue
		}

		seen[route.String()] = true
		newRoutes = append(newRoutes, route)
	}

	cs.routes = newRoutes
}
