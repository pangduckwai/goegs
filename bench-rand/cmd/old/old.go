package main

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"sea9.org/go/egs/rand/pkg/common"
	"sea9.org/go/egs/rand/pkg/old"
)

func main() {
	// ////////////////// pprof /////////////////////
	// fcpu, err := os.Create("cmd/old/pgo/cpu.pprof")
	// if err != nil {
	// 	log.Fatal("[PRF] Failed to create CPU profile", err)
	// }
	// defer fcpu.Close()
	// if err = pprof.StartCPUProfile(fcpu); err != nil {
	// 	log.Fatal("[PRF] Failed to start CPU profiling", err)
	// }
	// defer pprof.StopCPUProfile()
	// ////////////////// pprof ///////////////////*/

	var err error
	var run uint64 = common.RUN_NUM
	var cmd, tmp int
	env := os.Getenv("RAND_RUN_NUM")
	if env != "" {
		tmp, err = strconv.Atoi(env)
		if err == nil {
			run = uint64(tmp)
		}
	}

	switch len(os.Args) {
	case 3:
		tmp, err = strconv.Atoi(os.Args[2])
		if err == nil {
			run = uint64(tmp)
		}
		fallthrough
	case 2:
		cmd, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	default:
		common.Usage(true)
	}

	if cmd < 1 || cmd > 7 {
		common.Usage(true)
	}

	prt := message.NewPrinter(language.English)
	old.Run(uint8(cmd), run, prt.Sprintf("runs: %v", run))

	// ////////////////// pprof /////////////////////
	// fmem, err := os.Create("cmd/old/pgo/mem.pprof")
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
