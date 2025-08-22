package fast

import (
	"fmt"
	"math"
	"time"

	"sea9.org/go/egs/randBench/pkg/bytedance/fastrand"
	"sea9.org/go/egs/randBench/pkg/common"
)

func Run(c uint8, n int, run uint64) {
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
			rst, elapsed := simIntn(n, run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, rst, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&2 > 0 {
		var ttl time.Duration
		fmt.Printf("FAST: rand benchmark | fastrand - Uint64 | %v\n", run)
		for _, idx := range trtn {
			rst, elapsed := simUint64(n, run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, rst, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
}

func simIntn(
	n int,
	run uint64,
) (
	rst []uint64,
	lpsd time.Duration,
) {
	rst = make([]uint64, n)
	strt := time.Now()
	for range run {
		rst[fastrand.Intn(n)]++
	}
	lpsd = time.Since(strt)
	return
}

func simUint64(
	n int,
	run uint64,
) (
	rst []uint64,
	lpsd time.Duration,
) {
	rst = make([]uint64, n)
	strt := time.Now()
	for range run {
		rst[common.RandN(n, fastrand.Uint64())]++
	}
	lpsd = time.Since(strt)
	return
}
