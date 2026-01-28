package bench

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	val := 6000
	fmt.Printf("TestDivide() 1: %v\n", val>>1) // /2
	fmt.Printf("TestDivide() 2: %v\n", val>>2) // /4
	fmt.Printf("TestDivide() 3: %v\n", val>>3) // /8
}

func TestRange(t *testing.T) {
	out := make([]int, 0)
	for i := range 6 {
		out = append(out, i)
	}
	fmt.Printf("TestRange() %v\n", out)
}

func TestSim(t *testing.T) {
	totl := 1000000000
	bnch, lpsb := Bench(1, totl)
	stat, lpss := Sim(1, totl)
	fmt.Printf("TestSim() Bench duration: %v; statistic: %v\n", lpsb, bnch)
	fmt.Printf("TestSim() S i m duration: %v; statistic: %v\n", lpss, stat)
	fmt.Printf("  0, %v%%\n", float32(stat[0])/float32(totl)*100)
	fmt.Printf("  1, %v%%\n", float32(stat[1])/float32(totl)*100)
	fmt.Printf("  2, %v%%\n", float32(stat[2])/float32(totl)*100)
	fmt.Printf("  3, %v%%\n", float32(stat[3])/float32(totl)*100)
	fmt.Printf("  4, %v%%\n", float32(stat[4])/float32(totl)*100)
	fmt.Printf("  5, %v%%\n", float32(stat[5])/float32(totl)*100)
}
