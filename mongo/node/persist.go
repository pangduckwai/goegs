package node

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/////////////////////////////////////////////////////////////////////////////
// To prepare the mongodb
// use wdom
//
// db.createCollection('mctree', {storageEngine: {wiredTiger: {configString: 'block_compressor=zlib'}}})
// db.mctree.createIndex({"f": 1})
// db.mctree.createIndex({"p": -1})
//
// db.mctree.drop()
// db.runCommand({dropIndexes: "mctree", index: "e_1"})
// db.runCommand({dropIndexes: "mctree", index: "p_1"})
/////////////////////////////////////////////////////////////////////////////

var isConn bool

// var mngo *mongo.Database
// var coll *mongo.Collection
var conn *mongo.Client
var cntx context.Context

// Stop clean up connection
var Stop func()

// Connect connect to a mongo database
func Connect(url string) error {
	ctx, cancel := context.WithCancel(context.Background())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		defer cancel()
		return err
	}

	cntx = ctx
	conn = client
	// mngo = client.Database(database)
	// coll = mngo.Collection(collection)

	Stop = func() {
		defer cancel()
		defer client.Disconnect(ctx)
		isConn = false
	}
	isConn = true
	return nil
}

// Connected check if mongo connection is ready
func Connected() bool {
	return isConn
}

// Collection get collection from the active mongodb connection
func Collection(database string, collection string) (*mongo.Collection, error) {
	if !isConn {
		return nil, Err(0)
	}

	mngo := conn.Database(database)
	return mngo.Collection(collection), nil
}

// NOTE!!!!! split is "destructive" because in order to properly save document to mongo, the tree
// is cut to pieces according to the tree level. After child nodes (sub-tree) of a node is cut (by adding ID
// to the parent, and Refn to the head of the sub-tree), the []Next need to be clear (note the line "node.Next = nil"
// near the end of the split() function) before sending the parent tree to mongo to avoid duplication. This need to
// be done in-place to avoid making duplicate copies of the tree in memory.
//
// IMPORTANT!!! so as a result, after splitting the tree to persist, further simulations cannot use the in memory
// tree, need to load again from mongo

const levels = 90 // HERE!!!

func split(tree *Node) ([]*Node, error) {
	if tree.Parent != nil {
		return nil, Err(1)
	}

	var vertices []*Node
	var node *Node
	var nxtLvl int
	var modlus int

	// allowed depth of the current sub-tree according to the level of the current node
	// e.g. at the top of the tree allows fewer levels (e.g for node level 3-10, sub-tree can have 5 levels max)
	// It is to keep the size grow of the top level sub-tree not the exceed mongo's limit (16MB ?)
	var currDepth int

	tree.Level = 0
	queue := []*Node{tree}
	for len(queue) > 0 {
		node = queue[0]

		switch {
		case node.Level <= 2:
			currDepth = 1
		case node.Level <= 10:
			currDepth = 4
		case node.Level <= 60:
			currDepth = 19
		default:
			// Mongo only support 180 levels nested document, since the Next array seems to count as another level,
			// max level a sub-tree can have is 90
			currDepth = 89
		}

		modlus = node.Level % (currDepth + 1)
		if modlus == currDepth && node.ID == primitive.NilObjectID {
			// reached the allowed depth of the sub-tree, give each node an ID so the root of the next sub-tree can ref
			node.ID = primitive.NewObjectID()
		} else if modlus == 0 {
			// split to new sub-tree
			if node.ID == primitive.NilObjectID {
				node.ID = primitive.NewObjectID()
			}
			vertices = append(vertices, node)
		}

		nxtLvl = node.Level + 1
		for _, n := range node.Next {
			n.Level = nxtLvl
			if modlus == currDepth {
				// reached the allowed depth of the sub-tree, mark ref of all next level nodes with ID of the current node
				n.Refn = node.ID
			}
		}

		queue = append(queue[1:], node.Next...)
		if modlus == currDepth {
			// finally cut the nodes below so that won't store the whole thing in mongo
			node.Next = nil
		}
	}

	return vertices, nil
}

// Write write to mongodb
func Write(coll *mongo.Collection, node *Node) ([]*Node, int64, int64, int64, int64, time.Duration, error) {
	if !isConn {
		return nil, 0, 0, 0, 0, 0, Err(0)
	}

	vertices, err := split(node)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, err
	}

	matchedCount, modifiedCount, upsertedCount, elapsed, err := WriteNodes(coll, vertices)
	if err != nil {
		return nil, 0, 0, 0, 0, 0, err
	}
	return vertices, int64(len(vertices)), matchedCount, modifiedCount, upsertedCount, elapsed, nil
}

// WriteNodes write given nodes
func WriteNodes(coll *mongo.Collection, vertices []*Node) (matchedCount, modifiedCount, upsertedCount int64, elapsed time.Duration, err error) {
	if !isConn {
		return 0, 0, 0, 0, Err(0)
	}

	now := time.Now()
	cnt := 0
	for i, vertex := range vertices {
		if vertex.ID == primitive.NilObjectID {
			return matchedCount, modifiedCount, upsertedCount, 0, Err(4)
		}
		rst, err := coll.UpdateOne(
			cntx,
			bson.M{"_id": vertex.ID},
			bson.M{"$set": vertex},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return matchedCount, modifiedCount, upsertedCount, 0, err
		}
		matchedCount += rst.MatchedCount
		modifiedCount += rst.ModifiedCount
		upsertedCount += rst.UpsertedCount
		if i%50000 == 49999 {
			fmt.Printf("[WDOM-MC][WRTN] Vertices %v: updated %v; added %v\n", i, modifiedCount, upsertedCount)
		}
		cnt++
	}
	elapsed = time.Now().Sub(now)
	// fmt.Printf("[WDOM-MC][<<<<] Vertices %v: matched %v; updated %v; added %v; elapsed time %v\n", cnt, matchedCount, modifiedCount, upsertedCount, elapsed)

	return matchedCount, modifiedCount, upsertedCount, elapsed, nil
}

// Root read tree from mongo
func Root(coll *mongo.Collection, impl string, createIfNotFound bool) (*Node, time.Duration, error) {
	if !isConn {
		return nil, 0, Err(0)
	}
	if impl == "" {
		return nil, 0, Err(2)
	}
	now := time.Now()
	var read Node
	if err := coll.FindOne(cntx, bson.M{"p": impl}).Decode(&read); err != nil {
		if err == mongo.ErrNoDocuments && createIfNotFound {
			node := New(nil, impl, 0, 0, nil, 0)
			node.Runs = 1
			return node, time.Now().Sub(now), nil
		}
		return nil, 0, err
	}
	return &read, time.Now().Sub(now), nil
}

// Read read node by reference
func Read(coll *mongo.Collection, oid primitive.ObjectID) (*Node, time.Duration, error) {
	if !isConn {
		return nil, 0, Err(0)
	}
	if oid == primitive.NilObjectID {
		return nil, 0, Err(4)
	}
	now := time.Now()
	var read Node
	if err := coll.FindOne(cntx, bson.M{"_id": oid}).Decode(&read); err != nil {
		return nil, 0, err
	}
	return &read, time.Now().Sub(now), nil
}

// Next read next level tree but return the parent
func Next(coll *mongo.Collection, node *Node) (*Node, int, time.Duration, error) {
	if !isConn {
		return nil, 0, 0, Err(0)
	}
	if node == nil {
		return nil, 0, 0, Err(3)
	}
	if node.ID == primitive.NilObjectID {
		return nil, 0, 0, Err(4)
	}

	if len(node.Next) <= 0 { // if node.Next is not empty, don't need to read from mongo, just return the given node
		now := time.Now()
		cursor, err := coll.Find(cntx, bson.M{"f": node.ID})
		if err != nil {
			return nil, 0, 0, err
		}
		defer cursor.Close(cntx)

		var result []*Node
		for cursor.Next(cntx) {
			var read Node
			if err = cursor.Decode(&read); err != nil {
				return nil, 0, 0, err
			}
			result = append(result, &read)
		}

		node.Next = result
		return node, len(result), time.Now().Sub(now), nil
	}
	return node, len(node.Next), 0, nil
}

// Walk walk the entire tree stored in mongo. Note this won't build the entire tree,
// instead return all vertices (a vertex is a sub-tree stored as a separate document in mongo)
func Walk(coll *mongo.Collection, tree *Node) ([]*Node, error) {
	now := time.Now()
	var node *Node
	queue := []*Node{tree}
	vertices := []*Node{tree}
	ncnt := 0
	rcnt := 0
	vcnt := 0
	for len(queue) > 0 {
		node = queue[0]
		ncnt++

		if node.Next == nil && node.ID != primitive.NilObjectID {
			node, _, _, err := Next(coll, node)
			if err != nil {
				return nil, err
			}
			if len(node.Next) > 0 {
				rcnt++
				vertices = append(vertices, node.Next...)
				vcnt += len(node.Next)
			}
			if ncnt%100000 == 0 {
				fmt.Printf("[WDOM-MC][WLK] Traversing: nodes:%9v (%v)\n", ncnt, node)
			}
		}
		queue = append(queue[1:], node.Next...)
	}

	fmt.Printf("[WDOM-MC][WLK] Traversed nodes:%9v (vertices:%7v (%7v), read:%7v) in %v\n", ncnt, vcnt, len(vertices), rcnt, time.Now().Sub(now))
	return vertices, nil
}

// Err persist errors
type Err uint8

func (e Err) Error() string {
	switch e {
	case 0:
		return "Not conncted"
	case 1:
		return "Node provided is not the root node"
	case 2:
		return "Missing game implementation ID"
	case 3:
		return "Missing parent node"
	case 4:
		return "Missing node ID"
	case 5:
		return "Error choosing next move while tracking"
	case 8:
		return "UCB failed"
	case 9:
		return "Back-propagation failed, root not reached"
	default:
		return "Unknown error code " + strconv.Itoa(int(e))
	}
}
