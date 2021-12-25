package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// input := lib.ReadTest()
	input := lib.ReadInput()

	fmt.Println(position(input))
	fmt.Println(positionWithAim(input))
}

// 1
type Command struct {
	direction	string
	value		int
}
func parseCommand(command string) Command {
	parts := strings.Split(command, " ")
	
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	return Command{
		direction: parts[0],
		value: value,
	}
}

func position(commandList string) int {
	commands := lib.ToStrSlice(commandList)
	horizontal, depth := 0, 0

	for _, command := range commands {
		cmd := parseCommand(command)
		switch cmd.direction {
		case "forward":
			horizontal += cmd.value
		case "down":
			depth += cmd.value
		case "up":
			depth -= cmd.value
		}
	}

	return horizontal * depth
}

// 2
func positionWithAim(commandList string) int {
	commands := lib.ToStrSlice(commandList)
	horizontal, depth := 0, 0
	aim := 0

	for _, command := range commands {
		cmd := parseCommand(command)
		switch cmd.direction {
		case "forward":
			horizontal += cmd.value
			depth += aim * cmd.value
		case "down":
			aim += cmd.value
		case "up":
			aim -= cmd.value
		}
	}

	return horizontal * depth
}
