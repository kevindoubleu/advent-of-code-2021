package burrow

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strings"
)

type Burrow struct {
	rooms	[]Room
	spaces	[]byte
}

func NewBurrow(diagram string) Burrow {
	b := Burrow{
		rooms  : make([]Room, 4),
		spaces : make([]byte, 7),
	}

	lines := lib.ToStrSlice(diagram)

	// make rooms with appropriate size
	for i := range b.rooms {
		b.rooms[i] = newRoom(4)
	}
	
	// read amphipods in each room
	// we go from bottom to top so we can use .push method
	roomNum := 0
	for depth := len(lines)-1; depth >= 2; depth-- {
		for _, char := range lines[depth] {
			if strings.ContainsAny(string(char), "ABCD") {
				b.rooms[roomNum].push(byte(char))
				roomNum++
			}
		}
		roomNum = 0
	}

	return b
}

func cloneBurrow(original Burrow) Burrow {
	b := Burrow{
		rooms  : make([]Room, 4),
		spaces : make([]byte, 7),
	}

	for i := range original.rooms {
		b.rooms[i] = cloneRoom(original.rooms[i])
	}
	for i := range original.spaces {
		b.spaces[i] = original.spaces[i]
	}

	return b
}

var costs = map[byte]int{
	'A': 1,
	'B': 10,
	'C': 100,
	'D': 1000,
}

// each amphipod type's destination room numbers
var destinations = map[byte]int{
	'A': 0,
	'B': 1,
	'C': 2,
	'D': 3,
}

// fill space from room to space
// from is the room number
// to is the space number
// if cant get into space then return 0
func (b *Burrow) fillSpace(from int, to int) (cost int) {
	if b.spaces[to] != 0 { return 0 }
	if b.rooms[from].top() == 0 { return 0 }

	// check if amphipod is already at final destination
	if from == destinations[b.rooms[from].top()] {
		return 0
	}

	// check if path is clear
	// we do this just by checking the corridor
	roomIdx := toRoomIdx(from)
	spaceIdx := toSpaceIdx(to)
	for i := roomIdx; i < spaceIdx; i++ {
		if lib.ToStrSlice(b.String())[1][i] != '.' {
			return 0
		}
	}

	// calculate the cost
	cost = costs[b.rooms[from].top()] * b.distance(from, to)

	// actually move them
	b.spaces[to] = b.rooms[from].pop()

	return cost
}

// try to move the amphipod in space to its destination room
// from is the space number
// if cant move, then returns 0
func (b *Burrow) fillRoom(from int) (cost int) {
	if b.spaces[from] == 0 { return 0 }

	switch b.spaces[from] {
	case 'A':
		// room is not empty and filled with other amphipods
		if b.rooms[0].tenantCount() > 0 && !b.rooms[0].homogenous() {
			return 0
		}

		// path to room 1 must be clear
		// if space is on the left of room 1
		for i := from+1; i <= 1; i++ {
			if b.spaces[i] != 0 {
				return 0
			}
		}
		// if space is on the right of room 1
		for i := from-1; i >= 2; i-- {
			if b.spaces[i] != 0 {
				return 0
			}
		}

		// actually move the amphipod
		b.rooms[0].push(b.spaces[from])
		b.spaces[from] = 0

		// path clear return cost
		return (b.distance(0, from)+1) * costs[b.rooms[0].top()]
		
	case 'B':
		// room is not empty and filled with other amphipods
		if b.rooms[1].tenantCount() > 0 && !b.rooms[1].homogenous() {
			return 0
		}

		// path to room 2 must be clear
		// if space is on the left of room 2
		for i := from+1; i <= 2; i++ {
			if b.spaces[i] != 0 {
				return 0
			}
		}
		// if space is on the right of room 2
		for i := from-1; i >= 3; i-- {
			if b.spaces[i] != 0 {
				return 0
			}
		}

		// actually move the amphipod
		b.rooms[1].push(b.spaces[from])
		b.spaces[from] = 0

		// path clear return cost
		return (b.distance(1, from)+1) * costs[b.rooms[1].top()]
	case 'C':
		// room is not empty and filled with other amphipods
		if b.rooms[2].tenantCount() > 0 && !b.rooms[2].homogenous() {
			return 0
		}

		// path to room 3 must be clear
		// if space is on the left of room 3
		for i := from+1; i <= 3; i++ {
			if b.spaces[i] != 0 {
				return 0
			}
		}
		// if space is on the right of room 3
		for i := from-1; i >= 4; i-- {
			if b.spaces[i] != 0 {
				return 0
			}
		}

		// actually move the amphipod
		b.rooms[2].push(b.spaces[from])
		b.spaces[from] = 0

		// path clear return cost
		return (b.distance(2, from)+1) * costs[b.rooms[2].top()]

	case 'D':
		// room is not empty and filled with other amphipods
		if b.rooms[3].tenantCount() > 0 && !b.rooms[3].homogenous() {
			return 0
		}

		// path to room 4 must be clear
		// if space is on the left of room 2
		for i := from+1; i <= 4; i++ {
			if b.spaces[i] != 0 {
				return 0
			}
		}
		// if space is on the right of room 
		for i := from-1; i >= 5; i-- {
			if b.spaces[i] != 0 {
				return 0
			}
		}

		// actually move the amphipod
		b.rooms[3].push(b.spaces[from])
		b.spaces[from] = 0

		// path clear return cost
		return (b.distance(3, from)+1) * costs[b.rooms[3].top()]
	}
	return 0
}

// branching recursive function to try all possibilities
// each recursive call represents a state of the diagram (burrow)
// so it has to create copies, which is why we use value receiver
func (b Burrow) AllScenarios(currentCost int) {

	// for each space try going into room
	for spaceNumber := range b.spaces {

		newBranch := cloneBurrow(b)
		goToRoomCost := newBranch.fillRoom(spaceNumber)
		if goToRoomCost != 0 {
			// since each amphipod thats outside a room will only go to their destination room
			// we just continue this branch with the space freed
			b.fillRoom(spaceNumber)
			currentCost += goToRoomCost
		}
	}

	// for each room try to go to each empty space
	for roomNumber := range b.rooms {
		for spaceNumber := range b.spaces {

			newBranch := cloneBurrow(b)
			goToSpaceCost := newBranch.fillSpace(roomNumber, spaceNumber)
			// fmt.Println("cloned burrow for fillspace", newBranch)
			if goToSpaceCost != 0 {
				// fmt.Println("branched fillspace at room", roomNumber, "space", spaceNumber)
				newBranch.AllScenarios(currentCost + goToSpaceCost)
			}
		}
	}

	completed := 0
	for _, room := range b.rooms {
		// if room.homogenous() && room.tenantCount() == 4 {
		if room.homogenous() {
			completed++
		}
	}

	// if the rooms are not complete and we've done everything then this is a dead end
	// if all rooms are filled then this is a possible solution
	if completed == 4 {
		fmt.Println("successfult branch with cost", currentCost)
		fmt.Println(b)
	}
}
