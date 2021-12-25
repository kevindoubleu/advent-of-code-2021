package main

type Bracket byte

func (b Bracket) isOpener() bool {
	return b == '(' ||
		b == '[' ||
		b == '{' ||
		b == '<'
}

func (b Bracket) isCloser() bool {
	return b == ')' ||
		b == ']' ||
		b == '}' ||
		b == '>'
}

// check if opener bracket b is pair of otherBracket
// also returns what otherBracket should be in order to be a matching pair
func (b Bracket) isPairOf(otherBracket Bracket) (bool, Bracket) {
	if !b.isOpener() 			{ return false, 0 }

	switch b {
	case '(':
		return otherBracket == ')', ')'
	case '[':
		return otherBracket == ']', ']'
	case '{':
		return otherBracket == '}', '}'
	case '<':
		return otherBracket == '>', '>'
	}

	return false, 0
}


type BracketLine []Bracket

type BracketLineState int
const (
	complete   = iota
	incomplete // not enough closing brackets
	corrupted  // wrong closing brackets
	invalid    // not brackets
)
// check if BracketLine is complete or incomplete or corrupted
// if corrupted, also returns ptr to the corrupted closing bracket
// if incomplete, also returns ptr to stack of opener brackets
func (l BracketLine) check() (BracketLineState, BracketLine) {
	stack := []Bracket{}

	for _, bracket := range l {
		if bracket.isOpener() {
			stack = append(stack, bracket)
		} else if bracket.isCloser() {
			popped := stack[len(stack)-1]

			isPair, _ := popped.isPairOf(bracket)
			if !isPair {
				return corrupted, BracketLine{bracket}
			}

			stack = stack[: len(stack)-1]
		} else {
			return invalid, nil
		}
	}

	if len(stack) != 0 {
		return incomplete, stack
	}
	return complete, nil
}

func (l BracketLine) corruptedScore() int {
	if state, brackets := l.check(); state == corrupted {
		switch brackets[0] {
		case ')':
			return 3
		case ']':
			return 57
		case '}':
			return 1197
		case '>':
			return 25137
		}
	}

	return 0
}

func (l BracketLine) reverse() BracketLine {
	reversed := []Bracket{}

	// invert the bracket pairs
	for _, bracket := range l {
		_, closer := bracket.isPairOf(0)
		reversed = append(reversed, closer)
	}

	// reverse the order
	front, back := 0, len(reversed)-1
	for front < back {
		reversed[front], reversed[back] = reversed[back], reversed[front]
		front++
		back--
	}

	return reversed
}

func (l BracketLine) incompleteScore() int {
	if state, openerBrackets := l.check(); state == incomplete {
		closerBrackets := openerBrackets.reverse()

		score := 0
		for _, bracket := range closerBrackets {
			score *= 5
			switch bracket {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}

		return score
	}

	return 0
}
