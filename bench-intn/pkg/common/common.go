package common

import (
	"fmt"
	"log"
	"time"
)

const RUN_NUM = 10000000 // 10,000,000

const USAGE = "Usage: rand [1-3, 9-15]"

const TRUNC = time.Duration(time.Millisecond)

func Usage(fatal bool) {
	if fatal {
		log.Fatalf(USAGE)
	}
	log.Printf(USAGE)
}

func DisplayRun(
	idx, ran uint64,
	lpsd time.Duration,
	c0, c1, c2, c3, c4, c5, c6 uint64,
	pad int,
) {
	p0 := (float64(c0) / float64(ran)) * 100.0
	p1 := (float64(c1) / float64(ran)) * 100.0
	p2 := (float64(c2) / float64(ran)) * 100.0
	p3 := (float64(c3) / float64(ran)) * 100.0
	p4 := (float64(c4) / float64(ran)) * 100.0
	p5 := (float64(c5) / float64(ran)) * 100.0
	p6 := (float64(c6) / float64(ran)) * 100.0

	frm := fmt.Sprintf(" %%2v in %%-7v 0: %%%vv (%%7.4f%%%%) | 1: %%%vv (%%7.4f%%%%) | 2: %%%vv (%%7.4f%%%%) | 3: %%%vv (%%7.4f%%%%) | 4: %%%vv (%%7.4f%%%%) | 5: %%%vv (%%7.4f%%%%) | 6: %%%vv (%%7.4f%%%%) \n", pad, pad, pad, pad, pad, pad, pad)

	fmt.Printf(frm, idx, lpsd.Round(TRUNC), c0, p0, c1, p1, c2, p2, c3, p3, c4, p4, c5, p5, c6, p6)
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
