package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("'%v'\n", os.Getenv("ENV_TEST"))
}
