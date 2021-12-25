package dirac

import (
	"fmt"
	"strings"
)

func ParsePlayerPositions(positions []string) (pos1, pos2 int) {
	fmt.Fscanf(
		strings.NewReader(positions[0]),
		"Player 1 starting position: %d",
		&pos1)
	fmt.Fscanf(
		strings.NewReader(positions[1]),
		"Player 2 starting position: %d",
		&pos2)
		
	return pos1, pos2
}
