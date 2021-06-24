package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var holdings = []uint8{
	0, 1, 2, 3, 4, 5,
	1, 2, 3, 4, 5, 0,
	2, 3, 4, 5, 0, 1,
	3, 4, 5, 0, 1, 2,
	4, 5, 0, 1, 2, 3,
	5, 0, 1, 2, 3, 4,
	0, 1, 2, 3, 4, 5,
}

func dominion(p uint8) []uint8 {
	r := make([]uint8, 42)
	i := 0
	for j, o := range holdings {
		if o == p {
			r[i] = uint8(j)
			i++
		}
	}
	return r[:i]
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: test {num}")
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	d := dominion(uint8(num))
	d = nil
	fmt.Println(d)
	for _, j := range rand.Perm(len(d)) {
		fmt.Println(" ", j, d[j])
	}
}

/*
func (g *Game) Dominion(p uint8) []uint8 {
	r := make([]uint8, territoryCount)
	i := 0
	for j, o := range g.Holdings {
		if o == p {
			r[i] = uint8(j)
			i++
		}
	}
	return r[:i]
}

for _, j := range impl.game.Rand.Perm(len(domain)) {

2021/06/22 01:28:20 [WDOM-MC][ 440- 6- 80][Error] Expand() fatal error
panic: runtime error: index out of range [-1]

math/rand.(*rngSource).Uint64(...)
        /usr/local/go/src/math/rand/rng.go:249
math/rand.(*rngSource).Int63(0xc0003dd500, 0x3c6fdd596d864715)
        /usr/local/go/src/math/rand/rng.go:234 +0x98
math/rand.(*Rand).Int63(...)
        /usr/local/go/src/math/rand/rand.go:85
math/rand.(*Rand).Int31(...)
        /usr/local/go/src/math/rand/rand.go:99
math/rand.(*Rand).Int31n(0xc0002e7170, 0x3, 0xc000000001)
        /usr/local/go/src/math/rand/rand.go:134 +0x5f
math/rand.(*Rand).Intn(0xc0002e7170, 0x3, 0x1)
        /usr/local/go/src/math/rand/rand.go:172 +0x45
math/rand.(*Rand).Perm(0xc0002e7170, 0x26, 0xc02ea2c390, 0x26, 0x2a)
        /usr/local/go/src/math/rand/rand.go:225 +0x9c
sea9.org/go/wdomc/game.Variant1.ChooseMove(0xc03d9a70a0, 0x0, 0x967bc0, 0xc01101af80, 0xc01101b000, 0xc023b7fb0f)
        /home/paul_lai/go/wdom-mc/game/variants.go:137 +0x4de
*/
