package lib

import (
	"strings"
)

// splits on double newlines
// parses the parts into slice of strings
// saves the parts in the str pointers given
func UnpackToStrSlices(fileContent string, destinations ...*[]string) {
	parts := strings.Split(fileContent, "\n\n")

	for i, dest := range destinations {
		// initialize dest if not yet
		if dest == nil {
			initialized := make([]string, 0)
			*destinations[i] = initialized
		}

		lines := ToStrSlice(parts[i])
		*destinations[i] = lines
	}
}