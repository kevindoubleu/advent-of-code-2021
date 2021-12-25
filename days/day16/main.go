package main

import (
	"advent-of-code-2021/days/day16/packet"
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

	// 1
	p := packet.NewPacket(input)
	fmt.Println(p.TotalVersionNumber())

	// 2
	fmt.Println(p.Value)
}
