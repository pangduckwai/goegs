package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Err node errors
type Err struct {
	Msg string
}

func (e *Err) Error() string {
	return fmt.Sprintf("[node]%v", e.Msg)
}

const ctxTimeout = 10
const ctxTimeoutLong = 300

var isConn bool
var conn *mongo.Client

// Stop clean up connection
func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()
	isConn = false
	conn.Disconnect(ctx)
}

// Connect connect to a mongo database
func Connect(url string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	conn, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		defer cancel()
		return err
	}

	isConn = true
	return nil
}

// Connected check if mongo connection is ready
func Connected() bool {
	if !isConn {
		return false
	}
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()
	return conn.Ping(ctx, nil) == nil
}

// Collection get collection from the active mongodb connection
func Collection(database string, collection string) *mongo.Collection {
	if !isConn {
		return nil
	}

	mngo := conn.Database(database)
	return mngo.Collection(collection)
}

// Read read node by reference
func Read(db string, cl string, oid primitive.ObjectID) (obj *Objs, err error) {
	if !isConn {
		err = &Err{"Not connected"}
		return
	}
	coll := Collection(db, cl)
	if coll == nil {
		err = &Err{"Error getting collection"}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer cancel()

	if oid == primitive.NilObjectID {
		err = &Err{"Missing object ID"}
		return
	}

	var read Objs
	if err = coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&read); err != nil {
		return
	}

	obj = &read
	return
}

func Write(db string, cl string, obj *Objs) (matchedCount, modifiedCount, upsertedCount int64, err error) {
	if !isConn {
		err = &Err{"Not connected"}
		return
	}
	coll := Collection(db, cl)
	if coll == nil {
		err = &Err{"Error getting collection"}
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeoutLong*time.Second)
	defer cancel()

	if obj.ID == primitive.NilObjectID {
		err = &Err{"Missing object ID"}
		return
	}

	rst, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": obj.ID},
		bson.M{"$set": obj},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return
	}

	matchedCount = rst.MatchedCount
	modifiedCount = rst.ModifiedCount
	upsertedCount = rst.UpsertedCount
	return
}

type Objs struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Str   string             `bson:"s,omitempty"`
	Val1  uint64             `bson:"x,omitempty"`
	Val2  uint32             `bson:"y,omitempty"`
	Val3  uint16             `bson:"z,omitempty"`
	Level int                `bson:"-"`
}

const url = "mongodb://192.168.56.31:27017"
const db = "test"
const col = "obj"

func main() {
	err := Connect(url)
	if err != nil {
		log.Fatal(err)
	}
	defer Stop()

	if !Connected() {
		log.Fatal("[WDOM-MC][CONN][Error] Connection failed")
	}

	var matched, modified, upserted int64

	obj1 := &Objs{
		ID:    primitive.NewObjectID(),
		Str:   "UVWXYZA",
		Val1:  0x4000000000000000,
		Val2:  0x1000,
		Level: 2345234,
	}
	matched, modified, upserted, err = Write(db, col, obj1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Obj1: Matched:%v Modified:%v Upserted:%v\n", matched, modified, upserted)

	obj2 := &Objs{
		ID:    primitive.NewObjectID(),
		Str:   "UVWXYZB",
		Val2:  0x20000000,
		Val3:  0x1000,
		Level: 2345234,
	}
	matched, modified, upserted, err = Write(db, col, obj2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Obj2: Matched:%v Modified:%v Upserted:%v\n", matched, modified, upserted)
}
