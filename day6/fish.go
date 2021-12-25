package main

type Fishes []int

func newFishes(initial []int) Fishes {
	return Fishes(initial)
}

func (f Fishes) count() int {
	return len(f)
}

func (f *Fishes) age1Day() {
	fishCount := f.count()
	for i := 0; i < fishCount; i++ {
		if (*f)[i] == 0 {
			(*f)[i] = 6
			(*f) = append((*f), 8)
		} else {
			(*f)[i]--
		}
	}
}

// age the fish for `days` days
func (f *Fishes) age(days int) {
	for i := 0; i < days; i++ {
		f.age1Day()
	}
}
