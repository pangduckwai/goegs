package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// bench - Benchmark solution
func bench(n int) (
	result int,
	elapsed time.Duration,
) {
	now := time.Now()
	result = int(math.Sqrt(float64(n)))
	elapsed = time.Now().Sub(now)
	return
}

// sim - Direct simulation.
func sim(n int) (
	result int,
	elapsed time.Duration,
) {
	now := time.Now()
	cells := make([]bool, n) // zero is closed, 1 is opened
	for i := 0; i < n; i++ {
		for j := i; j < n; j += (i + 1) {
			cells[j] = !cells[j]
		}
	}
	for _, cell := range cells {
		if cell {
			result++
		}
	}
	elapsed = time.Now().Sub(now)
	return
}

func run(n int) (err error) {
	rs, es := sim(n)
	rb, eb := bench(n)
	if rs != rb {
		err = fmt.Errorf("Invalid simulation!, expected result is %v, got %v", rb, rs)
	}
	fmt.Printf("%4v - Sim: %2v %-7v / Bench: %2v %-7v\n", n, rs, es, rb, eb)
	return
}

func main() {
	var err error
	switch len(os.Args) {
	case 1:
		for n := range 100 {
			err = run(n + 1)
		}
	case 2:
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Usage: jailer [num], %v", err)
		}
		err = run(n)
	default:
		log.Fatal("Usage: jailer [num]")
	}
	if err != nil {
		log.Fatalf("Usage: jailer [num], %v", err)
	}
}
