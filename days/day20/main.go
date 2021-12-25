package main

import (
	"advent-of-code-2021/days/day20/image"
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

	var imageAlgo, imagePixels []string
	lib.UnpackToStrSlices(input, &imageAlgo, &imagePixels)
	// 1
	img := image.NewImage(imageAlgo[0], imagePixels)
	img.EnhanceMultiple(2)
	fmt.Println(img.LivePixelCount())
	
	// 2
	img = image.NewImage(imageAlgo[0], imagePixels)
	img.EnhanceMultiple(50)
	fmt.Println(img.LivePixelCount())
}
