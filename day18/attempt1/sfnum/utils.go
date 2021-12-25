package sfnum

import (
	"fmt"
	"strings"
)

func (n *SFNumber) Print() {
	fmt.Println("")
	n._print("")
	fmt.Println("")
}
func (n *SFNumber) _print(space string) {
	if n == nil { return }
	space += "   "
	
	n.right._print(space)
	fmt.Printf("%s%d\n", space, n.value)
	n.left._print(space)
}

// does not change the number back to nested array form
// but just enough to determine if a node has changed
func (n SFNumber) String() string {
	result := strings.Builder{}

	n.inorder(&result)

	return result.String()
}
func (n *SFNumber) inorder(result *strings.Builder) {
	if n == nil { return }

	n.left.inorder(result)
	// just print leaf nodes
	if n.value != -1 {
		result.WriteString(fmt.Sprint(n.value, ","))
	}
	n.right.inorder(result)
}




func (n *SFNumber) isBranch() bool {
	if n == nil { return false }

	return n.value == branchCode
}

func (n *SFNumber) isLeaf() bool {
	if n == nil { return false }

	return n.value != branchCode &&
		n.left == nil &&
		n.right == nil
}

// check if left and right are regular nums
func (n *SFNumber) isBranchToLeaves() bool {
	if n == nil { return false }
	if n.left == nil || n.right == nil { return false }

	return !n.left.isBranch() && !n.right.isBranch()
}


// 1
func AddMultiple(numbers []string) *SFNumber {
	mainNumber := NewSFNumberFromString(numbers[0])

	for i := 1; i < len(numbers); i++ {
		fmt.Println("now", mainNumber)
		mainNumber = mainNumber.Add(NewSFNumberFromString(numbers[i]))
	}

	return mainNumber
}
