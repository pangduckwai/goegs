package v2

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func Run() {
	var run uint64 = 10000000
	frm := fmt.Sprintf("%%%vv", int(math.Log10(float64(run))))

	var ttl time.Duration
	run0 := []uint64{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	for _, idx := range run0 {
		r0, r1, r2, r3, elapsed := sim0(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), idx)), run)
		ttl += elapsed
		fmt.Printf("randBench  (v2 0) %2v in %-14v - 0: "+frm+"; 1: "+frm+"; 2: "+frm+"; 3: "+frm+"\n", idx, elapsed, r0, r1, r2, r3)
	}
	fmt.Printf("                    avg %-14v\n\n", ttl/time.Duration(len(run0)))

	ttl = 0
	run1 := []uint64{30, 31, 32, 33, 34, 35, 36, 37, 38, 39}
	for _, idx := range run1 {
		r0, r1, r2, r3, elapsed := sim1(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), idx)), run)
		ttl += elapsed
		fmt.Printf("randBench  (v2 1) %2v in %-14v - 0: "+frm+"; 1: "+frm+"; 2: "+frm+"; 3: "+frm+"\n", idx, elapsed, r0, r1, r2, r3)
	}
	fmt.Printf("                    avg %-14v\n\n", ttl/time.Duration(len(run1)))

	ttl = 0
	run2 := []uint64{40, 41, 42, 43, 44, 45, 46, 47, 48, 49}
	for _, idx := range run2 {
		r0, r1, r2, r3, elapsed := sim2(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), idx)), run)
		ttl += elapsed
		fmt.Printf("randBench  (v2 2) %2v in %-14v - 0: "+frm+"; 1: "+frm+"; 2: "+frm+"; 3: "+frm+"\n", idx, elapsed, r0, r1, r2, r3)
	}
	fmt.Printf("                    avg %-14v\n", ttl/time.Duration(len(run2)))
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
		switch rnd.IntN(4) {
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
		val = rnd.Int64()
		if val < 2305843009213693952 {
			c0++
		} else if val < 4611686018427387904 {
			c1++
		} else if val < 6917529027641081856 {
			c2++
		} else {
			c3++
		}
	}
	lpsd = time.Since(strt)
	return
}

func sim2(
	rnd *rand.Rand,
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	var val uint64
	strt := time.Now()
	for range run {
		val = rnd.Uint64()
		if val < 4611686018427387904 {
			c0++
		} else if val < 9223372036854775808 {
			c1++
		} else if val < 13835058055282163712 {
			c2++
		} else {
			c3++
		}
	}
	lpsd = time.Since(strt)
	return
}
