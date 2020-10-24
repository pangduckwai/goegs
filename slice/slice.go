package main

import "fmt"

type test struct {
	s1 []int
	s2 []int
	s3 []int
	s4 []int
	s5 []int
}

func main() {
	s3 := make([]int, 5)
	for i := range s3 {
		s3[i] = i
	}
	s4 := []int{5, 4, 3, 2, 1}

	t := test{
		make([]int, 3),
		nil,
		s3,
		s4,
		make([]int, 0),
	}

	fmt.Println(t.s1)
	fmt.Println(t.s2)
	fmt.Println(t.s3)
	fmt.Println(t.s4)
	fmt.Println(t.s5)

	fmt.Println(len(t.s2))
	fmt.Println(len(t.s5))

	for i := 1; i <= 3; i++ {
		t.s2 = append(t.s2, i*2)
	}
	fmt.Println(t.s2)
}
