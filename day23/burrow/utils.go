package burrow

import (
	"fmt"
)

func difference(a, b int) int {
	if a > b {
		return a-b
	}
	return b-a
}

func (b Burrow) String() string {
	corridor := []byte{}
	for _, space := range b.spaces {
		if space == 0 {
			corridor = append(corridor, '.')
		} else {
			corridor = append(corridor, space)
		}
	}

	tenants := []byte{}
	for i := 3; i >= 0; i-- {
		for j := 0; j < 4; j++ {
			if b.rooms[j].tenants[i] == 0 {
				tenants = append(tenants, '.')	
			} else {
				tenants = append(tenants, b.rooms[j].tenants[i])
			}
		}
	}

	return fmt.Sprintf(
		"#############\n" +
		"#%c%c.%c.%c.%c.%c%c#\n" +
		"###%c#%c#%c#%c###\n" +
		"  #%c#%c#%c#%c#\n" +
		"  #%c#%c#%c#%c#\n" +
		"  #%c#%c#%c#%c#\n" +
		"  #########\n",
		corridor[0], corridor[1], corridor[2], corridor[3], corridor[4], corridor[5], corridor[6],
		tenants[0], tenants[1], tenants[2], tenants[3],
		tenants[4], tenants[5], tenants[6], tenants[7],
		tenants[8], tenants[9], tenants[10], tenants[11],
		tenants[12], tenants[13], tenants[14], tenants[15])
}

// translate room index to actual room index in diagram string
func toRoomIdx(roomNumber int) int {
	switch roomNumber {
	case 0:
		return 3
	case 1:
		return 5
	case 2:
		return 7
	case 3:
		return 9
	}
	panic("invalid room number")
}

// translate space index to actual space index in diagram string
func toSpaceIdx(spaceNumber int) int {
	switch spaceNumber {
	case 0:
		return 1
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 6
	case 4:
		return 8
	case 5:
		return 10
	case 6:
		return 11
	}
	panic("invalid space number")
}

// counts distance (in steps) between last tenant in room to space
func (b Burrow) distance(roomNumber, spaceNumber int) int {
	// translate space and room to actual index in diagram
	// and return their difference +
	// how many empty spots in the room +
	// 1 (for getting out of the room into the corridor)
	
	roomIdx := toRoomIdx(roomNumber)
	spaceIdx := toSpaceIdx(spaceNumber)

	return difference(roomIdx, spaceIdx) +
		(b.rooms[roomNumber].cap - len(b.rooms[roomNumber].tenants)) +
		1
}
