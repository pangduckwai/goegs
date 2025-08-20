package main

import (
	"fmt"

	"sea9.org/go/egs/rand/old"
	v2 "sea9.org/go/egs/rand/v2"
)

func main() {
	fmt.Println("randBench - rand benchmark")
	v2.Run()
	fmt.Println()
	old.Run()
}
