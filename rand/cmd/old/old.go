package main

import (
	"log"
	"os"
	"strconv"

	"sea9.org/go/egs/rand/pkg/common"
	"sea9.org/go/egs/rand/pkg/old"
)

func main() {
	if len(os.Args) != 2 {
		common.Usage(true)
	}
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if c < 1 || c > 3 {
		common.Usage(true)
	}

	old.Run(uint8(c), common.RUN_NUM)
}
