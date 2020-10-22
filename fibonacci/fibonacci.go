package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	last, next := 0, 1
	return func() int {
		v := last
		last, next = next, last+next // Note can't do that without multi assignment, as 'last' is overwritten by 'next', but 'next' need 'last'...
		return v
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Println(f())
	}
}
