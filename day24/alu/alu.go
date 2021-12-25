package alu

import (
	"fmt"
	"strings"
)

type ALU struct {
	// the instructions in the monad
	monad	[]string

	// the model number
	model	[]int

	// internal variables
	// w, x, y, z
	vars	[]int
}

func NewALU(monad []string) ALU {
	a := ALU{
		monad : monad,
		model : []int{9,9,9,9,9,9,9,9,9,9,9,9,9,9},
		// model : []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1},
		vars  : make([]int, 4),
	}
	return a
}

// check if the model number is valid according to the monad
func (a *ALU) Validate() bool {
	// reset vars
	a.vars = make([]int, 4)

	// the current index of model number being inputted into "inp w"
	digit := 0

	for _, instruction := range a.monad {
		parts := strings.Split(instruction, " ")
		switch parts[0] {
		case "inp":
			a.inp(parts[1], a.model[digit])
		case "add", "mul", "div", "mod":
			a.arithmetic(parts[0], parts[1], parts[2])
		case "eql":
			a.eql(parts[1], parts[2])
		}
		// fmt.Println(instruction)
		// fmt.Println("vars:")
		// fmt.Println("w =", a.vars[0])
		// fmt.Println("x =", a.vars[1])
		// fmt.Println("y =", a.vars[2])
		// fmt.Println("z =", a.vars[3])
		// fmt.Fscanln(os.Stdin)
	}

	return a.vars[3] == 0
}

// assigns the biggest valid model number to a.model
func (a *ALU) BiggestModelNumber() {
	// for each position try from 1..9
	for !a.Validate() {
		fmt.Println("valid fail for model", a.model, "got", a.vars)
		a.model = intToModel(modelToInt(a.model) - 1)
	}
}
