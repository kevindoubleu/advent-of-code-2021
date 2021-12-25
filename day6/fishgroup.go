package main

// this file counts the fish by groups instead of individually

type FishGroup struct {
	timers map[int]int // map of fish timers -> counts
}

func newFishGroup(initial []int) FishGroup {
	fg := FishGroup{
		timers: make(map[int]int),
	}

	for _, fish := range initial {
		fg.timers[fish]++
	}

	return fg
}

func (fg FishGroup) count() int {
	count := 0
	for _, fishCount := range fg.timers {
		count += fishCount
	}
	return count
}

func (fg *FishGroup) age1Day() {
	// age 0 are changed to 6 and add a new age 8 fish later
	justGaveBirth := fg.timers[0]
	fg.timers[0] = 0

	// the rest are moved
	for i := 1; i <= 8; i++ {
		fg.timers[i-1] = fg.timers[i]
	}

	// dont overwrite 6 bcs there might be fishes aging from 7 -> 6
	fg.timers[6] += justGaveBirth
	// overwrite the newborn babies bcs they have all aged to 7
	fg.timers[8] = justGaveBirth
}

func (fg *FishGroup) age(days int) {
	for i := 0; i < days; i++ {
		fg.age1Day()
	}
}
