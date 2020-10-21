package main

import (
	"fmt"
)

// Pic return a XxY board
func Pic(dx, dy int) [][]uint8 {
	buff := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		row := make([]uint8, dy)
		for y := range row {
			row[y] = uint8(x ^ y) // bitwise XOR, try others such as (x+y)/2, x*y
		}
		buff[x] = row
	}
	return buff
}

func main() {
	result := Pic(8, 7)
	for _, row := range result {
		for _, val := range row {
			fmt.Print(val)
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
