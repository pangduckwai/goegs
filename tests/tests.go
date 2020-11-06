package main

import (
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
	if err != nil {
		log.Fatal(err)
	}
	// for _, frag := range list {
	// 	fmt.Println(frag.Tree(3))
	// }
}
