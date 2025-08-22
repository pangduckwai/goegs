package fast

import (
	"fmt"
	"math"
	"time"

	"sea9.org/go/egs/randBench/pkg/bytedance/fastrand"
	"sea9.org/go/egs/randBench/pkg/common"
)

func Run(c uint8, run uint64) {
	if c > 3 {
		fmt.Println("Input range: 1, 2, 3")
		return
	}

	pad := int(math.Log10(float64(run)))

	trtn := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if c&1 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Intn() | %v\n", run)
		for _, idx := range trtn {
			r0, r1, r2, r3, r4, r5, r6, elapsed := sim0(run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, r4, r5, r6, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&2 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Uint64 | %v\n", run)
		for _, idx := range trtn {
			r0, r1, r2, r3, r4, r5, r6, elapsed := sim1(run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, r4, r5, r6, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
}

func sim0(
	run uint64,
) (
	c0, c1, c2, c3, c4, c5, c6 uint64,
	lpsd time.Duration,
) {
	strt := time.Now()
	for range run {
		switch fastrand.Intn(7) {
		case 0:
			c0++
		case 1:
			c1++
		case 2:
			c2++
		case 3:
			c3++
		case 4:
			c4++
		case 5:
			c5++
		case 6:
			c6++
		}
	}
	lpsd = time.Since(strt)
	return
}

func sim1(
	run uint64,
) (
	c0, c1, c2, c3, c4, c5, c6 uint64,
	lpsd time.Duration,
) {
	strt := time.Now()
	for range run {
		switch common.RandN(7, fastrand.Uint64()) {
		case 0:
			c0++
		case 1:
			c1++
		case 2:
			c2++
		case 3:
			c3++
		case 4:
			c4++
		case 5:
			c5++
		case 6:
			c6++
		}
	}
	lpsd = time.Since(strt)
	return
}
