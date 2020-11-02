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
	Round  uint32 // (pkey) indicate progress of the game
	Turn   uint8  // (pkey) identify the player
	Action uint8  // (pkey) Actions a player may takes
	Index1 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Index2 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Index3 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Value1 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Value2 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Value3 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Runs   uint64 // Total number of runs
	Wins   uint64 // Total number of won games of this player
	Parent *Node
	Next   []*Node
}

// MNode node in mongodb
type MNode struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Parent primitive.ObjectID `bson:"parent,omitempty"`
	Impl   string             `bson:"impl,omitempty"`
	Round  uint32             `bson:"round"`
	Turn   uint8              `bson:"turn"`
	Action uint8              `bson:"action"`
	Index1 uint8              `bson:"index1,omitempty"`
	Index2 uint8              `bson:"index2,omitempty"`
	Index3 uint8              `bson:"index3,omitempty"`
	Value1 uint32             `bson:"value1,omitempty"`
	Value2 uint32             `bson:"value2,omitempty"`
	Value3 uint32             `bson:"value3,omitempty"`
	Runs   uint64             `bson:"runs"`
	Wins   uint64             `bson:"wins"`
}

const url = "mongodb://192.168.56.101:27017"

// const url = "mongodb+srv://wdom:oY7v4xeqxO8zDFjz@m0-sea9.7zfms.mongodb.net/mctree?retryWrites=true&w=majority"

// Root read the root node entry from mongodb
func Root(impl string) (*MNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	var root MNode
	if err := tree.FindOne(ctx, bson.M{"impl": impl}).Decode(&root); err != nil {
		if err == mongo.ErrNoDocuments { // Root node not found, create one
			root = MNode{
				Impl:   impl,
				Round:  0,
				Turn:   0,
				Action: 0,
				Runs:   0,
				Wins:   0,
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
// func Find(leaf *Node) (*MNode, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer client.Disconnect(ctx)

// 	db := client.Database("wdom")
// 	tree := db.Collection("mctree")

// 	var mnode MNode
// 	opts := options.FindOne().SetSort(bson.D{{Key: "round", Value: 1}})
// 	filter := bson.M{
// 		"round":  leaf.Round,
// 		"turn":   leaf.Turn,
// 		"action": leaf.Action,
// 	}
// 	idx := join(leaf.Indices)
// 	val := join(leaf.Values)
// 	if len(idx) > 0 {
// 		filter["indices"] = idx
// 	}
// 	if len(val) > 0 {
// 		filter["values"] = val
// 	}

// 	if err := tree.FindOne(ctx, filter, opts).Decode(&mnode); err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &mnode, nil
// }

// Add add a leaf node
func Add(parent *MNode, leaf *Node) (*MNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)

	db := client.Database("wdom")
	tree := db.Collection("mctree")

	// mgnode := bson.M{
	// 	"$set": bson.M{
	// 		"round":  leaf.Round,
	// 		"turn":   leaf.Turn,
	// 		"action": leaf.Action,
	// 		"index1": leaf.Index1,
	// 	},
	// 	"$inc": bson.M{"runs": 1, "wins": 1},
	// }
	mgnode := bson.M{
		"round":  leaf.Round,
		"turn":   leaf.Turn,
		"action": leaf.Action,
		"index1": leaf.Index1,
	}

	incrmt := bson.M{"runs": 1, "wins": 1}

	filter := bson.M{}
	for k, v := range mgnode {
		filter[k] = v
	}

	mgnode["parent"] = parent.ID

	rslt, err := tree.UpdateOne(ctx, filter, bson.M{
		"$set": mgnode,
		"$inc": incrmt,
	}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	if rslt.UpsertedID != nil {
		fmt.Println("New node ID: ", rslt.UpsertedID, " Match count: ", rslt.MatchedCount, " Modify count: ", rslt.ModifiedCount, " Insert count: ", rslt.UpsertedCount) // TODO TEMP
		return &MNode{
			ID:     rslt.UpsertedID.(primitive.ObjectID),
			Parent: parent.ID,
			Round:  leaf.Round,
			Turn:   leaf.Turn,
			Action: leaf.Action,
			Index1: leaf.Index1,
		}, nil
	}

	fmt.Println("Existing node - Match count: ", rslt.MatchedCount, " Modify count: ", rslt.ModifiedCount, " Insert count: ", rslt.UpsertedCount) // TODO TEMP
	return &MNode{
		Parent: parent.ID,
		Round:  leaf.Round,
		Turn:   leaf.Turn,
		Action: leaf.Action,
		Index1: leaf.Index1,
	}, nil
}

// Next search for nodes in the next level of the tree
func Next(node *MNode) ([]MNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
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

	result := make([]MNode, 0)
	for cursor.Next(ctx) {
		var next MNode
		if err = cursor.Decode(&next); err != nil {
			return nil, err
		}
		result = append(result, next)
	}
	return result, nil
}

// func join(slice []uint8) string {
// 	rst := ""
// 	for _, v := range slice {
// 		rst += fmt.Sprintf(",%v", v)
// 	}
// 	if len(rst) > 0 {
// 		return rst[1:]
// 	}
// 	return rst
// }
