package old

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"sea9.org/go/egs/rand/pkg/common"
)

func Run(c uint8, run uint64) {
	if c > 3 {
		fmt.Printf("Input range: 1, 2, 3")
		return
	}

	pad := int(math.Log10(float64(run)))

	trtn := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if c&1 > 0 {
		var ttl time.Duration
		fmt.Println("randBench: rand benchmark | math/rand - Intn()")
		for _, idx := range trtn {
			r0, r1, r2, r3, elapsed := sim0(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&2 > 0 {
		var ttl time.Duration
		fmt.Println("randBench: rand benchmark | math/rand - Int63()")
		for _, idx := range trtn {
			r0, r1, r2, r3, elapsed := sim1(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
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
