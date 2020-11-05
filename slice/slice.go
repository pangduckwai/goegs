package main

import "fmt"

// type test struct {
// 	s1 []int
// 	s2 []int
// 	s3 []int
// 	s4 []int
// 	s5 []int
// }

// func main() {
// 	s3 := make([]int, 5)
// 	for i := range s3 {
// 		s3[i] = i
// 	}
// 	s4 := []int{5, 4, 3, 2, 1}

// 	t := test{
// 		make([]int, 3),
// 		nil,
// 		s3,
// 		s4,
// 		make([]int, 0),
// 	}

// 	fmt.Println(t.s1)
// 	fmt.Println(t.s2)
// 	fmt.Println(t.s3)
// 	fmt.Println(t.s4)
// 	fmt.Println(t.s5)

// 	fmt.Println(len(t.s2))
// 	fmt.Println(len(t.s5))

// 	for i := 1; i <= 3; i++ {
// 		t.s2 = append(t.s2, i*2)
// 	}
// 	fmt.Println(t.s2)
// 	fmt.Println()

// 	ss := []int{1, 2}
// 	fmt.Println("OH", ss)
// 	ss = ss[1:]
// 	fmt.Println("OH", ss)
// 	ss = ss[1:]
// 	fmt.Println("OH", ss)
// 	// ss = ss[1:]
// 	// fmt.Println("OH", ss)
// }

func main() {
	var s0 []int
	s1 := []int{1, 2, 3, 4, 5}
	var s2 []int
	s3 := []int{6, 7, 8, 9}

	s0 = append(s0, s1...)
	fmt.Println("1", s0)
	s0 = append(s0[1:], s2...)
	fmt.Println("2", s0)
	s0 = append(s0[1:], s3...)
	fmt.Println("3", s0)
}
