package main

import (
	"advent-of-code-2021/lib"
	"fmt"
)

var TESTING = false

func main() {
	var input string
	if TESTING {
		input = lib.ReadTest()
	} else {
		input = lib.ReadInput()
	}

	fmt.Println(getWinnerScore(getBoards(input), getCalls(input)))
	fmt.Println(getLastWinnerScore(getBoards(input), getCalls(input)))
}

// 1
func getWinnerScore(boards []Board, calls []int) int {
	game := newBingoGame(boards)

	for _, number := range calls {
		game.announce(number)
		if game.score > 0 {
			return game.score
		}
	}
	
	return 0
}

// 2
func getLastWinnerScore(boards []Board, calls []int) int {
	game := newBingoGame(boards)
	lastGameScore := 0

	for _, number := range calls {
		game.announce(number)
		if game.score > 0 {
			lastGameScore = game.score
			game.score = 0
		}
	}
	
	return lastGameScore
}
