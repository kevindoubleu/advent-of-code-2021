package main

import (
	"fmt"
	"strings"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func parseFold(instruction string) (axis byte, value int) {
	instReader := strings.NewReader(instruction)
	fmt.Fscanf(instReader, "fold along %c=%d", &axis, &value)

	return axis, value
}
