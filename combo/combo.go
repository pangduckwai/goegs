package main

import (
	"fmt"
	"math/bits"
)

// All returns all combinations for a given string array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
// func All(set []string) (subsets [][]string) {
// 	length := uint(len(set))

// 	// Go through all possible combinations of objects
// 	// from 1 (only first object in subset) to 2^length (all objects in subset)
// 	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
// 		var subset []string

// 		for object := uint(0); object < length; object++ {
// 			// checks if object is contained in subset
// 			// by checking if bit 'object' is set in subsetBits
// 			if (subsetBits>>object)&1 == 1 {
// 				// add object to subset
// 				subset = append(subset, set[object])
// 			}
// 		}
// 		// add subset to subsets
// 		subsets = append(subsets, subset)
// 	}
// 	return subsets
// }

// Combo1 returns combinations of n elements for a given string array.
// For n < 1, it equals to All and returns all combinations.
func Combo1(set []uint8, n int) (combo [][]uint8) {
	lgth := uint(len(set))
	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^lgth (all objects in subset)
	for b := 1; b < (1 << lgth); b++ {
		if n > 0 && bits.OnesCount(uint(b)) != n {
			continue
		}

		var c []uint8
		for o := uint(0); o < lgth; o++ {
			// checks if o is contained in c
			// by checking if bit 'object' is set in b
			if (b>>o)&1 == 1 {
				// add o to c
				c = append(c, set[o])
			}
		}
		// add c to combo
		combo = append(combo, c)
	}
	return combo
}

// Combo2 returns combinations of n elements for a given string array.
// func Combo2(set []uint8, n int) (combo [][]uint8) {
// 	lgth := uint(len(set))
// 	if n > len(set) {
// 		n = len(set)
// 	}

// 	// Go through all possible combinations of objects
// 	// from 1 (only first object in subset) to 2^lgth (all objects in subset)
// 	for b := 1; b < (1 << lgth); b++ {
// 		if n > 0 && bits.OnesCount(uint(b)) != n {
// 			continue
// 		}

// 		var c []uint8
// 		for o := uint(0); o < lgth; o++ {
// 			f := false
// 			for _, t := range c {
// 				if t == set[o] {
// 					f = true
// 					break
// 				}
// 			}
// 			if !f {
// 				// add o to c
// 				c = append(c, set[o])
// 			}
// 		}
// 		// add c to combo
// 		combo = append(combo, c)
// 	}
// 	return combo
// }

func main() {
	src := []uint8{3, 1, 2, 4}
	rn := Combo1(src, 3)
	for _, r := range rn {
		fmt.Println(r)
	}
	fmt.Println(src)

	// rs := All([]string{"a", "b", "c", "d"})
	// for _, r := range rs {
	// 	fmt.Println(r)
	// }
}
