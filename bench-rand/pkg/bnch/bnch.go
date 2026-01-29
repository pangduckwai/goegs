package bnch

import (
	"time"
)

func Run(
	id, rng, run int,
	rfunc func() int,
) (
	lps time.Duration,
	cnt []int,
	nmz []float32,
) {
	cnt = make([]int, 0)
	for range rng {
		cnt = append(cnt, 0)
	}

	now := time.Now()
	for range run {
		cnt[rfunc()] += 1
	}
	lps = time.Since(now)

	nmz = make([]float32, 0)
	for _, c := range cnt {
		nmz = append(nmz, float32(c)/float32(run))
	}

	return
}
