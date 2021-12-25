package sfnum

import (
	"strconv"
)

// essentially a binary tree
type SFNumber struct {
	// left and right item if this number is a pair
	left	*SFNumber
	right	*SFNumber

	// the parent of current node
	parent	*SFNumber

	// the value if this number is a regular number
	value	int
}

const (
	left = iota
	right
)

const (
	branchCode = -1
)

func NewSFNumberFromString(str string) *SFNumber {
	root := &SFNumber{}
	leftOrRight := left

	for _, char := range str {
		switch char {

		// go down branch
		// force next branch to be left
		case '[':
			newNode := &SFNumber{
				parent : root,
				value  : branchCode, // assumes input is always positive number
			}

			if leftOrRight == left {
				root.left = newNode
				root = root.left
			} else if leftOrRight == right {
				root.right = newNode
				root = root.right
			}
			leftOrRight = left
		
		// force next branch to be right
		case ',':
			leftOrRight = right

		// force next branch to be left, then
		// go up branch
		case ']':
			leftOrRight = left
			root = root.parent
		
		// create new regular number
		default:
			value, _ := strconv.Atoi(string(char))
			leaf := root.newLeaf(value)
			if leftOrRight == left {
				root.left = leaf
			} else if leftOrRight == right {
				root.right = leaf
			}
		}
	}

	// the outermost [] makes this tree go left 1x
	// before doing anything
	// so we undo it here
	root.left.parent = nil
	return root.left
}

func (n *SFNumber) newLeaf(value int) *SFNumber {
	leaf := SFNumber{
		value: value,
		parent: n,
	}
	return &leaf
}




func (n *SFNumber) Explode() {
	n._inorderExplode(0)
}
func (n *SFNumber) _inorderExplode(level int) {
	if n == nil { return }

	// if node is too deep, has to be exploded
	if level >= 4 && n.isBranchToLeaves() {
		// from the problem:
		// Exploding pairs will always consist of two regular numbers.

		// give value to parents
		n.parent.addToParent(left, n.left.value)
		n.parent.addToParent(right, n.right.value)
		// change myself into a leaf node with value 0
		n.parent.changeExplodedChild()

		return
	}

	// if leaf node
	if n.left == nil && n.right == nil {
		return
	}

	// if branch node
	n.left._inorderExplode(level+1)
	n.right._inorderExplode(level+1)
}

func (n *SFNumber) addToParent(direction, value int) {
	origin := n

	// go up until we're at root
	for n.parent != nil {

		// if found a leaf node, add to it
		if direction == left && n.left.isLeaf() {
			n.left.value += value
			return
		} else if direction == right && n.right.isLeaf() {
			n.right.value += value
			return
		}

		// if not then go up
		origin = n
		n = n.parent
	}

	// if we're at root and still cant find a way
	// try going down the branch where we want to put the value
	// for example if direction == right, then we go down right subtree
	// but the branch must be different then where we came from
	if direction == right {
		if n.right == origin {
			return
		}

		for n != nil {
			if n.left.isLeaf() {
				n.left.value += value
				return
			}
			
			n = n.right
		}
	} else if direction == left {
		if n.left == origin {
			return
		}

		for n != nil {
			if n.right.isLeaf() {
				n.right.value += value
				return
			}
			
			n = n.left
		}
	}

	// if still not found then the value is gone
}

func (n *SFNumber) changeExplodedChild() {
	zeroLeaf := n.newLeaf(0)

	if n.left.isBranchToLeaves() {
		n.left = zeroLeaf
	} else if n.right.isBranchToLeaves() {
		n.right = zeroLeaf
	}
}

func (n *SFNumber) Split() (done bool) {
	if n == nil { return }

	done = n.left.Split()
	if n.value > 9 {
		n.left = n.newLeaf(n.value / 2)
		n.right = n.newLeaf((n.value + 1) / 2)
		n.value = -1

		return true
	}

	if !done {
		n.right.Split()
	}
	return done
}



func (n *SFNumber) combineWith(otherNumber *SFNumber) *SFNumber {
	newRoot := &SFNumber{
		left: n,
		right: otherNumber,
		value: branchCode,
	}

	n.parent = newRoot
	otherNumber.parent = newRoot

	return newRoot
}

func (n *SFNumber) Add(otherNumber *SFNumber) *SFNumber {
	// create new root
	n = n.combineWith(otherNumber)

	// repeat until nothing changes
	reduced := n.String()
	for {
		n.Explode()
		n.Split()

		if n.String() == reduced {
			break
		} else {
			reduced = n.String()
		}
	}
	
	return n
}
