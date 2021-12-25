package dirac

type Dice interface {
	// returns the total steps from next 3 rolls
	Next3() int
	// returns the total times this dice has been rolled
	Rolls() int
}

type DeterministicDice struct {
	// the last value that showed up
	last	int
	rolls	int
}

func newDeterministicDice() DeterministicDice {
	dd := DeterministicDice{
		last: 0,
		rolls: 0,
	}
	return dd
}

func (dd *DeterministicDice) Next3() int {
	total := dd.last

	// 0 -> 1
	total += 1
	// 1 -> 1,1,1 (supposed to be 1,2,3)
	total *= 3
	// 1,1,1 -> 1,2,3
	total += 3

	dd.last += 3
	dd.rolls += 3
	return total
}

func (dd DeterministicDice) Rolls() int {
	return dd.rolls
}
