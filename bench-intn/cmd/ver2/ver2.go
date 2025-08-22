package main

import (
	"log"
	"os"
	"strconv"

	"sea9.org/go/egs/randBench/pkg/common"
	"sea9.org/go/egs/randBench/pkg/ver2"
)

func main() {
	// ////////////////// pprof /////////////////////
	// fcpu, err := os.Create("cmd/ver2/pgo/cpu.pprof")
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
	var n int = common.SEL_NUM
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
	case 4:
		tmp, err = strconv.Atoi(os.Args[3])
		if err == nil {
			run = uint64(tmp)
		}
		fallthrough
	case 3:
		tmp, err = strconv.Atoi(os.Args[2])
		if err == nil {
			n = tmp
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

	c := cmd & 7
	if c < 1 || c > 7 {
		common.Usage(true)
	}

	ver2.Run(uint8(cmd), n, run)

	// ////////////////// pprof /////////////////////
	// fmem, err := os.Create("cmd/ver2/pgo/mem.pprof")
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
