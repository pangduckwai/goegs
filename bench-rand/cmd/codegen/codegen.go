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
			generate(i + 2)
		}
	case 2:
		cmd, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		generate(cmd)
	default:
		common.Usage(true)
	}
}

func generate(num int) {
	var max uint64 = math.MaxUint64
	var stp uint64 = max / uint64(num)
	var lst = stp * uint64(num)
	var nxt uint64
	var dff = max - lst
	var idx int
	var fnxt, fstp float64

	if dff != 0 {
		fstp = float64(num) / float64(dff)
		fnxt = fstp
	}

	nxt = stp
	if int(dff) > num {
		panic("Not enough step to fill up the differences")
	} else if int(dff) == num { // start filling only if differences equal # of steps
		nxt += 1
		dff -= 1
	}
	fmt.Printf("\tcase %v:\n", num)
	fmt.Printf("\t\tif rnd < %-20v {        // diff:%v; %v÷%v=%v\n\t\t\treturn %v\n", nxt, dff, num, max-lst, fstp, idx)

	for idx = 1; idx < num-1; idx++ {
		nxt += stp
		if dff > 0 && (int(fnxt) == idx) {
			nxt += 1
			dff -= 1
			fnxt += fstp
		}
		fmt.Printf("\t\t} else if rnd < %-20v { // diff:%v\n\t\t\treturn %v\n", nxt, dff, idx)
	}

	nxt += stp
	if dff > 0 {
		nxt += 1
		dff -= 1
	}
	fmt.Printf("\t\t} else { //   < %-20v   // calc:%v\n\t\t\treturn %v\n\t\t} // num:%v\n", nxt, lst, idx, num)
	if nxt != max {
		panic("Not filled!!!")
	}
}
