package main

import (
	"fmt"
)

func main() {
	// if len(os.Args) != 2 {
	// 	log.Fatal("Usage: modular {num}")
	// }
	// num, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("       __%2__ __%3__ __%4__ __%?__")
	for i := 0; i < 30; i++ {
		var x int
		switch {
		case i < 6:
			x = 2
		case i < 12:
			x = 3
		default:
			x = 4
		}
		fmt.Printf("lvl:%2v mod:%2v mod:%2v mod:%2v mod:%2v\n", i, i%2, i%3, i%4, i%x)
	}
}
