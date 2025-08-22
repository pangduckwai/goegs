package fast

import (
	"fmt"
	"math"
	"time"

	"sea9.org/go/egs/rand/pkg/bytedance/fastrand"
	"sea9.org/go/egs/rand/pkg/common"
)

func Run(c uint8, run uint64, msg string) {
	if c > 7 {
		fmt.Println("Input range: 1, 2, 3, 4, 5, 6, 7")
		return
	}

	pad := int(math.Log10(float64(run)))

	trtn := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if c&1 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Intn() | %v\n", msg)
		for _, idx := range trtn {
			r0, r1, r2, r3, elapsed := sim0(run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&2 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Int63() | %v\n", msg)
		for _, idx := range trtn {
			r0, r1, r2, r3, elapsed := sim1(run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&4 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Uint64() | %v\n", msg)
		for _, idx := range trtn {
			r0, r1, r2, r3, elapsed := sim2(run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, r0, r1, r2, r3, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
}

func sim0(
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	var val int
	strt := time.Now()
	for range run {
		val = fastrand.Intn(4)
		if val < 1 {
			c0++
		} else if val < 2 {
			c1++
		} else if val < 3 {
			c2++
		} else {
			c3++
		}
	}
	lpsd = time.Since(strt)
	return
}

func sim1(
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	var val int64
	strt := time.Now()
	for range run {
		val = fastrand.Int63()
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

func sim2(
	run uint64,
) (
	c0, c1, c2, c3 uint64,
	lpsd time.Duration,
) {
	var val uint64
	strt := time.Now()
	for range run {
		val = fastrand.Uint64()
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
