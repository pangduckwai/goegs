package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"sea9.org/go/egs/rand/old"
	v2 "sea9.org/go/egs/rand/v2"
)

func main() {
	fmt.Println("randBench - rand benchmark")

	if len(os.Args) != 2 {
		log.Fatalf("Usage: rand [1-7, 9-15]")
	}
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if c < 1 || c > 15 || c == 8 {
		log.Fatalf("Usage: rand [1-7, 9-15]")
	}

	if c < 8 {
		old.Run(uint8(c))
	} else {
		v2.Run(uint8(c & 7))
	}
}
