package main

import (
	"advent-of-code-2021/day17/probe"
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
	maxY := highestVeloY(input)
	fmt.Println("maxY at", maxY, "with height", probe.GetMaxVelocity(maxY))
	
	// 2
	fmt.Println(allShots(input))
}

// 1
// returns highest y position to shoot while still being on target
func highestVeloY(target string) int {
	targetArea := probe.NewArea(target)
	veloX, veloY := 0, 0
	maxY := 0

	// to find the highest Y value, we start from 0 and
	// keep going up as long as there is at least one X value that will
	// make the probe reach the target
	
	// keep increasing Y until the first shot is higher than the distance
	// between 0 and target.YLow
	// for example Y = 11, target.YLow = -10
	// after we shoot it up, it will go back down to 0 at veloY of -11
	// then after being back at 0, the next step it will be at Y = -11
	// missing our target
	for veloY < difference(0, targetArea.YLow) {
		
		// as long as we will vertically align with our target area
		// keep increasing forward velo until we can reach target
		for probe.GetMaxVelocity(veloX) <= targetArea.XHigh {
			// fmt.Println("trying", veloX, veloY)

			currentProbe := probe.NewProbe(veloX, veloY, target)
			if currentProbe.WillBeInTarget() {
				// fmt.Println("found valid velocity", veloX, veloY)
				maxY = veloY
				break
			}

			veloX++
		}

		// keep going up
		veloY++
		// but reset the forward force
		veloX = 0
	}

	return maxY
}

// 2
func allShots(target string) int {
	targetArea := probe.NewArea(target)
	validShots := 0

	// start from X = 0, until target.XHigh
	for i := 0; i <= targetArea.XHigh; i++ {
		// start from Y = maxY, until target.YLow
		for j := highestVeloY(target); j >= targetArea.YLow; j-- {

			p := probe.NewProbe(i, j, target)
			if p.WillBeInTarget() {
				validShots++
			}
		}
	}

	return validShots
}
