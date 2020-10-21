package main

import (
	"fmt"
	"strings"
)

// WordCount return word counts
func WordCount(s string) map[string]int {
	buf := strings.Fields(s)
	rst := make(map[string]int)
	for _, str := range buf {
		if elm, exists := rst[str]; exists {
			rst[str] = elm + 1
		} else {
			rst[str] = 1
		}
	}
	return rst
}

func main() {
	fmt.Println(WordCount("Hello there, how are you? I'm fine, thanks, and you?"))
}
