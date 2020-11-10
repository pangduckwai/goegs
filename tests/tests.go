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
	tree := mongo.Build()
	err := mongo.Write(tree)
	// list, err := mongo.Split(tree, 3)

	// tree, err := mongo.Read()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mongo.Change(tree, 3)
	mongo.Fill(tree)

	err = mongo.Write(tree)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tree.Tree(10))
	// fmt.Println("Is the same:", mongo.Same(tree, mongo.Build()))

	// for _, frag := range list {
	// 	fmt.Println(frag.Tree(3))
	// }
}
