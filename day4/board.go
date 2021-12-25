package main

import (
	"fmt"
)

// to make it easy to look if a number is marked or not
type Number struct {
	value int
	marked bool
}

// to have board methods
type Board [5][5]Number

func (b *Board) print() {
	fmt.Println("====================")
	for _, row := range b {
		for _, cell := range row {
			if cell.marked {
				fmt.Printf("[%2d]", cell.value)
			} else {
				fmt.Printf(" %2d ", cell.value)
			}
		}
		fmt.Println("")
	}
	fmt.Println("====================")
}

// make a board invalid, quickfix for part 2
func (b *Board) nullify() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b[i][j] = Number{
				value: -1,
				marked: false,
			}
		}
	}
}

// marks number on board
// returns non-zero of board got a bingo
func (b *Board) markNumber(number int) int {
	for i := range b {
		for j := range b {
			if number == b[i][j].value {
				b[i][j].marked = true
				
				// check vertical and horizontal bingo
				marks := 0
				for k := 0; k < 5; k++ {
					if b[i][k].marked {
						marks++
					}
					if marks == 5 {
						return b.bingo(number)
					}
				}
				marks = 0
				for k := 0; k < 5; k++ {
					if b[k][j].marked {
						marks++
					}
					if marks == 5 {
						return b.bingo(number)
					}
				}
			}
		}
	}

	return 0
}

// count score after a bingo
// lastNumber is the number that was just called when the board won
func (b Board) getScore(lastNumber int) int {
	sum := 0

	for _, row := range b {
		for _, cell := range row {
			if !cell.marked {
				sum += cell.value
			}
		}
	}

	return sum * lastNumber
}

func (b *Board) bingo(lastNumber int) int {
	score := b.getScore(lastNumber)

	b.print()
	fmt.Println("bingo score:", score)

	b.nullify()
	return score
}

// to easily ask all boards to run something
type BingoGame struct {
	players		[]Board
	playerCount	int
	score		int
}

func newBingoGame(boards []Board) *BingoGame {
	game := BingoGame{
		players      : boards,
		playerCount  : len(boards),
		score        : 0,
	}

	return &game
}

func (g *BingoGame) announce(number int) {
	for i := 0; i < g.playerCount; i++ {
		// the rule is that if feedback is 0 then its an "ok i heard that"
		// if not, then its a bingo, and the feedback is the score
		feedback := g.players[i].markNumber(number)
		if feedback != 0 {
			g.score = feedback
		}
	}
}
