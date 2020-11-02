package main

import (
	"fmt"
	"log"

	"sea9.org/go/tests/mongo"
)

func main() {
	//////////////
	// Greetings
	//////////////
	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)

	// message, err := greetings.Hellos(os.Args[1:])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	//////////
	// Mongo
	//////////
	root, err := mongo.Root("BSCTTTCCDMT5XT0MTD-6")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Root: ", *root)

	// node, err := mongo.Find(&mongo.Node{
	node, err := mongo.Add(root, &mongo.Node{
		Round:  1,
		Turn:   0,
		Action: 2,
		Index1: 21,
		Runs:   0,
		Wins:   0,
	})
	if err != nil {
		log.Fatal(err)
	}
	if node == nil {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Node:", node)
	}

	next, err := mongo.Next(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Next: ", next)
}
