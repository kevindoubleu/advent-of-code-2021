package universe

type Multiverse struct {
	// map of score -> count of universes where that player has that score
	// max len of this map will be 21 since that is the winning score
	p1scores	map[int]int
	p2scores	map[int]int
}

func NewMultiverse() Multiverse {
	m := Multiverse{
		p1scores: make(map[int]int),
		p2scores: make(map[int]int),
	}
	return m
}

func (m *Multiverse) Play() {
	// outcomes of the 3 dice rolls of each turn
	outcomes := []int{
		3, /* 111 */	4, /* 211 */	5, /* 311 */
		4, /* 112 */	5, /* 212 */	6, /* 312 */
		5, /* 113 */	6, /* 213 */	7, /* 313 */
		4, /* 121 */	5, /* 221 */	6, /* 321 */
		5, /* 122 */	6, /* 222 */	7, /* 322 */
		6, /* 123 */	7, /* 223 */	8, /* 323 */
		5, /* 131 */	6, /* 231 */	7, /* 331 */
		6, /* 132 */	7, /* 232 */	8, /* 332 */
		7, /* 133 */	8, /* 233 */	9, /* 333 */
	}

	// at every step, each player will get every outcome
	for _, outcome := range outcomes {
			
		// each player will have their scores increased by each outcome
		for score := range m.p1scores {
			m.p1scores[score + outcome] += m.p1scores[score]
		}

	}

	// each time a player wins, their total wins is increased by
	// how many scores the opponent still has that hasnt won yet
	// bcs this player is matched against all the other player's scores
}


