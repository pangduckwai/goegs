package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// fibonacci returns a function that returns the next term in the Fibonacci sequence, without overflowing, for each call.
// i0 - the first starting term.
// i1 - the second starting term.
func fibonacci(i0, i1 uint64) func() (int, uint64, error) {
	var max uint64 = math.MaxUint64
	var t int
	var p, f uint64 = i0, i1 // Item i minus 2 and i minus 1

	return func() (int, uint64, error) {
		t++

		if p > (max - f) {
			return t, f, fmt.Errorf("Calculating the %v-th term of the Fibonacci sequence will results in overflow, maximum 64-bits integer is %v.", t, max)
		}

		if t == 1 {
			return t, p, nil
		} else if t == 2 {
			return t, f, nil
		}

		p, f = f, p+f // Note: can't do that without multi assignment, as 'p' is overwritten by 'f', but 'f' need 'p' to calculate...
		return t, f, nil
	}
}

func main() {
	var err error
	var cnt int
	var i0, i1 uint64 = 0, 1

	switch len(os.Args) {
	case 4:
		i1, err = strconv.ParseUint(os.Args[3], 10, 64)
		if err != nil {
			log.Fatalf("Usage: fibonacci [num of terms] {1st term} {2nd term}, %v", err)
		}
		fallthrough
	case 3:
		i0, err = strconv.ParseUint(os.Args[2], 10, 64)
		if err != nil {
			log.Fatalf("Usage: fibonacci [num of terms] {1st term} {2nd term}, %v", err)
		}
		fallthrough
	case 2:
		cnt, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Usage: fibonacci [num of terms] {1st term} {2nd term}, %v", err)
		}
	default:
		log.Fatal("Usage: fibonacci [num of terms] {1st term} {2nd term}")
	}

	fn := fibonacci(i0, i1)

	now := time.Now()
	fmt.Println("  #   value")
	fmt.Println("==== =======")
	var t int
	var f uint64
	for range cnt {
		t, f, err = fn()
		if err != nil {
			fmt.Printf("%3v   %v\n", t, err)
			t--
			break
		}
		fmt.Printf("%3v   %v\n", t, f)
	}
	fmt.Printf("Calculated %v terms in the Fibonacci sequence, elapsed time %v\n", t, time.Now().Sub(now))
}

// Simple answer
// func main() {
// 	p := 0
// 	f := 1
// 	for i := 3; i <= 30; i++ {
// 		t := f
// 		f = f + p
// 		p = t
// 		fmt.Printf("%2v - %v\n", i, f)
// 	}
// }
