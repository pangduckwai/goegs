package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	troops := []uint32{
		5, 6, 12, 1, 44, 1, 10, 99, 8, 114, 25, 2, 31, 12, 4, 8, 3, 44, 282, 143, 21, 168, 36, 49, 84, 125, 17, 31, 20, 30, 11, 12, 8, 39, 7, 7, 15, 82, 43, 75, 211, 236,
	}
	// holdng := []uint8{
	// 	0, 0, 0, 1, 4, 1, 0, 2, 1, 2, 2, 4, 2, 0, 1, 2, 2, 2, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
	// }

	// sorted := make([]uint32, len(troops))
	// copy(sorted, troops)
	// sort.Slice(sorted, func(i, j int) bool {
	// 	return sorted[i] < sorted[j]
	// })
	// tmax := sorted[len(sorted)-1]
	tmax := uint32(0)
	for _, v := range troops {
		if v > tmax {
			tmax = v
		}
	}

	prcntg := make([]uint8, len(troops))
	for i, t := range troops {
		prcntg[i] = uint8((t * 100) / tmax)
	}

	scores := make([]uint8, len(troops))
	for i, t := range troops {
		scores[i] = uint8((t * 100) / tmax)
	}

	// fmt.Println(sorted)
	fmt.Println(tmax)
	fmt.Println(troops)
	fmt.Println(prcntg)
	fmt.Println(scores)

	fmt.Printf("%x\n", sha1.Sum(scores))
}
