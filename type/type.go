package main

import (
	"crypto/md5"
	"fmt"
)

const ID_LEN_MEM = 16 // ID length is 16 bytes since current implementation uses md5
const ID_LEN_SQL = 2  // Since current implementation uses md5, ID lenght is 128 bits, the IDs are stored as 2 64-bits signed int in the database
const BITS_PER_BYTE = 8
const ID_RATION = ID_LEN_MEM / ID_LEN_SQL // number of bytes can be fitted into an uint64

type NodeId [ID_LEN_SQL]int64

func CalcId(variant string, path ...uint8) (nid NodeId) {
	h := md5.New()
	fmt.Printf("x1: %v\n", h.Size())
	fmt.Printf("x2: %v\n", h.BlockSize())

	h.Write([]byte(variant))
	h.Write(path)

	inp := h.Sum(nil)
	for k := 0; k < ID_LEN_MEM; k++ {
		nid[k/ID_RATION] = nid[k/ID_RATION] | int64(inp[k])<<(((ID_LEN_MEM-k-1)%BITS_PER_BYTE)*ID_RATION)
	}
	return
}

func main() {
	var id0 NodeId
	fmt.Printf("01: %v\n", id0)

	id1 := CalcId("ABC123")
	fmt.Printf("02: %v\n", id1)
	fmt.Printf("03: %v\n", id1[0])
	fmt.Printf("04: %v\n", id1[1])
}
