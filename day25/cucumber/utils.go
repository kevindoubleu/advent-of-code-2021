package cucumber

import "strings"

func (c Cucumber) String() string {
	result := strings.Builder{}

	for _, row := range c.seamap {
		for _, cell := range row {
			result.WriteByte(cell)
		}
		result.WriteByte('\n')
	}

	return result.String()[: result.Len()-1 ]
}
