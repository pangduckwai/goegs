package main

import (
	"fmt"
	"log"
	"os"

	"sea9.org/go/tests/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hellos(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
