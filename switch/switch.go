package main

import (
	"fmt"
)

func main() {
	vals := []int{10, 11, 12, 13, 14}
	switch lgth := len(vals); {
	case lgth > 4:
		fmt.Printf("5: %v\n", vals[4])
		fallthrough
	case lgth > 3:
		fmt.Printf("4: %v\n", vals[3])
		fallthrough
	case lgth > 2:
		fmt.Printf("3: %v\n", vals[2])
		fallthrough
	case lgth > 1:
		fmt.Printf("2: %v\n", vals[1])
		fallthrough
	case lgth > 0:
		fmt.Printf("1: %v\n", vals[0])
	}
}
