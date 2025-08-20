package old

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Run(c uint8) {
	if c > 3 {
		fmt.Printf("Input range: 1, 2, 3")
		return
	}

	var run uint64 = 10000000
	frm := fmt.Sprintf("%%%vv", int(math.Log10(float64(run))))
	var ttl time.Duration

	runs := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	switch {
	case c&1 > 0:
		for _, idx := range runs {
			r0, r1, r2, r3, elapsed := sim0(rand.New(rand.NewSource(time.Now().UnixNano()+idx)), run)
			ttl += elapsed
			fmt.Printf("randBench (old 0) %2v in %-14v - 0: "+frm+"; 1: "+frm+"; 2: "+frm+"; 3: "+frm+"\n", idx, elapsed, r0, r1, r2, r3)
		}
	case c&2 > 0:
		for _, idx := range runs {
			r0, r1, r2, r3, elapsed := sim1(rand.New(rand.NewSource(time.Now().UnixNano()+idx)), run)
			ttl += elapsed
			fmt.Printf("randBench (old 1) %2v in %-14v - 0: "+frm+"; 1: "+frm+"; 2: "+frm+"; 3: "+frm+"\n", idx, elapsed, r0, r1, r2, r3)
		}
	}
	fmt.Printf("                    avg %-14v\n\n", ttl/time.Duration(len(runs)))
}

func sim0(
	rnd *rand.Rand,
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	strt := time.Now()
	for range run {
		switch rnd.Intn(4) {
		case 0:
			c0++
		case 1:
			c1++
		case 2:
			c2++
		case 3:
			c3++
		}
	}
	lpsd = time.Since(strt)
	return
}

func sim1(
	rnd *rand.Rand,
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	var val int64
	strt := time.Now()
	for range run {
		val = rnd.Int63()
		if val < 2305843009213693951 {
			c0++
		} else if val < 4611686018427387902 {
			c1++
		} else if val < 6917529027641081853 {
			c2++
		} else {
			c3++
		}
	}
	lpsd = time.Since(strt)
	return
}
