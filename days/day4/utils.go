package main

import (
	"strconv"
	"strings"
)

func getCalls(fileContent string) []int {
	callsByComa := strings.Split(fileContent, "\n")[0]
	callsAsString := strings.Split(callsByComa, ",")
	calls := []int{}

	for _, call := range callsAsString {
		if callNum, err := strconv.Atoi(call); err != nil {
			panic(err)
		} else {
			calls = append(calls, callNum)
		}
	}

	return calls
}

func getBoards(fileContent string) []Board {
	boards := make([]Board, 0)
	
	// split the input into boards
	boardStrings := strings.Split(fileContent, "\n\n")[1 : ] // remove the number calls
	// parse each board
	for _, boardString := range boardStrings {
		var board Board

		// split each board into rows
		rows := strings.Split(boardString, "\n")
		// parse each row
		for i, row := range rows {

			// split into cells
			cells := strings.Fields(row)
			// parse each cell
			for j, cell := range cells {

				if num, err := strconv.Atoi(cell); err != nil {
					panic(err)
				} else {
					board[i][j] = Number{
						value  : num,
						marked : false,
					}
				}
			}
		}

		boards = append(boards, board)
	}

	return boards
}