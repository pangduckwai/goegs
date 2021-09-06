package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func build(types int, amount int) []int {
	result := make([]int, types*amount)
	// idx := 0
	for i := 0; i < types; i++ {
		for j := 0; j < amount; j++ {
			result[amount*i+j] = i
		}
	}
	return result
}

func sim(types int, amount int) (totals int, same bool, diff bool) {
	balls := build(types, amount)
	drawn := []int{} // TODO TEMP
	result := make([]int, types)

	for !same && !diff && len(balls) > 0 {
		index := rand.Intn(len(balls))
		drawn = append(drawn, balls[index]) // TODO TEMP
		result[balls[index]]++
		totals++
		balls = append(balls[:index], balls[index+1:]...)

		count := 0
		for _, cnt := range result {
			if cnt > 0 {
				count++
			}
			if cnt >= types {
				same = true
				break
			}
		}
		if count >= types {
			diff = true
			break
		}
	}
	fmt.Print(result, drawn)

	return
}

func main() {
	const trials = 10000

	max := 0
	cnts := 0
	cntd := 0
	for i := 0; i < trials; i++ {
		totals, same, diff := sim(3, 6)
		fmt.Println(" ->", totals, same, diff) // TODO TEMP
		if totals > max {
			max = totals
		}
		if same {
			cnts++
		}
		if diff {
			cntd++
		}
	}
	fmt.Printf("max: %v | same %v | different %v\n", max, cnts, cntd)
}
