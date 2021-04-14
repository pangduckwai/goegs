package node

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

////////////////////////////////////////////////////////////////////////////////////////////////
// NOTES TO SELF
// Why a node only need to contain the move taken, and not the current game board situation?
// - It is because the game's current situation is recorded in the tree structure.
// - e.g.:
// 1. at the start of game valid moves for the 1st player is put 1 troop on any territory
// 2. 1st player chose to put 1 troop in Brazil. a node is added to the root node
// 3. the tree got a branch from Start-of-Game to Claiming-Brazil
// 4. the 2nd player now can only choose any territory except Brazil
// 5. this way each node in the tree already represent a particular game board situation
////////////////////////////////////////////////////////////////////////////////////////////////

// Node each step of the game play
type Node struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Refn   primitive.ObjectID `bson:"f,omitempty"` // Reference to a parent node
	Impl   string             `bson:"p,omitempty"` // Game implementation string
	Data   uint32             `bson:"d"`
	Value  uint32             `bson:"v,omitempty"`
	Moves  uint16             `bson:"x"`           // Number of all possible moves
	Runs   uint64             `bson:"r,omitempty"` // Total number of runs
	Wins   uint64             `bson:"w,omitempty"` // Total number of won games of this player
	Next   []*Node            `bson:"n,omitempty"`
	Parent *Node              `bson:"-"`
	Level  int                `bson:"-"`
}

// New create a new node
// NOTE: not adding the newly created node to the parent's []Next here because Expand() may fail
func New(node *Node, impl string, turn uint8, action uint8, idx []uint8, val uint32) *Node {
	next := Node{}
	if node != nil {
		next = Node{
			Data:   0,
			Moves:  0,
			Runs:   0,
			Wins:   0,
			Parent: node,
			Next:   nil,
		}
		next.SetTurn(turn)
		next.SetAction(action)
	} else {
		next.Impl = impl
	}

	lidx := len(idx)
	if lidx > 0 {
		next.SetValue1(idx[0])
		if lidx > 1 {
			next.SetValue2(idx[1])
			if lidx > 2 {
				next.SetValue3(idx[2])
				if lidx > 3 {
					next.SetValue4(idx[3])
					if lidx > 4 {
						next.SetValue5(idx[4])
					}
				}
			}
		}
	}

	if val > 0 {
		next.Value = val
	}

	return &next
}

// Select select the optimal node from the existing tree
// return Select-node, Status, Error
func Select(node *Node, turn uint8, nextCard uint8, wons bool, battle func(attacker, defender uint32) uint8) (*Node, int, error) {
	idx, err := Ucb(node)
	if err != nil {
		return nil, 0, err
	}

	next := node.Next[idx]
	next.Parent = node
	act := next.Action()
	switch act {
	case 4:
		val := uint8(63)
		if wons {
			val = nextCard
		}
		for _, c := range next.Next {
			if c.Value1() == val {
				return c, 1, nil // 4a
			}
		}
		derv := New(next, "", turn, 8+act, []uint8{val}, 0)
		next.Next = append(next.Next, derv)
		return derv, 8, nil // 4b
	case 6:
		val := uint8(63)
		if wons {
			val = nextCard
		}
		for _, c := range next.Next {
			if c.Value3() == val {
				return c, 2, nil // 6a
			}
		}
		derv := New(next, "", turn, 8+act, []uint8{next.Value1(), next.Value2(), val}, next.Value)
		next.Next = append(next.Next, derv)
		return derv, 16, nil // 6b
	case 5:
		ridx := battle(uint32(next.Value4()), uint32(next.Value5()))
		for _, c := range next.Next {
			if c.Value3() == ridx {
				return c, 4, nil // 5a
			}
		}
		derv := New(next, "", turn, 8+act, []uint8{next.Value1(), next.Value2(), ridx, next.Value4(), next.Value5()}, next.Value)
		next.Next = append(next.Next, derv)
		return derv, 32, nil // 5b
	}

	return next, 0, nil
}

// Expand expand the tree to explore possible moves
func Expand(
	node *Node, isSim bool, turn uint8, nextCard uint8, wons bool, cardCount int,
	chooseMove func() (uint8, []uint8, uint32, uint16), battle func(attacker, defender uint32) uint8,
) (*Node, int, int) {
	var next *Node
	retry := ExpandRetries
	for retry > 0 {
		act, idx, val, mov := chooseMove() // Choose a legal move randomly

		// Build the node representing the move just chosen
		next = New(node, "", turn, act, idx, val)

		// Test if the chosen move was already explored
		if !node.Exists(next) {
			if !isSim {
				if len(node.Next) <= 0 {
					node.Moves = mov
				}
				node.Next = append(node.Next, next)
			}

			ridx := uint8(63)
			switch {
			case act == 4 || act == 6:
				if wons {
					ridx = nextCard
				}
				derv := New(next, "", turn, 8+act, append(idx, ridx), val)
				next.Moves = uint16(cardCount)
				next.Next = append(next.Next, derv)
				return derv, 1, ExpandRetries - retry
			case act == 5:
				idx[2] = battle(uint32(idx[3]), uint32(idx[4]))
				derv := New(next, "", turn, 8+act, idx, val)
				next.Moves = 2
				next.Next = append(next.Next, derv)
				return derv, 2, ExpandRetries - retry
			}

			return next, 0, ExpandRetries - retry
		}

		retry--
	}
	return nil, 0, ExpandRetries // Exhausted all legal move !?
}

var mux sync.Mutex

func won(node *Node) {
	mux.Lock()
	defer mux.Unlock()
	node.Runs++
	node.Wins++
}

func lost(node *Node) {
	mux.Lock()
	defer mux.Unlock()
	node.Runs++
}

// BackProp back propagate phase
func BackProp(tail *Node, winner uint8, impl string) (int64, error) {
	count := int64(0)
	for node := tail; node != nil; node = node.Parent {
		if node.Turn() == winner {
			won(node)
		} else {
			lost(node)
		}

		if node.Parent == nil {
			if node.Impl != impl {
				return 0, Err(9)
			}
		}

		count++
	}

	return count, nil
}

// Ucb calculate scorce for each node
func Ucb(node *Node) (int, error) {
	ni := math.Log(float64(node.Runs))
	if node.Runs == 0 {
		ni = 1
	}

	max := -1.0
	var idx0 []int
	idx1 := []int{-1}
	for i, x := range node.Next {
		r := float64(x.Runs)
		if r == 0 {
			idx0 = append(idx0, i)
		} else {
			u := float64(x.Wins)/r + Exploration*math.Sqrt(ni/r)
			if u > max {
				max = u
				idx1 = idx1[:1] // Remove all indices with same old high score, since a new high score is found
				idx1[0] = i
			} else if u == max {
				idx1 = append(idx1, i)
			}
		}
	}

	lgth := len(idx0)
	if lgth == 1 {
		return idx0[0], nil
	} else if lgth > 1 {
		tmp := rand.Intn(lgth)
		return idx0[tmp], nil
	}

	if len(idx1) > 1 {
		return idx1[rand.Intn(len(idx1))], nil
	}
	if idx1[0] >= 0 {
		return idx1[0], nil
	}
	return -1, Err(8)
}

// Validate validate a MCT by checking the # of Runs of a node - 1 equals the sum of the number of Runs
// of all its direct child nodes 1 level below
func Validate(tree *Node) (bool, int) {
	var node *Node
	queue := []*Node{tree}
	count := 0
	for len(queue) > 0 {
		node = queue[0]
		queue = append(queue[1:], node.Next...)

		if len(node.Next) <= 0 {
			continue
		}

		sum := uint64(0)
		for _, x := range node.Next {
			sum += x.Runs
		}
		if sum != (node.Runs - 1) {
			fmt.Printf("[Error] Validation failed: %v vs %v\n%v\n", sum, node.Runs, node.Tree(5, 0))
			return false, count
		}

		count++
	}

	return true, count
}

///////////////////////
// Utilities methods //
///////////////////////

func (n *Node) String() string {
	return toString(n, 0)
}

// Tree display the structure of the tree
func (n *Node) Tree(lvls int, cnts int) string {
	return show("", n, 0, 0, 0, lvls, cnts)
}
func show(pfx string, n *Node, idx int, lst int, lvl int, max int, cnt int) string {
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

	frmt := "─"
	if cnt > 0 {
		dl := 1
		if max > 1 {
			dl = int(math.Log10(float64(max-1))) + 1
		}
		dn := 1
		if cnt > 1 {
			dn = int(math.Log10(float64(cnt-1))) + 1
		}
		temp := fmt.Sprintf("─[%%%dv/%%-%dv]", dl, dn)
		frmt = fmt.Sprintf(temp, lvl, idx)
	}
	buff += fmt.Sprintf(frmt+" %v\n", toString(n, max-lvl))
	if lgth > 0 && lvl < max {
		for i, x := range n.Next {
			buff += show(pfx+nxlv, x, i, lgth-1, lvl+1, max, cnt)
		}
	}
	return buff
}
func toString(n *Node, p int) string {
	idx := ""
	pad := strings.Repeat(" ", p)
	switch n.Action() {
	case 2:
		idx = fmt.Sprintf("%2v, -, -,-,-,   -", n.Value1())
	case 3:
		idx = fmt.Sprintf("%2v, -, -,-,-,   -", n.Value1())
	case 4:
		idx = fmt.Sprintf(" -, -, -,-,-,   -")
	case 12:
		idx = fmt.Sprintf("%2v, -, -,-,-,   -", n.Value1())
	case 5:
		idx = fmt.Sprintf("%2v,%2v, -,%v,%v,%4v", n.Value1(), n.Value2(), n.Value4(), n.Value5(), n.Value)
	case 13:
		idx = fmt.Sprintf("%2v,%2v,%2v,%v,%v,%4v", n.Value1(), n.Value2(), n.Value3(), n.Value4(), n.Value5(), n.Value)
	case 6:
		idx = fmt.Sprintf("%2v,%2v, -,-,-,%4v", n.Value1(), n.Value2(), n.Value)
	case 14:
		idx = fmt.Sprintf("%2v,%2v,%2v,-,-,%4v", n.Value1(), n.Value2(), n.Value3(), n.Value)
	case 7:
		idx = fmt.Sprintf("%2v,%2v,%2v,-,-,   -", n.Value1(), n.Value2(), n.Value3())
	}
	rt := "»"
	if n.Parent == nil {
		rt = "≡"
	} else if n.Current() {
		rt = "*"
	}
	cnt := len(n.Next)
	if n.Level > 0 { // Borrow node.Level (a transient field) to store the number of child nodes during tracking
		cnt = n.Level
	}
	return fmt.Sprintf(
		"%v%v|%d|%x[%-17v/%7v/%8v]%3v/%3v|%v",
		n.ID, pad, n.Turn(), n.Action(), idx, n.Wins, n.Runs, n.Moves, cnt, rt,
	)
}

// Same check if 2 nodes are the same
func (n *Node) Same(o *Node) bool {
	if n.Data != o.Data {
		return false
	}
	if n.Action() == 5 || n.Action() == 6 {
		if n.Value != o.Value {
			return false
		}
	}
	return true
}

// Exists find if given child node already exists
func (n *Node) Exists(c *Node) bool {
	for _, sibl := range n.Next {
		if c.Same(sibl) {
			return true
		}
	}
	return false
}

// Format of Node.Data:
//   3322 2222 2222 1111 1111 1100 0000 0000
//   1098 7654 3210 9876 5432 1098 7654 3210
//   .... .... .... .... .... .... .... .111 Turn
//   .... .... .... .... .... .... .111 1000 Action
//   .... .... .... .... ...1 1111 1000 0000 Value1
//   .... .... .... .111 1110 0000 0000 0000 Value2
//   ...1 1111 1000 0000 0000 0000 0000 0000 Value3
//   .... .... ...1 1000 0000 0000 0000 0000 Value4 (Dice R)
//   .... .... .110 0000 0000 0000 0000 0000 Value5 (Dice W)
//   1000 0000 0000 0000 0000 0000 0000 0000 Current (transient)

// Current check if node is current or not
func (n *Node) Current() bool {
	return ((n.Data & 0x80000000) != 0)
}

// SetCurrent mark the node as current
func (n *Node) SetCurrent() {
	n.Data = n.Data | 0x80000000
}

// Turn get turn value
func (n *Node) Turn() uint8 {
	return uint8(n.Data & 0x00000007)
}

// SetTurn set turn value
func (n *Node) SetTurn(t uint8) {
	n.Data = n.Data | uint32(t&0x00000007)
}

// Action get action value
func (n *Node) Action() uint8 {
	return uint8((n.Data & 0x00000078) >> 3)
}

// SetAction set action value
func (n *Node) SetAction(a uint8) {
	n.Data = n.Data | ((uint32(a) << 3) & 0x00000078)
}

// Value1 get short value #1 (was Index1)
func (n *Node) Value1() uint8 {
	return uint8((n.Data & 0x00001F80) >> 7)
}

// SetValue1 set short value #1
func (n *Node) SetValue1(v uint8) {
	n.Data = n.Data | ((uint32(v) << 7) & 0x00001F80)
}

// Value2 get short value #2 (was Index2)
func (n *Node) Value2() uint8 {
	return uint8((n.Data & 0x0007E000) >> 13)
}

// SetValue2 set short value #2
func (n *Node) SetValue2(v uint8) {
	n.Data = n.Data | ((uint32(v) << 13) & 0x0007E000)
}

// Value3 get short value #3 (was Index3)
func (n *Node) Value3() uint8 {
	return uint8((n.Data & 0x1F800000) >> 23)
}

// SetValue3 set short value #3
func (n *Node) SetValue3(v uint8) {
	n.Data = n.Data | ((uint32(v) << 23) & 0x1F800000)
}

// Value4 get tiny value #4
func (n *Node) Value4() uint8 {
	return uint8((n.Data & 0x00180000) >> 19)
}

// SetValue4 set tiny value #4
func (n *Node) SetValue4(v uint8) {
	n.Data = n.Data | ((uint32(v) << 19) & 0x00180000)
}

// Value5 get tiny value #5
func (n *Node) Value5() uint8 {
	return uint8((n.Data & 0x00600000) >> 21)
}

// SetValue5 set tiny value #5
func (n *Node) SetValue5(v uint8) {
	n.Data = n.Data | ((uint32(v) << 21) & 0x00600000)
}
