package main

import (
	"fmt"
	"sort"
)

func main() {
	ogn := []uint8{4, 2, 7, 3, 9, 8}
	cln := make([]uint8, len(ogn))
	copy(cln, ogn)
	sort.Slice(cln, func(i, j int) bool {
		return cln[i] < cln[j]
	})
	fmt.Println(ogn)
	fmt.Println(cln)
}
