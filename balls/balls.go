package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// func build(types int, amount int) []int {
// 	result := make([]int, types*amount)
// 	// idx := 0
// 	for i := 0; i < types; i++ {
// 		for j := 0; j < amount; j++ {
// 			result[amount*i+j] = i
// 		}
// 	}
// 	return result
// }

// func sim(types int, amount int) (totals int, same bool, diff bool) {
// 	balls := build(types, amount)
// 	drawn := []int{} // TODO TEMP
// 	result := make([]int, types)

// 	for !same && !diff && len(balls) > 0 {
// 		index := rand.Intn(len(balls))
// 		drawn = append(drawn, balls[index]) // TODO TEMP
// 		result[balls[index]]++
// 		totals++
// 		balls = append(balls[:index], balls[index+1:]...)

// 		count := 0
// 		for _, cnt := range result {
// 			if cnt > 0 {
// 				count++
// 			}
// 			if cnt >= types {
// 				same = true
// 				break
// 			}
// 		}
// 		if count >= types {
// 			diff = true
// 			break
// 		}
// 	}
// 	fmt.Print(result, drawn)

// 	return
// }

func main() {
	// const trials = 10000

	// max := 0
	// cnts := 0
	// cntd := 0
	// for i := 0; i < trials; i++ {
	// 	totals, same, diff := sim(3, 6)
	// 	fmt.Println(" ->", totals, same, diff) // TODO TEMP
	// 	if totals > max {
	// 		max = totals
	// 	}
	// 	if same {
	// 		cnts++
	// 	}
	// 	if diff {
	// 		cntd++
	// 	}
	// }
	// fmt.Printf("max: %v | same %v | different %v\n", max, cnts, cntd)

	//for _, k := range g.Rand.Perm(len(rules.Territories[domain[j]].Connected)) {
	// x := rand.Perm(5)
	// fmt.Println(x)
	// for _, k := range x {
	// 	fmt.Println(k)
	// }

	// cnt := 2
	// for i := 0; i < 4; i++ {
	// 	// fmt.Println(x[cnt])
	// 	fmt.Print(cnt, " ")
	// 	if cnt != 0 {
	// 		fmt.Println(rand.Intn(cnt))
	// 	} else {
	// 		fmt.Println()
	// 	}
	// 	cnt--
	// }

	/*
	   func (g *Game) shuffleDeck() {
	   	for last := len(g.Deck) - 1; last > 0; last-- {
	   		indx := g.Rand.Intn(last)
	   		g.Deck[indx], g.Deck[last] = g.Deck[last], g.Deck[indx]
	   	}
	   }
	*/

	// fmt.Println(rand.Perm(10))

	arr := []int{} //{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	shuffle(arr)
	fmt.Println(arr)
}

func shuffle(arr []int) {
	for lst := len(arr) - 1; lst > 0; lst-- {
		idx := rand.Intn(lst)
		fmt.Printf("%v[%v] <-> %v[%v]\n", arr[lst], lst, arr[idx], idx)
		arr[idx], arr[lst] = arr[lst], arr[idx]
	}
}
