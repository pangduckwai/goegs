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
	strt := time.Now()
	for range run {
		switch fastrand.Intn(4) {
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

func IntN(n int) int {
	val := fastrand.Uint64()

	switch n { // 0xFFFFFFFFFFFFFFFF = 18446744073709551616
	/*case 0:
	if val < 1 {
		return 0
	} else if val < 2 {
		return 1
	} else if val < 3 {
		return 2
	} else if val < 4 {
		return 3
	} else { // 5
		return 4
	}*/
	case 4:
		if val < 4611686018427387904 {
			return 0
		} else if val < 9223372036854775808 {
			return 1
		} else if val < 13835058055282163712 {
			return 2
		} else { // 18446744073709551616
			return 3
		}
	case 5:
		if val < 3689348814741910323 {
			return 0
		} else if val < 7378697629483820646 {
			return 1
		} else if val < 11068046444225730969 {
			return 2
		} else if val < 14757395258967641292 {
			return 3
		} else { // 18446744073709551615
			return 4
		}
	case 6:
		if val < 3074457345618258603 {
			return 0
		} else if val < 6148914691236517205 {
			return 1
		} else if val < 9223372036854775808 {
			return 2
		} else if val < 12297829382473034411 {
			return 3
		} else if val < 15372286728091293013 {
			return 4
		} else { // 18446744073709551614
			return 5
		}
	default:
		return -1 // TEMP!!! TODO: use Int2, Int3, Int10 instead
	}
}
