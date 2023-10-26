package nodes

import (
	"fmt"
)

//////////
// Tree
//////////

// Tree header of a MC Tree
type Tree struct {
	V string // Game variant ID
	R *Node  // Root node of the MC Tree
}

// func (t *Tree) GetTree() *Tree {
// 	return t
// }

// func (t *Tree) GetNode() *Node {
// 	return nil
// }

// Init initialize a new tree
func Init(variant string) *Tree {
	if variant == "" {
		return nil
	}

	head := &Tree{
		V: variant,
	}

	root := Node{
		Parent: nil,
		R:      1,
	}
	head.R = &root
	return head
}

func (t *Tree) String() string {
	return fmt.Sprintf("Variant:%v", t.V)
}
