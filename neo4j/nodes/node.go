package nodes

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

///////////////////////////////////////////////////////////////////////////////////////////////
// NOTES TO SELF
// Why a node only need to contain the move taken, and not the current game board situation?
// - It is because the game's current situation is recorded in the tree structure.
// - e.g.:
// 1. at the start of game valid moves for the 1st player is put 1 troop on any territory
// 2. 1st player chose to put 1 troop in Brazil. a node is added to the root node
// 3. the tree got a branch from Start-of-Game to Claiming-Brazil
// 4. the 2nd player now can only choose any territory except Brazil
// 5. this way each node in the tree already represent a particular game board situation
///////////////////////////////////////////////////////////////////////////////////////////////

// Exploration the exploration parameter in the Upper Confidence Bound (UCB) algorithm.
// Theoretically it equals to √2, but in practice usually chosen empirically.
const Exploration = 1.4142135623730950488016887242097 // √2

//////////
// Node
//////////

// Maximum value of node.Value
const MaxValue = 65535 // 16 bits

type Nid interface{}

// Node represent a node in the MC Tree, correspond to each step of the game play.
type Node struct {
	Next   []*Node // Child nodes
	Parent *Node   // Parent node
	D      uint32  // Node data
	V      uint16  // data value
	R      uint64  // Total number of runs
	W      uint64  // Total number of won games of this player
	L      int     // Level in the MC Tree (transient)
}

// func (n *Node) GetTree() *Tree {
// 	return nil
// }

// func (n *Node) GetNode() *Node {
// 	return n
// }

// New create a new node
// NOTE: not adding the newly created node to the parent's []Next here because Expand() may fail
func New(prnt *Node, turn uint8, action uint8, vals []uint8, val uint16) *Node {
	// if prnt == nil {
	// 	return nil
	// }

	next := Node{
		Parent: prnt,
	}
	next.SetTurn(turn)
	next.SetAction(action)

	for i := 0; i < len(vals); i++ {
		next.SetParam(i+1, vals[i])
	}

	if val > 0 {
		next.V = val
	}

	return &next
}

//////////////////
// Node methods
//////////////////

// Sorted for sorting child nodes array
type Sorted []*Node

func (nodes Sorted) Len() int {
	return len(nodes)
}
func (nodes Sorted) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
func (nodes Sorted) Less(i, j int) bool {
	return nodes[i].Less(nodes[j])
}

// Less for sorting the child array in Node
func (n *Node) Less(l *Node) bool {
	return Less(n.D, n.V, l.D, l.V)
}

// Same check if 2 nodes are the same
func (n *Node) Same(o *Node) bool {
	return Same(n.D, n.V, o.D, o.V)
}

// Exists find if given child node already exists
func (n *Node) Exists(c *Node) (found bool, index int) {
	if c == nil {
		return false, -1
	}

	length := len(n.Next)
	index = sort.Search(length, func(i int) bool {
		return !n.Next[i].Less(c)
	})

	found = (index < length) && n.Next[index].Same(c)
	return
}

// Filter filter out unwanted child node denoted by the given index from the child nodes array
func (n *Node) Filter(indices []int) []*Node {
	rst := make([]*Node, len(n.Next))
	copy(rst, n.Next)
	idx := sort.IntSlice(indices)
	off := 0
	idx.Sort()

	for _, i := range idx {
		l := len(rst) - 1
		j := i - off
		if j == 0 {
			rst = rst[1:]
			off++
		} else if j == l {
			rst = rst[:l]
			off++
		} else if j > 0 && j < l {
			rst = append(rst[:j], rst[j+1:]...)
			off++
		}
	}

	return rst
}

// AddChild add a node as a child of the given node while keeping the sort order of the child nodes
func (n *Node) AddChild(c *Node) (found bool, idx int, err error) {
	if found, idx = n.Exists(c); !found {
		if idx < 0 {
			err = Err(fmt.Sprintf("Unable to add %v as child of %v", c, n))
			return
		}

		if idx >= len(n.Next) {
			idx = len(n.Next)
			n.Next = append(n.Next, c)
		} else {
			n.Next = append(n.Next[:idx+1], n.Next[idx:]...)
			n.Next[idx] = c
		}
		c.Parent = n
	} else {
		if idx < 0 || idx >= len(n.Next) {
			err = Err(fmt.Sprintf("Fatal error adding %v as child of %v", c, n))
		}
	}
	return
}

// // BackTrack back track the path of a given node
// // NOTE: do not mix up with "back propagate"
// func (n *Node) BackTrack() (path []uint8, variant string, err error) {
// 	c := n
// 	for ; c.Parent.GetNode() != nil; c = c.Parent.GetNode() {
// 		if found, idx := c.Parent.GetNode().Exists(c); found {
// 			path = append(path, uint8(idx))
// 		} else {
// 			err = Err(fmt.Sprintf("%v not found in its parent node", c))
// 			return
// 		}
// 	}

// 	if c.Parent.GetTree() == nil {
// 		err = Err("Root not reached")
// 		return
// 	}
// 	variant = c.Parent.GetTree().V

// 	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
// 		path[i], path[j] = path[j], path[i]
// 	}

// 	return
// }

// ClearTransient clear transient fields
func (n *Node) ClearTransient() {
	n.D = ClearTransient(n.D)
}

// Busy check if node is busy
func (n *Node) Busy() bool {
	return Busy(n.D)
}

// SetBusy mark the node as busy
func (n *Node) SetBusy() {
	n.D = SetBusy(n.D)
}

// ClearBusy clear the busy flag
func (n *Node) ClearBusy() {
	n.D = ClearBusy(n.D)
}

// Current check if node is current or not
func (n *Node) Current() bool {
	return Current(n.D)
}

// SetCurrent mark the node as current
func (n *Node) SetCurrent() {
	n.D = SetCurrent(n.D)
}

// ClearCurrent clear the current flag
func (n *Node) ClearCurrent() {
	n.D = ClearCurrent(n.D)
}

// EndGame indicate a game ending node
func (n *Node) EndGame() bool {
	return EndGame(n.D)
}

// SetEndGame mark the node as game ending
func (n *Node) SetEndGame() {
	n.D = SetEndGame(n.D)
}

// ClearEndGame clear the end-game flag
func (n *Node) ClearEndGame() {
	n.D = ClearEndGame(n.D)
}

// Turn get turn value
func (n *Node) Turn() uint8 {
	return Turn(n.D)
}

// SetTurn set turn value
func (n *Node) SetTurn(t uint8) {
	n.D = SetTurn(n.D, t)
}

// Action get action value
func (n *Node) Action() uint8 {
	return Action(n.D)
}

// SetAction set action value
func (n *Node) SetAction(a uint8) {
	n.D = SetAction(n.D, a)
}

// Param get short value
func (n *Node) Param(i int) uint8 {
	return Param(n.D, i)
}

// SetParam set short value
func (n *Node) SetParam(i int, v uint8) {
	n.D = SetParam(n.D, i, v)
}

///////////////////////
// Utility functions
///////////////////////

// Ucb calculate scorce for each node
func Ucb(
	rnd *rand.Rand,
	node *Node,
	busy func(*Node) bool, // Func to indicate the node is busy, choose the next one
) (int, error) {
	var ni float64
	if node.R == 0 {
		ni = 1
	} else {
		ni = math.Log(float64(node.R))
	}

	max := -1.0
	var count int
	indices := []int{-1}
	for i, x := range node.Next {
		if x.R == 0 {
			if busy == nil || !busy(x) {
				return i, nil
			}
			count++
		} else {
			r := float64(x.R)
			u := float64(x.W)/r + Exploration*math.Sqrt(ni/r)
			if u > max {
				max = u
				indices = indices[:1] // Remove all indices with same old high score, since a new high score is found
				indices[0] = i
			} else if u == max {
				indices = append(indices, i)
			}
		}
	}

	if len(indices) > 1 {
		return indices[rnd.Intn(len(indices))], nil
	}
	if indices[0] >= 0 {
		return indices[0], nil
	}
	if count > 0 {
		return -1, nil // All unexplored node busy, but no node selected
	}
	return -1, Err("UCB failed")
}

// LeastExplored find the least explored next node
func LeastExplored(
	node *Node,
	busy func(*Node) bool,
) int {
	l := len(node.Next)
	if l <= 0 {
		return -2
	}

	i := 0
	for j := 1; j < l; j++ {
		if node.Next[j].R < node.Next[i].R {
			i = j
		}
	}

	if busy == nil || !busy(node.Next[i]) {
		return i
	}
	return -1
}

// Validate validate a single node by:
// 1. all direct children of the node are sorted
// 2. # of Runs of the node equals to the sum of the # of Runs (minus 1) of all its direct children nodes
// 3. # of Runs of the node equals to the sum of the # of Runs of all its derived children nodes
// 4. node marked as game ending move, but still have child nodes
func (n *Node) Validate() (
	flag int, // range from -4 to 1. -4 to -1 correspond to the 4 validation rules described above. 1 means validation pass, 0 means undetermined as child nodes array is empty.
	sum uint64, // the sum of the number of runs of the child nodes
	expected uint64, // the number of runs of the current node
) {
	if len(n.Next) > 0 {
		flag = 1

		if !sort.IsSorted(Sorted(n.Next)) {
			flag = -1
			return // Child nodes not sorted
		}

		if EndGame(n.D) {
			flag = -4
			return // Game ending nodes with children found
		}

		for _, x := range n.Next {
			sum += x.R
		}

		expected = n.R
		if HasDerived(n.D) {
			if sum != n.R {
				flag = -3
				return // # of runs mismatched
			}
		} else {
			if sum != (n.R - 1) {
				flag = -2
				return // # of runs mismatched
			}
		}
	} else {
		flag = 0
	}
	return
}

///////////////////////////////
// Display related functions
///////////////////////////////

func (n *Node) String() string {
	return ToString(n.D, n.V, n.W, n.R, 0, n.Parent == nil, len(n.Next), n.L)
}

// Tree display the in-memory structure of the tree.
// 'lvls'
//   - controls the maximum number of levels to display, which also affect the number of <space> padding before the node contents.
//
// 'cnts'
//   - the number of children of the node with the most child nodes in the entire (known) tree has, which is used to calculate the number of padding for the child node indices,
//     level and child node index display is turned off if 'cnts' equals 0.
//
// Both parameters affect the alignment of the displayed node contents.
func (n *Node) Tree(lvls int, cnts int) string {
	return show("", n, 0, 0, 0, lvls, cnts)
}

// show helper function called recursivly by node.Show()
func show(pfx string, n *Node, idx int, lst int, lvl int, lvls int, cnts int) (buff string) {
	if n == nil {
		return "[nil]"
	}
	buff = pfx
	lgth := len(n.Next)
	nxlv := " "
	if idx == lst {
		buff += "└"
	} else {
		buff += "├"
		nxlv = "│"
	}
	if lgth <= 0 || lvl >= lvls {
		buff += "─"
	} else {
		buff += "┬"
	}

	frmt := "─"
	if cnts > 0 {
		dl := 1
		if lvls > 1 {
			dl = int(math.Log10(float64(lvls-1))) + 1
		}
		dn := 1
		if cnts > 1 {
			dn = int(math.Log10(float64(cnts-1))) + 1
		}
		temp := fmt.Sprintf("─[%%%dv/%%-%dv]", dl, dn)
		frmt = fmt.Sprintf(temp, lvl, idx)
	}
	buff += fmt.Sprintf(frmt+"%v\n", ToString(n.D, n.V, n.W, n.R, lvls-lvl, n.Parent == nil, len(n.Next), n.L))
	if lgth > 0 && lvl < lvls {
		for i, x := range n.Next {
			buff += show(pfx+nxlv, x, i, lgth-1, lvl+1, lvls, cnts)
		}
	}
	return
}

// Err node errors
type Err string

func (e Err) Error() string {
	return fmt.Sprintf("[NODE] %v", string(e))
}
