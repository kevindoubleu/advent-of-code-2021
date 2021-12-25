package main

import (
	"advent-of-code-2021/days/day14/polygrp"
	"advent-of-code-2021/days/day14/polymer"
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

	var templates, rules []string
	lib.UnpackToStrSlices(input, &templates, &rules)

	// 1
	p := polymer.NewPolymer(templates[0], rules)
	p.PolymerizeMultiple(10)
	fmt.Println("most common", p.MostCommonCharCount())
	fmt.Println("least common", p.LeastCommonCharCount())
	fmt.Println("difference", p.MostCommonCharCount() - p.LeastCommonCharCount())
	
	fmt.Println("")

	// 2
	pg := polygrp.NewPolymerGroup(templates[0], rules)
	pg.PolymerizeMultiple(40)
	fmt.Println("most common", pg.MostCommonCharCount())
	fmt.Println("least common", pg.LeastCommonCharCount())
	fmt.Println("difference", pg.MostCommonCharCount() - pg.LeastCommonCharCount())
}
