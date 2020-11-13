package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	v0 := 7
	v1 := 87
	v2 := 777
	format := fmt.Sprintf("Hello [%%%dv]", int(math.Log10(float64(v1)))+1)
	fmt.Println(format)
	fmt.Printf(format+" there.\n", v0)
	fmt.Printf(format+" there?\n", v1)
	fmt.Printf(format+" there!\n", v2)
}
