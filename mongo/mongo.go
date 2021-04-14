package main

import (
	"fmt"
	"log"

	"sea9.org/go/mongo-test/node"
)

const impl = "MONGOTEST"
const monurl = "mongodb://192.168.56.31:27017"
const mondb = "test"
const moncoll = "mct"

func buildNodes(head *node.Node) {
	var next *node.Node
	curr := head
	curr.Moves = 42
	curr.Runs = 11
	curr.Wins = 1

	// 1st level
	next = node.New(curr, "", 0, 2, []uint8{10}, 0)
	next.Moves = 41
	next.Runs = 2
	curr.Next = append(curr.Next, next)

	next = node.New(curr, "", 0, 2, []uint8{17}, 0)
	next.Moves = 41
	next.Runs = 2
	next.Wins = 1
	curr.Next = append(curr.Next, next)

	next = node.New(curr, "", 0, 2, []uint8{34}, 0)
	next.Moves = 41
	next.Runs = 2
	curr.Next = append(curr.Next, next)

	next = node.New(curr, "", 0, 2, []uint8{39}, 0)
	next.Moves = 41
	next.Runs = 2
	curr.Next = append(curr.Next, next)

	next = node.New(curr, "", 0, 2, []uint8{7}, 0)
	next.Moves = 41
	next.Runs = 2
	curr.Next = append(curr.Next, next)

	// 2nd level
	curr = head.Next[0]
	next = node.New(curr, "", 1, 2, []uint8{20}, 0)
	next.Moves = 40
	next.Runs = 1
	curr.Next = append(curr.Next, next)

	curr = head.Next[1]
	next = node.New(curr, "", 1, 2, []uint8{5}, 0)
	next.Moves = 40
	next.Runs = 1
	next.Wins = 1
	curr.Next = append(curr.Next, next)

	curr = head.Next[2]
	next = node.New(curr, "", 1, 2, []uint8{4}, 0)
	next.Moves = 40
	next.Runs = 1
	curr.Next = append(curr.Next, next)

	curr = head.Next[3]
	next = node.New(curr, "", 1, 2, []uint8{29}, 0)
	next.Moves = 40
	next.Runs = 1
	curr.Next = append(curr.Next, next)

	curr = head.Next[4]
	next = node.New(curr, "", 1, 2, []uint8{21}, 0)
	next.Moves = 40
	next.Runs = 1
	curr.Next = append(curr.Next, next)
}

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
	buildNodes(head)

	if valid, count := node.Validate(head); valid {
		fmt.Printf("[MONGO-TEST] Validated %v nodes: okay\n", count)
		fmt.Println(head.Tree(50, 0))
	} else {
		fmt.Println(head.Tree(50, 0))
		log.Fatal("[Error] Validation failed")
	}
}
