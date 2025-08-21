package main

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sea9.org/go/egs/rand/pkg/common"
	"sea9.org/go/egs/rand/pkg/fast"
)

func main() {
	// ////////////////// pprof /////////////////////
	// fcpu, err := os.Create("cmd/fast/pgo/cpu.pprof")
	// if err != nil {
	// 	log.Fatal("[PRF] Failed to create CPU profile", err)
	// }
	// defer fcpu.Close()
	// if err = pprof.StartCPUProfile(fcpu); err != nil {
	// 	log.Fatal("[PRF] Failed to start CPU profiling", err)
	// }
	// defer pprof.StopCPUProfile()
	// ////////////////// pprof ///////////////////*/

	if len(os.Args) != 2 {
		common.Usage(true)
	}
	c, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if c < 1 || c > 7 {
		common.Usage(true)
	}

	var r uint64 = common.RUN_NUM
	e := os.Getenv("RAND_RUN_NUM")
	if e != "" {
		t, err := strconv.Atoi(e)
		if err == nil {
			r = uint64(t)
		}
	}

	prt := message.NewPrinter(language.English)
	fast.Run(uint8(c), r, prt.Sprintf("runs: %v", r))

	// ////////////////// pprof /////////////////////
	// fmem, err := os.Create("cmd/fast/pgo/mem.pprof")
	// if err != nil {
	// 	log.Fatal("[PRF] Failed to create Memory profile", err)
	// }
	// defer fmem.Close()
	// runtime.GC() // get up-to-date statistics
	// if err = pprof.WriteHeapProfile(fmem); err != nil {
	// 	log.Fatal("[PRF] Failed to start Memory profiling", err)
	// }
	// ////////////////// pprof ///////////////////*/
}
