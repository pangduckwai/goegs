package main

import (
	"fmt"
	"strings"
)

// Node tree nodes
type Node struct {
	Row  int
	Col  int
	Next []*Node
}

func populate() *Node {
	rslt := &Node{0, 0,
		[]*Node{
			&Node{1, 0,
				[]*Node{
					&Node{2, 10, nil},
				},
			},
			&Node{1, 1,
				[]*Node{
					&Node{2, 20,
						[]*Node{
							&Node{3, 200, nil},
						},
					},
					&Node{2, 21, nil},
					&Node{2, 22, nil},
				},
			},
			&Node{1, 2, nil},
			&Node{1, 3,
				[]*Node{
					&Node{2, 40, nil},
					&Node{2, 41, nil},
					&Node{2, 42, nil},
				},
			},
		},
	}
	return rslt
}

// Tree display the tree
func (n *Node) Tree(max int) string {
	return show("", n, 0, 0, 0, max)
}
func show(pfx string, n *Node, idx int, lst int, lvl int, max int) string {
	buff := pfx
	lgth := len(n.Next)
	nxlv := " "
	if idx == lst {
		buff += "└"
	} else {
		buff += "├"
		nxlv = "│"
	}
	if lgth <= 0 || lvl >= max {
		buff += "─"
	} else {
		buff += "┬"
	}
	buff += "─ " + toString(n, max-lvl) + "\n"
	if lgth > 0 && lvl < max {
		for i, x := range n.Next {
			buff += show(pfx+nxlv, x, i, lgth-1, lvl+1, max)
		}
	}
	return buff
}
func toString(n *Node, p int) string {
	pad := strings.Repeat(" ", p)
	return fmt.Sprintf("*%v%2d|%-3d|%v)", pad, n.Row, n.Col, len(n.Next))
}

func main() {
	root := populate()
	fmt.Println(root.Tree(3))
}
