package probe

type Probe struct {
	position	Coordinate
	velocity	Coordinate
	target		Area
}

func NewProbe(veloX, veloY int, target string) Probe {
	p := Probe{
		position : Coordinate{0,0},
		velocity : Coordinate{veloX, veloY},
		target   : NewArea(target),
	}
	return p
}

// go forward in time to calculate the probe's next location
func (p *Probe) step() {
	// The probe's x position increases by its x velocity.
	p.position.x += p.velocity.x

	// The probe's y position increases by its y velocity.
	p.position.y += p.velocity.y

	// Due to drag, the probe's x velocity changes by 1 toward the value 0;
	// that is, it decreases by 1 if it is greater than 0,
	// increases by 1 if it is less than 0,
	// or does not change if it is already 0.
	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}

	// Due to gravity, the probe's y velocity decreases by 1.
	p.velocity.y--
}

func (p *Probe) StepMultiple(times int) {
	for i := 0; i < times; i++ {
		p.step()
	}
}

// given a probe's position and velocity
// determines if it will reach it's target at any step
func (p *Probe) WillBeInTarget() bool {
	for {
		if p.position.isWithin(p.target) {
			return true
		}

		// if a probe has gone below it's target
		// it will never go back up, meaning
		// it will never reach it's target
		if p.position.isBelow(p.target) {
			return false
		}

		// if a probe has gone past it's target
		// it will never go backwards due to drag, meaning
		// it will never reach it's target
		if p.position.isAfter(p.target) {
			return false
		}

		p.step()
	}
}
