package ver2

import (
	"math/bits"
	"math/rand/v2"
	"time"
)

func v2Rand(rnd *rand.Rand, n int) int {
	hi, _ := bits.Mul64(rnd.Uint64(), uint64(n))
	return int(hi)
}

func Sim(id, n, run int) (lps time.Duration, cnt []int, nmz []float32) {
	rnd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(id)))
	cnt = make([]int, 0)
	for range n {
		cnt = append(cnt, 0)
	}

	now := time.Now()
	for range run {
		cnt[v2Rand(rnd, n)] += 1
	}
	lps = time.Since(now)

	nmz = make([]float32, 0)
	for _, c := range cnt {
		nmz = append(nmz, float32(c)/float32(run))
	}

	return
}
