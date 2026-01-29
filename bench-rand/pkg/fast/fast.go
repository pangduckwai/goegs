package fast

import (
	"time"

	"sea9.org/go/egs/randBench/pkg/bytedance/fastrand"
)

// fastRand returns a uniform value in [0,n)
func fastRand(n int) int {
	return int(fastrand.Uint32n(uint32(n)))
}

func Sim(id, n, run int) (lps time.Duration, cnt []int, nmz []float32) {
	cnt = make([]int, 0)
	for range n {
		cnt = append(cnt, 0)
	}

	now := time.Now()
	for range run {
		cnt[fastRand(n)] += 1
	}
	lps = time.Since(now)

	nmz = make([]float32, 0)
	for _, c := range cnt {
		nmz = append(nmz, float32(c)/float32(run))
	}

	return
}
