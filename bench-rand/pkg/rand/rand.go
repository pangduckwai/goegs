package rand

import (
	"math/bits"
	"math/rand"
	"time"
)

func rndRand(rnd *rand.Rand, n int) int {
	hi, _ := bits.Mul64(rnd.Uint64(), uint64(n))
	return int(hi)
}

func Sim(id, n, run int) (lps time.Duration, cnt []int, nmz []float32) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
	cnt = make([]int, 0)
	for range n {
		cnt = append(cnt, 0)
	}

	now := time.Now()
	for range run {
		cnt[rndRand(rnd, n)] += 1
	}
	lps = time.Since(now)

	nmz = make([]float32, 0)
	for _, c := range cnt {
		nmz = append(nmz, float32(c)/float32(run))
	}

	return
}
