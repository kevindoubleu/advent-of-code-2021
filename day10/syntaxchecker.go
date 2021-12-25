package main

import (
	"sort"
)

type SyntaxChecker struct {
	linesOfCode []BracketLine
}

func newSyntaxChecker(code []string) SyntaxChecker {
	s := SyntaxChecker{
		linesOfCode: make([]BracketLine, len(code)),
	}

	for i, loc := range code {
		s.linesOfCode[i] = BracketLine(loc)
	}

	return s
}

func (s SyntaxChecker) totalCorruptedScore() int {
	total := 0

	for _, line := range s.linesOfCode {
		total += line.corruptedScore()
	}

	return total
}

func (s SyntaxChecker) totalIncompleteScore() int {
	scores := []int{}

	for _, line := range s.linesOfCode {
		// only take into account the lines that are incomplete
		if state, _ := line.check(); state == incomplete {
			scores = append(scores, line.incompleteScore())
		}
	}

	// total score is median of all scores
	sort.Ints(scores)

	return scores[ len(scores) / 2 ]
}
