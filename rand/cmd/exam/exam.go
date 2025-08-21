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
	var err error
	var cmd int

	switch len(os.Args) {
	case 1:
		for i := range 42 {
			exam(i + 2)
			fmt.Println()
		}
	case 2:
		cmd, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		exam(cmd)
	default:
		common.Usage(true)
	}
}

func exam(cnt int) {
	var max uint64 = math.MaxUint64
	var stp uint64 = max / uint64(cnt)
	var lst = stp * uint64(cnt)
	var nxt uint64
	var dff = max - lst
	var spd, idx int

	if dff != 0 {
		spd = int(math.Round(float64(cnt) / float64(dff)))
	}

	nxt = stp
	if int(dff) > cnt {
		panic("Not enough step to fill up the differences")
	} else if int(dff)*spd >= cnt { // start filling only if differences equal # of steps
		nxt += 1
		dff -= 1
	}
	fmt.Printf("  0  if rnd < %v { // diff: %v (%v); spread: %v\n", nxt, dff, max-lst, spd)

	for idx = 1; idx < cnt-1; idx++ {
		nxt += stp
		if dff > 0 && (spd == 1 || idx%spd > 0) {
			nxt += 1
			dff -= 1
		}
		fmt.Printf("%3v  } else if rnd < %v { // diff: %v\n", idx, nxt, dff)
	}

	nxt += stp
	if dff > 0 {
		nxt += 1
		dff -= 1
	}
	fmt.Printf("%3v  } else { // %v (%20v)\n%3v  }\n", idx, nxt, lst, cnt)
	if nxt != max {
		panic("Not filled!!!")
	}
}
