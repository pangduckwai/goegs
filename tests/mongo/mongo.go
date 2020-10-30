package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Node in memory node
type Node struct {
	Round   uint32  // (pkey) indicate progress of the game
	Turn    uint8   // (pkey) identify the player
	Action  uint8   // (pkey) Actions a player may takes
	Indices []uint8 // (pkey) indices of players, territories or cards (Depends on the event)
	Values  []uint8 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Runs    uint64  // Total number of runs
	Wins    uint64  // Total number of won games of this player
	Parent  *Node
	Next    []*Node
}

// Mnode node in mongodb
type Mnode struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Parent  primitive.ObjectID `bson:"parent,omitempty"`
	Impl    string             `bson:"impl,omitempty"`
	Round   uint32             `bson:"round"`             // (pkey) indicate progress of the game
	Turn    uint8              `bson:"turn"`              // (pkey) identify the player
	Action  uint8              `bson:"action"`            // (pkey) Actions a player may takes
	Indices string             `bson:"indices,omitempty"` // (pkey) indices of players, territories or cards (Depends on the event)
	Values  string             `bson:"values,omitempty"`  // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Runs    uint64             `bson:"runs"`              // Total number of runs
	Wins    uint64             `bson:"wins"`              // Total number of won games of this player
}

// Root read the root node entry from mongodb
func Root(impl string) (*Mnode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	var root Mnode
	if err := tree.FindOne(ctx, bson.M{"impl": impl}).Decode(&root); err != nil {
		if err == mongo.ErrNoDocuments { // Root node not found, create one
			root = Mnode{
				Impl:   impl,
				Round:  0,
				Turn:   0,
				Action: 0,
			}
			if root.ID == primitive.NilObjectID {
				fmt.Println("Node is new", root.ID)
			} else {
				fmt.Println("Node exists", root.ID)
			}

			rslt, err := tree.InsertOne(ctx, root)
			if err != nil {
				return nil, err
			}
			fmt.Println("New root ID: ", rslt.InsertedID) // TODO TEMP
			root.ID, _ = rslt.InsertedID.(primitive.ObjectID)
			return &root, nil
		}
		return nil, err
	}
	return &root, nil
}

// Find find a node
func Find(leaf *Node) (*Mnode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	var mnode Mnode
	opts := options.FindOne().SetSort(bson.D{{Key: "round", Value: 1}})
	filter := bson.M{
		"round":  leaf.Round,
		"turn":   leaf.Turn,
		"action": leaf.Action,
	}
	idx := join(leaf.Indices)
	val := join(leaf.Values)
	if len(idx) > 0 {
		filter["indices"] = idx
	}
	if len(val) > 0 {
		filter["values"] = val
	}

	if err := tree.FindOne(ctx, filter, opts).Decode(&mnode); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &mnode, nil
}

// Add add a leaf node
func Add(parent *Mnode, leaf *Node) (*Mnode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	filter := bson.M{
		"round":  leaf.Round,
		"turn":   leaf.Turn,
		"action": leaf.Action,
	}
	idx := join(leaf.Indices)
	val := join(leaf.Values)
	if len(idx) > 0 {
		filter["indices"] = idx
	}
	if len(val) > 0 {
		filter["values"] = val
	}
	mnode := Mnode{
		Parent:  parent.ID,
		Round:   leaf.Round,
		Turn:    leaf.Turn,
		Action:  leaf.Action,
		Indices: join(leaf.Indices),
		Values:  join(leaf.Values),
		Runs:    1, // leaf.Runs,
		Wins:    0, // leaf.Wins,
	}
	rslt, err := tree.UpdateOne(ctx, filter, bson.M{
		"$set": mnode,
	}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	if rslt.UpsertedID != nil {
		fmt.Println("New node ID: ", rslt.UpsertedID, " Match count: ", rslt.MatchedCount, " Modify count: ", rslt.ModifiedCount) // TODO TEMP
		mnode.ID, _ = rslt.UpsertedID.(primitive.ObjectID)
	} else {
		fmt.Println("Existing node - Match count: ", rslt.MatchedCount, " Modify count: ", rslt.ModifiedCount) // TODO TEMP
	}
	return &mnode, nil
}

// Next search for nodes in the next level of the tree
func Next(node *Mnode) ([]Mnode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority",
	))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	cursor, err := tree.Find(ctx, bson.M{"parent": node.ID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]Mnode, 0)
	for cursor.Next(ctx) {
		var next Mnode
		if err = cursor.Decode(&next); err != nil {
			return nil, err
		}
		result = append(result, next)
	}
	return result, nil
}

func join(slice []uint8) string {
	rst := ""
	for _, v := range slice {
		rst += fmt.Sprintf(",%v", v)
	}
	if len(rst) > 0 {
		return rst[1:]
	}
	return rst
}
