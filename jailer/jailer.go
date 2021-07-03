package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: jailer {num}")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	cells := make([]bool, num) // zero is closed, 1 is opened

	now := time.Now()
	for i := 0; i < num; i++ {
		for j := i; j < num; j += (i + 1) {
			cells[j] = !cells[j]
		}
	}
	elasped := time.Now().Sub(now)

	result := 0
	for _, cell := range cells {
		if cell {
			result++
		}
	}

	fmt.Printf("Result: %v, elapsed time: %v\n", result, elasped)
}
