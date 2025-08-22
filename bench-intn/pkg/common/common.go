package common

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const RUN_NUM = 10000000 // 10,000,000
const SEL_NUM = 4

const USAGE = "Usage: rand [1-3, 9-15]"

const TRUNC = time.Duration(time.Millisecond)

func Usage(fatal bool) {
	if fatal {
		log.Fatalf(USAGE)
	}
	log.Printf(USAGE)
}

func DisplayRaw(
	idx uint64,
	lpsd time.Duration,
) {
	fmt.Printf(" %2v in %-7v\n", idx, lpsd.Round(TRUNC))
}

func DisplayRun(
	idx, ran uint64,
	lpsd time.Duration,
	rsts []uint64,
	pad int,
) {
	frm := fmt.Sprintf(" | %%2v: %%%vv (%%7.4f%%%%)", pad)

	var sbr strings.Builder
	for j, rst := range rsts {
		ptg := (float64(rst) / float64(ran)) * 100.0
		sbr.WriteString(fmt.Sprintf(frm, j, rst, ptg))
	}

	fmt.Printf(" %2v in %-7v%v\n", idx, lpsd.Round(TRUNC), sbr.String())
}

func DisplayAvg(
	ttl time.Duration,
	iter int,
	ran uint64,
) {
	avg := ttl / time.Duration(iter)
	idv := avg / time.Duration(ran)
	fmt.Printf("   avg %-7v (%v)\n\n", avg.Round(TRUNC), idv)
}
