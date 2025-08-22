package ver2

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

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
		fmt.Printf("VER2: rand benchmark | math/rand/v2 - Intn() | %v\n", run)
		for _, idx := range trtn {
			rst, elapsed := simIntn(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), idx)), n, run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, rst, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if c&2 > 0 {
		var ttl time.Duration
		fmt.Printf("VER2: rand benchmark | math/rand/v2 - Uint64 | %v\n", run)
		for _, idx := range trtn {
			rst, elapsed := simUint64(rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), idx)), n, run)
			ttl += elapsed
			common.DisplayRun(idx, run, elapsed, rst, pad)
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
}

func simIntn(
	rnd *rand.Rand,
	n int,
	run uint64,
) (
	rst []uint64,
	lpsd time.Duration,
) {
	rst = make([]uint64, n)
	strt := time.Now()
	for range run {
		rst[rnd.IntN(n)]++
	}
	lpsd = time.Since(strt)
	return
}

func simUint64(
	rnd *rand.Rand,
	n int,
	run uint64,
) (
	rst []uint64,
	lpsd time.Duration,
) {
	rst = make([]uint64, n)
	strt := time.Now()
	for range run {
		rst[common.RandN(n, rnd.Uint64())]++
	}
	lpsd = time.Since(strt)
	return
}
