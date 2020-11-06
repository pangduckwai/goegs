package mongo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const url = "mongodb://192.168.56.31:27017" // "mongodb://localhost:27017" "mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority"
const database = "wdom"
const collection = "test"

var conn bool
var mngo *mongo.Database
var coll *mongo.Collection
var cntx context.Context
var stop func()

func connect() error {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		defer cancel()
		return err
	}

	cntx = ctx
	mngo = client.Database(database)
	coll = mngo.Collection(collection)
	stop = func() {
		defer cancel()
		defer client.Disconnect(ctx)
		conn = false
	}
	conn = true
	return nil
}

// Node a tree node
type Node struct {
	ID     string  `bson:"_id,omitempty"` // primitive.ObjectID
	Ref    string  `bson:"e,omitempty"`   // Reference to a parent node
	Value  int     `bson:"v,omitempty"`
	Data   int     `bson:"d"`
	Parent *Node   `bson:"-"`
	Next   []*Node `bson:"n,omitempty"` // Total number of won games of this player
	Level  int     `bson:"-"`
}

func (n *Node) String() string {
	return toString(n, 0)
}

// Tree display the structure of the tree
func (n *Node) Tree(fwd int) string {
	return show("", n, 0, 0, 0, fwd)
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
	pad := strings.Repeat(".", p)
	return fmt.Sprintf(" %v. [ %-11v %4v ] %2v| %10v | %v -> %v", pad, n.Value, n.Data, 10-p, n.Level, n.ID, n.Ref)
}

func (n *Node) hash() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

// Write write to mongodb
func Write(root *Node) error {
	err := connect()
	if err != nil {
		return err
	}
	defer stop()

	vertices, err := Split(root, 3)
	if err != nil {
		return err
	}
	for i, vertex := range vertices {
		rst, err := coll.UpdateOne(
			cntx,
			bson.M{"_id": vertex.ID},
			bson.M{"$set": vertex},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return err
		}
		fmt.Println("Update results", vertex.ID, i, rst.MatchedCount, rst.ModifiedCount, rst.UpsertedCount)
	}

	return nil
}

// Read read from mongodb
func Read() (*Node, error) {
	err := connect()
	if err != nil {
		return nil, err
	}
	defer stop()

	var tree Node
	if err := coll.FindOne(cntx, bson.M{"_id": "TEST001"}).Decode(&tree); err != nil {
		return nil, err
	}

	var node *Node
	queue := []*Node{&tree}
	for len(queue) > 0 {
		node = queue[0]

		if node.Next == nil && node.ID != "" {
			cursor, err := coll.Find(cntx, bson.M{"e": node.ID})
			if err != nil {
				return nil, err
			}
			defer cursor.Close(cntx)

			result := make([]*Node, 0)
			for cursor.Next(cntx) {
				var read Node
				if err = cursor.Decode(&read); err != nil {
					return nil, err
				}
				result = append(result, &read)
			}
			node.Next = result
		}

		queue = append(queue[1:], node.Next...)
	}

	return &tree, nil
}

// Split split up the tree breadth-first by level
func Split(tree *Node, depth int) ([]*Node, error) {
	if depth < 3 || depth > 100 {
		return nil, errors.New("Depth must be >= 3 or <= 100")
	}

	var vertices []*Node
	var nxtLvl int
	var modlus int
	maxLvl := depth - 1
	tree.Level = 0

	var node *Node
	queue := []*Node{tree}
	for len(queue) > 0 {
		node = queue[0]

		modlus = node.Level % depth
		if modlus == maxLvl && node.ID == "" {
			node.ID = node.hash()
		} else if modlus == 0 {
			if node.ID == "" {
				node.ID = node.hash()
			}
			vertices = append(vertices, node)
		}

		nxtLvl = node.Level + 1
		for _, n := range node.Next {
			n.Level = nxtLvl
			if modlus == maxLvl {
				n.Ref = node.ID
			}
		}

		queue = append(queue[1:], node.Next...)
		if modlus == maxLvl {
			node.Next = nil
		}
	}

	return vertices, nil
}

// Change update the tree
func Change(tree *Node, code int) {
	find := func(tree *Node, val int) *Node {
		var node *Node
		queue := []*Node{tree}
		for len(queue) > 0 {
			node = queue[0]
			queue = append(queue[1:], node.Next...)
			if node.Value == val {
				return node
			}
		}
		return nil
	}

	switch code {
	case 1:
		node := find(tree, 1220112)
		node.Data += 17
	case 2:
		node := find(tree, 102)
		node.Next = []*Node{
			&Node{Value: 1020, Data: 79, Parent: node, Next: nil},
		}
	case 3:
		node := find(tree, 102)
		node.Next = append(node.Next, &Node{Value: 1021, Data: 7, Parent: node, Next: nil})
	}
}

// Same check if 2 MCTrees are the same
func Same(tree1, tree2 *Node) bool {
	build := func(tree *Node) []int {
		var rst []int

		var node *Node
		queue := []*Node{tree}
		for len(queue) > 0 {
			node = queue[0]
			queue = append(queue[1:], node.Next...)
			rst = append(rst, node.Value)
		}
		return rst
	}

	rst1 := build(tree1)
	rst2 := build(tree2)
	if len(rst1) != len(rst2) {
		return false
	}

	for i := 0; i < len(rst1); i++ {
		if rst1[i] != rst2[i] {
			return false
		}
	}
	return true
}

// Build build the test tree
func Build() *Node {
	// Level 0
	head := &Node{ID: "TEST001", Value: 1, Parent: nil, Next: nil}
	// Level 1
	head.Next = []*Node{
		&Node{Value: 10, Parent: head, Next: nil},
		&Node{Value: 11, Parent: head, Next: nil},
		&Node{Value: 12, Parent: head, Next: nil},
		&Node{Value: 13, Parent: head, Next: nil},
		&Node{Value: 14, Parent: head, Next: nil},
	}
	// Level 2
	head.Next[0].Next = []*Node{
		&Node{Value: 100, Parent: head.Next[0], Next: nil},
		&Node{Value: 101, Parent: head.Next[0], Next: nil},
		&Node{Value: 102, Parent: head.Next[0], Next: nil},
	}
	head.Next[2].Next = []*Node{
		&Node{Value: 120, Parent: head.Next[2], Next: nil},
		&Node{Value: 121, Parent: head.Next[2], Next: nil},
		&Node{Value: 122, Parent: head.Next[2], Next: nil},
		&Node{Value: 123, Parent: head.Next[2], Next: nil},
	}
	head.Next[3].Next = []*Node{
		&Node{Value: 130, Parent: head.Next[3], Next: nil},
		&Node{Value: 131, Parent: head.Next[3], Next: nil},
	}
	head.Next[4].Next = []*Node{
		&Node{Value: 140, Parent: head.Next[4], Next: nil},
		&Node{Value: 141, Parent: head.Next[4], Next: nil},
	}
	// Level 3
	head.Next[0].Next[1].Next = []*Node{
		&Node{Value: 1010, Parent: head.Next[0].Next[1], Next: nil},
		&Node{Value: 1011, Parent: head.Next[0].Next[1], Next: nil},
		&Node{Value: 1012, Parent: head.Next[0].Next[1], Next: nil},
	}
	head.Next[2].Next[2].Next = []*Node{
		&Node{Value: 1220, Parent: head.Next[2].Next[2], Next: nil},
		&Node{Value: 1221, Parent: head.Next[2].Next[2], Next: nil},
	}
	head.Next[2].Next[3].Next = []*Node{
		&Node{Value: 1230, Parent: head.Next[2].Next[3], Next: nil},
		&Node{Value: 1231, Parent: head.Next[2].Next[3], Next: nil},
		&Node{Value: 1232, Parent: head.Next[2].Next[3], Next: nil},
	}
	head.Next[3].Next[0].Next = []*Node{
		&Node{Value: 1300, Parent: head.Next[3].Next[0], Next: nil},
		&Node{Value: 1301, Parent: head.Next[3].Next[0], Next: nil},
		&Node{Value: 1302, Parent: head.Next[3].Next[0], Next: nil},
	}
	// Level 4
	head.Next[0].Next[1].Next[1].Next = []*Node{
		&Node{Value: 10110, Parent: head.Next[0].Next[1].Next[1], Next: nil},
		&Node{Value: 10111, Parent: head.Next[0].Next[1].Next[1], Next: nil},
		&Node{Value: 10112, Parent: head.Next[0].Next[1].Next[1], Next: nil},
	}
	head.Next[0].Next[1].Next[2].Next = []*Node{
		&Node{Value: 10120, Parent: head.Next[0].Next[1].Next[2], Next: nil},
		&Node{Value: 10121, Parent: head.Next[0].Next[1].Next[2], Next: nil},
		&Node{Value: 10122, Parent: head.Next[0].Next[1].Next[2], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next = []*Node{
		&Node{Value: 12200, Parent: head.Next[2].Next[2].Next[0], Next: nil},
		&Node{Value: 12201, Parent: head.Next[2].Next[2].Next[0], Next: nil},
		&Node{Value: 12202, Parent: head.Next[2].Next[2].Next[0], Next: nil},
	}
	head.Next[2].Next[2].Next[1].Next = []*Node{
		&Node{Value: 12210, Parent: head.Next[2].Next[2].Next[1], Next: nil},
		&Node{Value: 12211, Parent: head.Next[2].Next[2].Next[1], Next: nil},
	}
	// Level 5
	head.Next[0].Next[1].Next[2].Next[1].Next = []*Node{
		&Node{Value: 101210, Parent: head.Next[0].Next[1].Next[2].Next[1], Next: nil},
		&Node{Value: 101211, Parent: head.Next[0].Next[1].Next[2].Next[1], Next: nil},
		&Node{Value: 101212, Parent: head.Next[0].Next[1].Next[2].Next[1], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[0].Next = []*Node{
		&Node{Value: 122000, Parent: head.Next[2].Next[2].Next[0].Next[0], Next: nil},
		&Node{Value: 122001, Parent: head.Next[2].Next[2].Next[0].Next[0], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[1].Next = []*Node{
		&Node{Value: 122010, Parent: head.Next[2].Next[2].Next[0].Next[1], Next: nil},
		&Node{Value: 122011, Parent: head.Next[2].Next[2].Next[0].Next[1], Next: nil},
		&Node{Value: 122012, Parent: head.Next[2].Next[2].Next[0].Next[1], Next: nil},
		&Node{Value: 122013, Parent: head.Next[2].Next[2].Next[0].Next[1], Next: nil},
	}
	// Level 6
	head.Next[0].Next[1].Next[2].Next[1].Next[0].Next = []*Node{
		&Node{Value: 1012100, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[0], Next: nil},
		&Node{Value: 1012101, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[0], Next: nil},
	}
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next = []*Node{
		&Node{Value: 1012110, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1], Next: nil},
		&Node{Value: 1012111, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1], Next: nil},
		&Node{Value: 1012112, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[1].Next[1].Next = []*Node{
		&Node{Value: 1220110, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1], Next: nil},
		&Node{Value: 1220111, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1], Next: nil},
		&Node{Value: 1220112, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1], Next: nil},
		&Node{Value: 1220113, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1], Next: nil},
	}
	// Level 7
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[0].Next = []*Node{
		&Node{Value: 10121100, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[0], Next: nil},
		&Node{Value: 10121101, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[0], Next: nil},
	}
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next = []*Node{
		&Node{Value: 10121120, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2], Next: nil},
		&Node{Value: 10121121, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[1].Next = []*Node{
		&Node{Value: 12201110, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[1], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2].Next = []*Node{
		&Node{Value: 12201120, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2], Next: nil},
		&Node{Value: 12201121, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2], Next: nil},
		&Node{Value: 12201122, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2], Next: nil},
	}
	// Level 8
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next = []*Node{
		&Node{Value: 101211210, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1], Next: nil},
	}
	head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2].Next[0].Next = []*Node{
		&Node{Value: 122011200, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2].Next[0], Next: nil},
		&Node{Value: 122011201, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2].Next[0], Next: nil},
		&Node{Value: 122011202, Parent: head.Next[2].Next[2].Next[0].Next[1].Next[1].Next[2].Next[0], Next: nil},
	}
	// Level 9
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next = []*Node{
		&Node{Value: 1012112100, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0], Next: nil},
		&Node{Value: 1012112101, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0], Next: nil},
		&Node{Value: 1012112102, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0], Next: nil},
	}
	// Level 10
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next[1].Next = []*Node{
		&Node{Value: 10121121010, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next[1], Next: nil},
		&Node{Value: 10121121011, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next[1], Next: nil},
	}
	head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next[2].Next = []*Node{
		&Node{Value: 10121121020, Parent: head.Next[0].Next[1].Next[2].Next[1].Next[1].Next[2].Next[1].Next[0].Next[2], Next: nil},
	}

	return head
}
