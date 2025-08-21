package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"sea9.org/go/egs/rand/pkg/common"
)

func main() {
	if len(os.Args) != 2 {
		common.Usage(true)
	}
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var max uint64 = math.MaxUint64
	var stp uint64 = max / uint64(c)
	var nxt = stp * 2
	var lst = stp * uint64(c)

	fmt.Printf("  0 - %20v\n", stp)
	for i := 1; i < c; i++ {
		fmt.Printf("%3v - %20v\n", i, nxt)
		nxt += stp
	}
	fmt.Printf("\n  max %20v*\n", max)
	fmt.Printf(" last %20v\n", lst)
	fmt.Printf("    = %20v\n", max-lst+1)
}
