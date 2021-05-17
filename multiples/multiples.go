package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func simulate(sim int) (time.Duration, error) {
	quit := make(chan string)
	sims := make(chan string)

	simCnt := sim / 5
	now := time.Now()

	for i := 0; i < simCnt; i++ {
		for j := 0; j < 5; j++ {
			go func(m, n int) {
				str := time.Now()
				if n < 3 {
					time.Sleep(7 * time.Second)
				} else if n < 5 {
					time.Sleep(9 * time.Second)
				} else {
					quit <- "Error!!!"
				}
				sims <- fmt.Sprintf("%v %v elapsed %v", m, n, time.Now().Sub(str))
			}(i, j)
		}

		done := 0
		for done < 5 {
			select {
			case msg := <-quit:
				return time.Now().Sub(now), errors.New(msg)
			case msg := <-sims:
				fmt.Println(msg)
				done++
			}
		}
	}

	return time.Now().Sub(now), nil
}

func main() {
	sims := 12
	for sims%5 != 0 {
		sims++
	}
	fmt.Println("Run ", sims, "times")

	elapsed, err := simulate(sims)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Finished, total elasped time is %v\n", elapsed)
}
