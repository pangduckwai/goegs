package main

import (
	"fmt"
)

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

// func main() {
// 	var s0 []int
// 	s1 := []int{1, 2, 3, 4, 5}
// 	var s2 []int
// 	s3 := []int{6, 7, 8, 9}

// 	s0 = append(s0, s1...)
// 	fmt.Println("1", s0)
// 	s0 = append(s0[1:], s2...)
// 	fmt.Println("2", s0)
// 	s0 = append(s0[1:], s3...)
// 	fmt.Println("3", s0)
// }

// func main() {
// 	s1 := []int{5, 6, 7, 8}
// 	fmt.Println("1", s1[1:])

// 	s2 := []int{9}
// 	fmt.Println("2", s2[1:])
// }

func main() {
	s := []int{5, 6, 7, 8, 9, 10, 11, 12}
	t := 5
	u := 8
	v := 9
	w := 12

	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("Hello %v\n", s[i])
		l := len(s) - 1
		if s[i] == t || s[i] == u || s[i] == v || s[i] == w {
			if i < 1 {
				s = s[i+1:]
			} else if i >= l {
				s = s[:l]
			} else {
				s = append(s[:i], s[i+1:]...)
			}
			i--
		}
	}

	fmt.Println(s)
}

// Obj Object
// type Obj struct {
// 	Name  string
// 	Value int
// }

// func (a Obj) Less(b Obj) bool {
// 	return a.Value < b.Value
// }

// func (o Obj) String() string {
// 	return fmt.Sprintf("%s", o.Name)
// }

// // Objs array of Obj
// type Objs []Obj

// func (o Objs) Len() int {
// 	return len(o)
// }
// func (o Objs) Swap(i, j int) {
// 	o[i], o[j] = o[j], o[i]
// }
// func (o Objs) Less(i, j int) bool {
// 	return o[i].Less(o[j])
// }

// // Find find an element in a slice
// func Find(list []Obj, item Obj) (found bool, idx int) {
// 	lgt := len(list)
// 	idx = sort.Search(
// 		lgt,
// 		func(i int) bool {
// 			return !list[i].Less(item) // list[i].Value >= item.Value
// 		},
// 	)
// 	found = (idx < lgt) && (list[idx].Value == item.Value)
// 	return
// }

// func Insert(list []Obj, item Obj, index int) []Obj {
// 	lgt := len(list)
// 	if index >= lgt {
// 		return append(list, item)
// 	}
// 	list = append(list[:index+1], list[index:]...)
// 	list[index] = item
// 	return list
// }

// func main() {
// 	s := []Obj{
// 		{"you", 17},
// 		{"there", 3},
// 		{"today", 19},
// 		{"hello", 2},
// 		{"are", 13},
// 	}

// 	fmt.Println(s)
// 	sort.Sort(Objs(s))
// 	fmt.Println(s)
// 	fmt.Println()

// 	t := Obj{"how", 11}
// 	u := Obj{"are", 13}
// 	v := Obj{"?", 21}

// 	start1 := time.Now()
// 	if f, i := Find(s, t); !f {
// 		s = Insert(s, t, i)
// 		fmt.Printf("%v -- %v\n", t, s)
// 	} else {
// 		fmt.Printf("%v already exists at %v\n", t, i)
// 	}

// 	start2 := time.Now()
// 	find1 := start2.Sub(start1)
// 	if f, i := Find(s, u); !f {
// 		s = Insert(s, u, i)
// 		fmt.Printf("%v -- %v\n", u, s)
// 	} else {
// 		fmt.Printf("%v already exists at %v\n", u, i)
// 	}

// 	start3 := time.Now()
// 	find2 := start3.Sub(start2)
// 	if f, i := Find(s, v); !f {
// 		s = Insert(s, v, i)
// 		fmt.Printf("%v -- %v\n", v, s)
// 	} else {
// 		fmt.Printf("%v already exists at %v\n", v, i)
// 	}
// 	find3 := time.Now().Sub(start3)

// 	fmt.Printf("Elapsed: %v %v %v -- %v", find1, find2, find3, find1+find2+find3)
// }
