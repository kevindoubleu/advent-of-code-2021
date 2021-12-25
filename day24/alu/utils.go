package alu

import (
	"strconv"
	"strings"
)

var varIdx = map[byte]int{
	'w': 0,
	'x': 1,
	'y': 2,
	'z': 3,
}

func isVar(arg string) bool {
	return strings.ContainsAny(arg, "wxyz")
}

func (a *ALU) inp(arg string, value int) {
	a.vars[varIdx[arg[0]]] = value
}

func (a *ALU) arithmetic(op, arg1, arg2 string) {
	var1 := varIdx[arg1[0]]
	var2 := varIdx[arg2[0]]

	if isVar(arg2) {
		switch op {
		case "add":
			a.vars[var1] += a.vars[var2]
		case "mul":
			a.vars[var1] *= a.vars[var2]
		case "div":
			a.vars[var1] /= a.vars[var2]
		case "mod":
			a.vars[var1] %= a.vars[var2]
		}
	} else {
		val, _ := strconv.Atoi(arg2)
		switch op {
		case "add":
			a.vars[var1] += val
		case "mul":
			a.vars[var1] *= val
		case "div":
			a.vars[var1] /= val
		case "mod":
			a.vars[var1] %= val
		}
	}
}

func (a *ALU) eql(arg1, arg2 string) {
	var1 := varIdx[arg1[0]]
	var2 := varIdx[arg2[0]]
	var isEqual bool

	if isVar(arg2) {
		isEqual = a.vars[var1] == a.vars[var2]
	} else {
		val, _ := strconv.Atoi(arg2)
		isEqual = a.vars[var1] == val
	}

	if isEqual {
		a.vars[var1] = 1
	} else {
		a.vars[var1] = 0
	}
}


// translate model num to actual integer
func modelToInt(modelNumber []int) int {
	modelAsInt := 0

	for _, num := range modelNumber {
		modelAsInt += num
		modelAsInt *= 10
	}

	return modelAsInt / 10
}

// translate int to model num
func intToModel(modelNumber int) []int {
	modelAsSlice := make([]int, 0)

	// append number from back to front
	for modelNumber > 0 {
		modelAsSlice = append(modelAsSlice, modelNumber % 10)
		modelNumber /= 10
	}

	// reverse the slice
	for i, j := 0, len(modelAsSlice)-1; i < j; i, j = i+1, j-1 {
		modelAsSlice[i], modelAsSlice[j] = modelAsSlice[j], modelAsSlice[i]
	}

	return modelAsSlice
}
