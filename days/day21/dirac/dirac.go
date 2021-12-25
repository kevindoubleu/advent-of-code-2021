package dirac

type Dirac struct {
	// player scores
	score1		int
	score2		int

	// player positions
	pos1		int
	pos2		int

	// dice being used in this game
	dice		Dice

	// how many points needed to win this game
	winScore	int
}

func NewPracticeGame(pos1, pos2, winScore int) Dirac {
	dice := newDeterministicDice()

	game := Dirac{
		pos1: pos1,
		pos2: pos2,
		dice: &dice,
		winScore: winScore,
	}

	return game
}

func (d *Dirac) Play() {
	turn := 1

	for d.score1 < d.winScore && d.score2 < d.winScore {
		diceRoll := d.dice.Next3()

		if turn == 1 {
			// player rolls dice and moves
			d.pos1 += diceRoll - 1
			d.pos1 = d.pos1 % 10 + 1

			// player gets score based on where he landed
			d.score1 += d.pos1

			// change player
			turn = 2
		} else {
			d.pos2 += diceRoll - 1
			d.pos2 = d.pos2 % 10 + 1
			d.score2 += d.pos2
			turn = 1
		}
	}
}

func (d Dirac) LoserTimesDiceRolls() int {
	if d.score1 >= 1000 {
		return d.score2 * d.dice.Rolls()
	}
	return d.score1 * d.dice.Rolls()
}
