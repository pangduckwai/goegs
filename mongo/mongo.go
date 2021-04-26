package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"sea9.org/go/mongo-test/node"
)

const impl = "BSCTTTCCDMT5XT0MTD-6"
const monurl = "mongodb://192.168.56.31:27017"
const dbsrc = "test1"
const dbtgt = "test"
const colsrc = "mct"
const coltgt = "mct"

func walk(database string, collection string) {
	fmt.Printf("[MONGO-TEST] %v.%v\n", database, collection)
	coll, err := node.Collection(database, collection)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}

	// Read existing nodes
	head, elapsed, err := node.Root(coll, impl, false)
	if err != nil {
		log.Fatalf("[Error] %v\n", err)
	}
	fmt.Println("[MONGO-TEST] node read:", elapsed)

	// Walk thru the entire tree
	var curr *node.Node
	queue := []*node.Node{head}
	for len(queue) > 0 {
		curr = queue[0]

		if curr.Next == nil && curr.ID != primitive.NilObjectID {
			_, _, _, err := node.Next(coll, curr)
			if err != nil {
				log.Fatalf("[Error] %v\n", err)
			}
		}

		queue = append(queue[1:], curr.Next...)
	}

	if valid, count := node.Validate(head); valid {
		fmt.Printf("[MONGO-TEST] Validated %v nodes: okay\n", count)
		fmt.Println(head.Tree(10, 0))
	} else {
		fmt.Println(head.Tree(10, 0))
		log.Fatal("[Error] Validation failed")
	}
}

// func merge() {
// 	collSrc, err := node.Collection(dbsrc, colsrc)
// 	if err != nil {
// 		log.Fatalf("[Error] %v\n", err)
// 	}
// 	headSrc, _, err := node.Root(collSrc, impl, false)
// 	if err != nil {
// 		log.Fatalf("[Error] %v\n", err)
// 	}
// 	var currSrc *node.Node
// 	queSrc := []*node.Node{headSrc}

// 	collTgt, err := node.Collection(dbtgt, coltgt)
// 	if err != nil {
// 		log.Fatalf("[Error] %v\n", err)
// 	}
// 	headTgt, _, err := node.Root(collTgt, impl, false)
// 	var currTgt *node.Node
// 	queTgt := []*node.Node{headTgt}

// 	for len(queSrc) > 0 && len(queTgt) > 0 {
// 		currSrc = queSrc[0]
// 		currTgt = queTgt[0]

// 		fmt.Printf("Source: %v\n", currSrc)
// 		fmt.Printf("Target: %v\n", currTgt)

// 		queSrc = append(queSrc[1:], cu)
// 	}
// }

func main() {
	fmt.Println("[MONGO-TEST]######################################################")
	fmt.Printf("[MONGO-TEST] Connecting to %v %v.%v -> %v.%v\n", monurl, dbsrc, colsrc, dbtgt, coltgt)

	// Connect to mongo
	err := node.Connect(monurl)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Stop()

	if !node.Connected() {
		log.Fatal("[Error] Connection failed")
	}

	walk(dbtgt, coltgt)
	walk(dbsrc, colsrc)
}
