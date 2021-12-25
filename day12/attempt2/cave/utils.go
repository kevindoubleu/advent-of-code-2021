package cave

import (
	"strings"
)

func parseCaveNames(connection string) (name1, name2 string) {
	names := strings.Split(connection, "-")
	name1 = names[0]
	name2 = names[1]
	
	return name1, name2
}
