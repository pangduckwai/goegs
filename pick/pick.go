package main

import "fmt"

func bsearch(vals []int, val int) int {
	for a, z := 0, len(vals)-1; ; {
		if a == z {
			return a
		}

		i := (z - a) / 2
		fmt.Println(i)

		if val < vals[i] {
			z = i
		} else if val > vals[i] {
			a = i
		} else {
			return i
		}
	}
}

func main() {
	vals := idx{2, 3, 4, 6, 7, 9}
	pick := []int{7, 9, 2}

	for _, p := range pick {
		for i := 0; i < len(vals); i++ {
			if vals[i] == p {
				vals = append(vals[:i], vals[i+1:]...)
				break
			}
		}
	}
	fmt.Println(vals)
}
