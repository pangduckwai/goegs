package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// Node test
type Node struct {
	ID     string
	Move   uint32 // (pkey) indicate progress of the game
	Turn   uint8  // (pkey) identify the player
	Action uint8  // (pkey) Actions a player may takes
	Index1 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Index2 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Index3 uint8  // (pkey) indices of players, territories or cards (Depends on the event)
	Value1 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Value2 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Value3 uint32 // (pkey) Event payload, mainly troop amount, etc (Depends on the event)
	Runs   uint64 // Total number of runs
	Wins   uint64 // Total number of won games of this player
	Parent *Node
	Next   []*Node
}

func hash(n *Node) string {
	a := make([]byte, 4)
	b := make([]byte, 4)
	c := make([]byte, 4)
	d := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, n.Move)
	binary.LittleEndian.PutUint32(b, n.Value1)
	binary.LittleEndian.PutUint32(c, n.Value2)
	binary.LittleEndian.PutUint32(d, n.Value3)
	return fmt.Sprintf("%x\n", sha1.Sum([]byte{
		a[0], a[1], a[2], a[3], n.Turn, n.Action, n.Index1, n.Index2, n.Index3, b[0], b[1], b[2], b[3], c[0], c[1], c[2], c[3], d[0], c[1], c[2], c[3],
	}))
}

func main() {
	fmt.Println(hash(&Node{}))
}
