package bench

import "math/rand"

// max val 9223372036854775807
const T1 = 1537228672809129301
const T2 = 3074457345618258602
const T3 = 4611686018427387904
const T4 = 6148914691236517205
const T5 = 7686143364045646506

func Dice(rnd *rand.Rand) int64 {
	v := rnd.Int63()
	if v < T1 {
		return 0
	} else if v < T2 {
		return 1
	} else if v < T3 {
		return 2
	} else if v < T4 {
		return 3
	} else if v < T5 {
		return 4
	} else {
		return 5
	}
}

// func Sim() {
// 	rnd := rand.New(rand.NewSource(time.Now().UnixNano()+1))
// }
