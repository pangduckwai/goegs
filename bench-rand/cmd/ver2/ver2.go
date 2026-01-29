package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sea9.org/go/egs/randBench/pkg/ver2"
)

func main() {
	// go build -pgo=pgo/cpu.pprof
	/*///////////////// pprof /////////////////////
	dir, _ := filepath.Split(os.Args[0])
	fcpu, e := os.Create(filepath.Join(dir, "pgo", "cpu.pprof"))
	if e != nil {
		log.Fatal("[PRF] Failed to create CPU profile", e)
	}
	defer fcpu.Close()
	if e = pprof.StartCPUProfile(fcpu); e != nil {
		log.Fatal("[PRF] Failed to start CPU profiling", e)
	}
	defer pprof.StopCPUProfile()
	////////////////// pprof ///////////////////*/

	var err error
	run := 1000000000 // 1,000,000,000
	rng := 6
	switch len(os.Args) {
	case 3:
		run, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fallthrough
	case 2:
		rng, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	case 1:
	default:
		log.Println("Usage: cmd/ver2 [range] [num-of-runs]")
	}

	lps, cnt, nmz := ver2.Sim(1, rng, run)

	var buf strings.Builder
	for i, v := range cnt {
		fmt.Fprintf(&buf, " %3v: %v (%.4f%%)\n", i, v, nmz[i]*100)
	}

	prt := message.NewPrinter(language.English)
	prt.Printf("[VER2] %v simulations with [0,%v) range, elapsed time: %12v (%v per op)\n%v", run, rng, lps, lps/time.Duration(run), buf.String())

	/*///////////////// pprof /////////////////////
	fmem, e := os.Create(filepath.Join(dir, "pgo", "mem.pprof"))
	if e != nil {
		log.Fatal("[PRF] Failed to create Memory profile", e)
	}
	defer fmem.Close()
	runtime.GC() // get up-to-date statistics
	if e = pprof.WriteHeapProfile(fmem); e != nil {
		log.Fatal("[PRF] Failed to start Memory profiling", e)
	}
	////////////////// pprof ///////////////////*/
}
