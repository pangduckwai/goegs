package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// var holdings = []uint8{
// 	0, 1, 2, 3, 4, 5,
// 	1, 2, 3, 4, 5, 0,
// 	2, 3, 4, 5, 0, 1,
// 	3, 4, 5, 0, 1, 2,
// 	4, 5, 0, 1, 2, 3,
// 	5, 0, 1, 2, 3, 4,
// 	0, 1, 2, 3, 4, 5,
// }

// func dominion(p uint8) []uint8 {
// 	r := make([]uint8, 42)
// 	i := 0
// 	for j, o := range holdings {
// 		if o == p {
// 			r[i] = uint8(j)
// 			i++
// 		}
// 	}
// 	return r[:i]
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		log.Fatal("Usage: test {num}")
// 	}
// 	num, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	d := dominion(uint8(num))
// 	d = nil
// 	fmt.Println(d)
// 	for _, j := range rand.Perm(len(d)) {
// 		fmt.Println(" ", j, d[j])
// 	}
// }

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
sea9.org/go/wdomc/node.Expand(0xc0002e7170, 0xc01101b000, 0x6, 0x1c03, 0x2c, 0xc023b7fc98, 0xa80860, 0xc01101b000, 0xc01101b000, 0x26, ...)
        /home/paul_lai/go/wdom-mc/node/node.go:154 +0x5b
sea9.org/go/wdomc/game.Playout(0xc0002d0750, 0x4, 0xc0002d0754, 0x6, 0xc03d9a70a0, 0xc00ba00600, 0xa80878, 0x100000e, 0x23cb000023cc, 0x64, ...)
        /home/paul_lai/go/wdom-mc/game/game.go:275 +0x5e9
main._simulate.func1(0xc0002da000, 0xc134f76780, 0xc070a12948, 0x16, 0x1b8, 0xc0002d4e00, 0x70, 0x70, 0xc00ba00600, 0xc134f76840, ...)
        /home/paul_lai/go/wdom-mc/simulation.go:235 +0x2f3
created by main._simulate
        /home/paul_lai/go/wdom-mc/simulation.go:227 +0x270
*/

// func main() {
// 	n := 0
// 	v := math.Log(float64(n))
// 	fmt.Println("Log", n, "is", v)
// }

// func main() {
// 	buf := []uint8{1, 2, 3, 4, 5, 6, 7}
// 	fmt.Println(buf)

// 	for j, v := range buf {
// 		if v == 0 {
// 			buf = append([]uint8{v}, append(buf[:j], buf[j+1:]...)...)
// 			break
// 		}
// 	}

// 	fmt.Println(buf)
// }

// func main() {
// 	quit, sims := make(chan string), make(chan string)
// 	var wgrp sync.WaitGroup
// 	n := rand.Intn(10)
// 	fmt.Println("Choice", n)

// 	wgrp.Add(1)
// 	go func() {
// 		defer wgrp.Done()
// 		time.Sleep(8 * time.Second)
// 		sims <- fmt.Sprintf("run %vs", 8)
// 	}()

// 	wgrp.Add(1)
// 	go func() {
// 		defer wgrp.Done()
// 		time.Sleep(3 * time.Second)
// 		sims <- fmt.Sprintf("run %vs", 3)
// 	}()

// 	wgrp.Add(1)
// 	go func() {
// 		time.Sleep(4 * time.Second)
// 		if n > 5 {
// 			quit <- "QUIT"
// 		} else {
// 			defer wgrp.Done()
// 			sims <- fmt.Sprintf("run %vs", 4)
// 		}
// 	}()

// 	go func() {
// 		defer close(quit)
// 		wgrp.Wait()
// 	}()

// wait:
// 	for {
// 		select {
// 		case msg := <-sims:
// 			fmt.Println(msg)
// 		case msg := <-quit:
// 			if msg != "" {
// 				log.Println(msg)
// 			}
// 			break wait
// 		}
// 	}
// 	fmt.Println("The End!")
// }

// type node struct {
// 	values uint32
// }

// func (n *node) value1() uint16 {
// 	return uint16(n.values & 0x0000FFFF)
// }

// func (n *node) setValue1(v uint16) {
// 	n.values = n.values | uint32(v)
// }

// func (n *node) value2() uint16 {
// 	return uint16((n.values & 0xFFFF0000) >> 16)
// }

// func (n *node) setValue2(v uint16) {
// 	n.values = n.values | (uint32(v) << 16)
// }

// func main() {
// 	n := &node{}
// 	n.setValue1(7890)
// 	n.setValue2(65534)

// 	v1 := n.value1()
// 	v2 := n.value2()
// 	fmt.Printf("%v: 1:%v / 2:%v", n.values, v1, v2)
// }

// func test(val int) {
// 	fmt.Println("1", val)
// 	val += 20
// 	fmt.Println("2", val)
// 	val = 99
// 	fmt.Println("3", val)
// }

// func main() {
// 	v := []uint8{4, 3, 5, 8, 1, 7, 2}
// 	w := strings.Replace(fmt.Sprintf("%v", v), " ", ",", -1)
// 	fmt.Printf("%v\n", v)
// 	fmt.Println(w)
// 	test(24)
// }

// func main() {
// 	dat := uint8(0)
// 	fmt.Printf("%08b | 0x%02x | Ready:%5v | Finished:%5v\n", dat, dat, (dat&0x1) > 0, (dat&0x80) > 0)

// 	dat |= 0x1
// 	fmt.Printf("%08b | 0x%02x | Ready:%5v | Finished:%5v\n", dat, dat, (dat&0x1) > 0, (dat&0x80) > 0)

// 	dat |= 0x2
// 	fmt.Printf("%08b | 0x%02x | Ready:%5v | Finished:%5v\n", dat, dat, (dat&0x1) > 0, (dat&0x80) > 0)

// 	dat |= 0x10
// 	fmt.Printf("%08b | 0x%02x | Ready:%5v | Finished:%5v\n", dat, dat, (dat&0x1) > 0, (dat&0x80) > 0)

// 	dat &= 0xFE
// 	fmt.Printf("%08b | 0x%02x | Ready:%5v | Finished:%5v\n", dat, dat, (dat&0x1) > 0, (dat&0x80) > 0)
// }

// func main() {
// 	val := uint8(1)
// 	val--
// 	fmt.Println(val)
// 	val--
// 	fmt.Println(val)

// 	var a1 []int
// 	a1 = append(a1, 1, 3)
// 	fmt.Println(a1)

// 	var a2 [][]int
// 	var a3 []int
// 	a3 = append(a3, 2, 3)
// 	a2 = append([][]int{a3}, a1)
// 	fmt.Println(a2)
// }

// type Data uint16

// func disp(val uint16) {
// 	fmt.Printf("%b\n", val)
// }

// func main() {
// 	// var dat Data
// 	dat := Data(0xEFD3)
// 	fmt.Println(dat)
// 	disp(uint16(dat))
// }

// func main() {
// 	str := "abcde-fghi-jklmno"
// 	fmt.Printf("%v\n", str[0:len(str)-7])
// }

// func main() {
// 	n := 42
// 	t := float64(1)
// 	l := float64(1)
// 	fmt.Println(4)
// 	fmt.Println(1)
// 	for i := n; i > 38; i-- {
// 		l = l * float64(i)
// 		t += l
// 		fmt.Printf("‡ %.0f\n", l)
// 	}
// 	fmt.Printf("‡ %.0f\n", t)
// }

// func main() {
// 	fmt.Printf("%v\n", Version("11.139.04.beta.202306"))
// 	fmt.Printf("%v\n", Version("11.139.04.beta.202306", false))
// 	fmt.Printf("%v\n", Version("11.139.04.beta.202306", true))
// 	fmt.Printf("%v\n", Version("11.139.04.202306"))
// 	fmt.Printf("%v\n", Version("11.139.04.202306", false))
// 	fmt.Printf("%v\n", Version("11.139.04.202306", true))
// 	fmt.Printf("%v\n", Version("11.139.04"))
// 	fmt.Printf("%v\n", Version("11.139.04", false))
// 	fmt.Printf("%v\n", Version("11.139.04", true))
// }

// // Version wdomc version. Format: {Major-version}.{Minor-version}.{Revision}.[Tag].[Branch].{Timestamp}
// func Version(version string, full ...bool) string {
// 	if strings.Count(version, ".") <= 3 || (len(full) > 0 && full[0]) {
// 		return version + " (is full)"
// 	}

// 	return version[0:strings.LastIndex(version, ".")] + " (not full)"
// }

func main() {
	player := 6
	CardCount := 44

	var deck []uint8
	cards := make([][]uint8, player)
	for i, j := 0, 0; i < CardCount; i++ {
		j = i / 2
		if j < player {
			cards[j] = append(cards[j], uint8(i))
			// fmt.Printf("i:%v -> cards[%v][?]\n", i, j)
		} else {
			deck = append(deck, uint8(i))
			// fmt.Printf("deck[%v]\n", i)
		}
	}

	fmt.Println(cards)
	fmt.Println(deck)
}
