package bench

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	val := 6000
	fmt.Printf("1: %v\n", val>>1) // /2
	fmt.Printf("2: %v\n", val>>2) // /4
	fmt.Printf("3: %v\n", val>>3) // /8
}
