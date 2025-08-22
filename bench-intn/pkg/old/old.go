package old

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"sea9.org/go/egs/randBench/pkg/common"
)

func Run(c uint8, n int, run uint64) {
	rw := false
	if c&8 > 0 {
		rw = true
	}
	cmd := c & 7
	if cmd > 3 {
		fmt.Println("Input range: 1, 2, 3")
		return
	}

	pad := int(math.Log10(float64(run)))

	trtn := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if cmd&1 > 0 {
		var ttl time.Duration
		fmt.Printf("OLD : rand benchmark | math/rand - Intn() | %v\n", run)
		for _, idx := range trtn {
			if !rw {
				rst, elapsed := simIntn(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), n, run)
				ttl += elapsed
				common.DisplayRun(idx, run, elapsed, rst, pad)
			} else {
				elapsed := simIntnRaw(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), run)
				ttl += elapsed
				common.DisplayRaw(idx, elapsed)
			}
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
	if cmd&2 > 0 {
		var ttl time.Duration
		fmt.Printf("OLD : rand benchmark | math/rand - Uint64 | %v\n", run)
		for _, idx := range trtn {
			if !rw {
				rst, elapsed := simUint64(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), n, run)
				ttl += elapsed
				common.DisplayRun(idx, run, elapsed, rst, pad)
			} else {
				elapsed := simUint64Raw(rand.New(rand.NewSource(time.Now().UnixNano()+int64(idx))), run)
				ttl += elapsed
				common.DisplayRaw(idx, elapsed)
			}
		}
		common.DisplayAvg(ttl, len(trtn), run)
	}
}

func simIntnRaw(
	rnd *rand.Rand,
	run uint64,
) (lpsd time.Duration) {
	strt := time.Now()
	for range run {
		rnd.Intn(7)
	}
	lpsd = time.Since(strt)
	return
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
		rst[rnd.Intn(n)]++
	}
	lpsd = time.Since(strt)
	return
}

func simUint64Raw(
	rnd *rand.Rand,
	run uint64,
) (lpsd time.Duration) {
	strt := time.Now()
	for range run {
		rnd.Uint64()
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
