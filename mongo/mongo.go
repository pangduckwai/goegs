package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"sea9.org/go/mongo-test/node"
)

const impl = "BSCTTTCCDMT5XT0MTD-6"
const monurl = "mongodb://192.168.56.31:27017"
const mondb = "test"
const moncoll = "mct"

func main() {
	fmt.Println("[MONGO-TEST]######################################################")
	fmt.Printf("[MONGO-TEST] Connecting to %v %v.%v\n", monurl, mondb, moncoll)

	// Connect to mongo
	err := node.Connect(monurl, mondb, moncoll)
	if err != nil {
		log.Fatal(err)
	}
	defer node.Stop()

	if !node.Connected() {
		log.Fatal("[Error] Connection failed")
	}

	// Read existing nodes
	head, elapsed, err := node.Root(impl, true)
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
			_, _, _, err := node.Next(curr)
			if err != nil {
				log.Fatalf("[Error] %v\n", err)
			}
		}

		queue = append(queue[1:], curr.Next...)
	}

	if valid, count := node.Validate(head); valid {
		fmt.Printf("[MONGO-TEST] Validated %v nodes: okay\n", count)
		fmt.Println(head.Tree(50, 0))
	} else {
		fmt.Println(head.Tree(50, 0))
		log.Fatal("[Error] Validation failed")
	}
}
