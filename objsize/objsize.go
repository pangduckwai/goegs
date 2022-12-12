package main

import (
	"fmt"
	"runtime"
)

const buffsz = 10000000

// const start = maxWR - 5 - buffsz // 4284967290 - 4294967290
const start = maxWR - (buffsz / 2) // 4289967295 - 4299967295
// const start = maxWR + 5 // 4294967300 - 4304967300
const mib = float64(1048576)

var mem runtime.MemStats

func main() {
	// n0 := make([]*Node0, buffsz)
	n0 := make([]*Node1, buffsz)

	for i := 0; i < buffsz; i++ {
		// n0[i] = New0("x", uint64(start+i+1))
		// n0[i].Run = uint64(start + i)
		// n0[i].Win = uint64(start + i - 1)
		n0[i] = New1("x", uint64(start+i+1))
		n0[i].SetRuns(uint64(start + i))
		n0[i].SetWins(uint64(start + i - 1))

		if i >= 4999995 && i < 5000005 {
			r, _ := n0[i].Runs()
			w, _ := n0[i].Wins()
			fmt.Printf("%v -- %v / %v\n", n0[i], r, w)
		}
	}

	runtime.ReadMemStats(&mem)
	fmt.Printf("%v, %v\n", float64(mem.Alloc)/mib, float64(mem.Sys)/mib)
}

type Node0 struct {
	ID   string
	Data uint64
	Run  uint64
	Win  uint64
}

func New0(id string, data uint64) *Node0 {
	next := Node0{ID: id, Data: data, Run: 0}
	return &next
}

// ****************************************************************************************************************************************************************************
// Findings
// 1. Node1 won't work, because using a slice, even with only 1 item in the slice, take up more RAM than Node0 above!!!!!!!!!!!!!!!!!!!!!
// 2. In Golang struct empty field also take up RAM.
// 3. In Golang, 2 structs with different (uint64 change to uint16) only in a single field does not change the RAM needed for the structs
// 4. So from #2 and #3 above, memory saving only need to be done on database level, the space saving code in wdom-wc (Dat vs Data, combine Run & Win in 1 field) is useless
// ****************************************************************************************************************************************************************************

type Node1 struct {
	ID    string
	Data  uint64
	Stats []uint64 // [1] - Wins, [0] - Runs / WinsRuns
}

func New1(id string, data uint64) *Node1 {
	next := Node1{id, data, make([]uint64, 1)}
	return &next
}

// Format of Node.Wins and Node.Runs
// 1. For both runs < 4294967295, and wins < 2147483647 store in node.runs with format:
//  6666 5555 5555 5544 4444 4444 3333 3333 3322 2222 2222 1111 1111 1100 0000 0000
//  3210 9876 5432 1098 7654 3210 9876 5432 1098 7654 3210 9876 5432 1098 7654 3210
//  .... .... .... .... .... .... .... .... 1111 1111 1111 1111 1111 1111 1111 1111 Runs in node.runs
//  .111 1111 1111 1111 1111 1111 1111 1111 0000 0000 0000 0000 0000 0000 0000 0000 Wins in node.runs
//  1000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 Flag indicate node.runs is storing both values
//
// 2. Otherwise store in node.wins and node.runs respectively
//  6666 5555 5555 5544 4444 4444 3333 3333 3322 2222 2222 1111 1111 1100 0000 0000
//  3210 9876 5432 1098 7654 3210 9876 5432 1098 7654 3210 9876 5432 1098 7654 3210
//  .111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 Runs in node.runs
//  .111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 1111 Wins in node.wins

// const maxWins = 2147483647           // 31 bits integer
const maxWR = 4294967295 // 32 bits integer
// const m64Compct = 0x8000000000000000 // 1000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000 0000
const m64FdWins = 0xFFFFFFFF00000000 // 1111 1111 1111 1111 1111 1111 1111 1111 0000 0000 0000 0000 0000 0000 0000 0000
const m64IvWins = 0x00000000FFFFFFFF // 0000 0000 0000 0000 0000 0000 0000 0000 1111 1111 1111 1111 1111 1111 1111 1111
const offWins = 32
const m64FdRuns = 0x00000000FFFFFFFF // 0000 0000 0000 0000 0000 0000 0000 0000 1111 1111 1111 1111 1111 1111 1111 1111
const m64IvRuns = 0xFFFFFFFF00000000 // 1111 1111 1111 1111 1111 1111 1111 1111 0000 0000 0000 0000 0000 0000 0000 0000

// Wins read the number of wins of a node
func (n *Node1) Wins() (uint64, error) {
	switch len(n.Stats) {
	case 1:
		return (n.Stats[0] & m64FdWins) >> offWins, nil
	case 2:
		return n.Stats[1], nil
	default:
		return 0, fmt.Errorf("invalid num of statistic: %v", len(n.Stats))
	}
}

// SetWins set the number of wins of a node
func (n *Node1) SetWins(v uint64) error {
	switch len(n.Stats) {
	case 1: // wins & runs stored combined
		if v < maxWR { // new wins < 32 bits
			n.Stats[0] = n.Stats[0]&m64IvWins | ((v << offWins) & m64FdWins)
		} else {
			n.Stats = append(n.Stats, v)
			n.Stats[0] = n.Stats[0] & m64IvWins // NOTE: discard the upper 32 bits here because, originally runs is < 32 bits (stored combined)
		}
	case 2: // wins & runs already stored separately
		if v < maxWR { // new wins < 32 bits
			if n.Stats[0] < maxWR { // runs is also < 32 bits, combine wins and runs
				stat := make([]uint64, 1)
				stat[0] = n.Stats[0]&m64IvWins | ((v << offWins) & m64FdWins)
				n.Stats = stat
			} else {
				n.Stats[1] = v
			}
		} else {
			n.Stats[1] = v
		}
	default:
		return fmt.Errorf("invalid num of statistic: %v", len(n.Stats))
	}
	return nil
}

// Runs read the number of runs of a node
func (n *Node1) Runs() (uint64, error) {
	switch len(n.Stats) {
	case 1:
		return n.Stats[0] & m64FdRuns, nil
	case 2:
		return n.Stats[0], nil
	default:
		return 0, fmt.Errorf("invalid num of statistic: %v", len(n.Stats))
	}
}

// SetRuns set the number oruns of a node
func (n *Node1) SetRuns(v uint64) error {
	switch len(n.Stats) {
	case 1:
		if v < maxWR {
			n.Stats[0] = n.Stats[0]&m64IvRuns | v&m64FdRuns
		} else {
			n.Stats = append(n.Stats, (n.Stats[0]&m64IvRuns)>>offWins)
			n.Stats[0] = v // NOTE: DO NOT discard the upper 32 bits here, because the new run value is >= 32 bits!!!
		}
	case 2:
		if v < maxWR {
			if n.Stats[1] < maxWR { // wins is also < 32 bits, combine wins and runs
				stat := make([]uint64, 1)
				stat[0] = n.Stats[0]&m64IvRuns | v&m64FdRuns
				n.Stats = stat
			} else {
				n.Stats[0] = v
			}
		} else {
			n.Stats[0] = v
		}
	default:
		return fmt.Errorf("invalid num of statistic: %v", len(n.Stats))
	}
	return nil
}
