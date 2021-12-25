package sfnum

import (
	"fmt"
	"regexp"
	"strings"
)

type SFNumber struct {
	notation	string
}

func NewSFNumber(str string) SFNumber {
	n := SFNumber{
		notation: str,
	}
	return n
}

func (n *SFNumber) explode() {
	openerBrackets := 0
	left, right := 0, 0

	strLen := len(n.notation)
	for i := 0; i < strLen; i++ {
		if n.notation[i] == '[' {
			openerBrackets++
		}

		// Exploding pairs will always consist of two regular numbers.
		// this means we can immediately parse the numbers
		// bcs there will be no more nests (opener brackets)
		if openerBrackets == 4 {
			// get the numbers
			strReader := strings.NewReader(n.notation[i+1 :])
			fmt.Fscanf(strReader, "%d,%d", &left, &right)
			
			// replace the current [x,y] with 0
			pairMatcher := regexp.MustCompile(`\[\d,\d\]`)
			pair := pairMatcher.Find([]byte(n.notation[i+1 :]))
			n.notation = strings.Replace(n.notation, string(pair), "0", 1)

			// give the value of left & right to others
			

			// restart from the beginning of str
			i = 0
			openerBrackets = 0
		}
	}
}
