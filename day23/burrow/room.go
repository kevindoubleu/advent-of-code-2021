package burrow

// basically like a stack
type Room struct {
	// 0 is bottom
	// last is top
	tenants []byte
	cap		int
}

func newRoom(cap int) Room {
	r := Room{
		tenants: make([]byte, cap),
		cap: cap,
	}
	return r
}

func cloneRoom(original Room) Room {
	r := Room{
		tenants: make([]byte, 0),
		cap: original.cap,
	}

	r.tenants = append(r.tenants, original.tenants...)

	return r
}

func (r *Room) push(amphipod byte) (success bool) {
	for i, tenant := range r.tenants {
		if tenant == 0 {
			r.tenants[i] = amphipod
			return true
		}
	}
	return false
}

func (r *Room) pop() byte {
	// find the topmost tenant
	var top byte
	for i := r.cap-1; i >= 0; i-- {
		if r.tenants[i] != 0 {
			top = r.tenants[i]
			r.tenants[i] = 0
			return top
		}
	}
	return 0
}

func (r *Room) top() byte {
	for i := 3; i >= 0; i-- {
		if r.tenants[i] != 0 {
			return r.tenants[i]
		}
	}
	return 0
}

// check if the room is homogenous
// only checks if all tenants are the same amphipods
func (r Room) homogenous() bool {
	for _, tenant := range r.tenants {
		if tenant != r.tenants[0] {
			return false
		}
	}
	return true
}

func (r Room) tenantCount() int {
	total := 0

	for _, tenant := range r.tenants {
		if tenant != r.tenants[0] {
			total++
		}
	}

	return total
}
