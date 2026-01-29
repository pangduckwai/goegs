package bnch

import (
	"fmt"
	"math/bits"
	"testing"
)

func TestDivide(t *testing.T) {
	val := 6000
	fmt.Printf("TestDivide() 1: %v\n", val>>1) // /2
	fmt.Printf("TestDivide() 2: %v\n", val>>2) // /4
	fmt.Printf("TestDivide() 3: %v\n", val>>3) // /8
}

func TestRange(t *testing.T) {
	out := make([]int, 0)
	for i := range 6 {
		out = append(out, i)
	}
	fmt.Printf("TestRange() %v\n", out)
}

// max val 18446744073709551615
const U64a = 3074457345618258603 // +1
const U64b = 6148914691236517205
const U64c = 9223372036854775808 // +1
const U64d = 12297829382473034410
const U64e = 15372286728091293013 // +1

func TestMult64(t *testing.T) {
	v00, _ := bits.Mul64(U64a-10, 6)
	v01, _ := bits.Mul64(U64a+17, 6)
	v10, _ := bits.Mul64(U64b-100, 6)
	v11, _ := bits.Mul64(U64b+9000, 6)
	v20, _ := bits.Mul64(U64c-192, 6)
	v21, _ := bits.Mul64(U64c+321, 6)
	v30, _ := bits.Mul64(U64d-1234, 6)
	v31, _ := bits.Mul64(U64d+4321, 6)
	v40, _ := bits.Mul64(U64e-5555, 6)
	v41, _ := bits.Mul64(U64e+7777, 6)

	fmt.Printf("TestMult64() 00: %x\n", v00)
	fmt.Printf("TestMult64() 01: %x\n", v01)
	fmt.Printf("TestMult64() 10: %x\n", v10)
	fmt.Printf("TestMult64() 11: %x\n", v11)
	fmt.Printf("TestMult64() 20: %x\n", v20)
	fmt.Printf("TestMult64() 21: %x\n", v21)
	fmt.Printf("TestMult64() 30: %x\n", v30)
	fmt.Printf("TestMult64() 31: %x\n", v31)
	fmt.Printf("TestMult64() 40: %x\n", v40)
	fmt.Printf("TestMult64() 41: %x\n", v41)
}

// max val 4294967295
const U32a = 715827883 // +1
const U32b = 1431655765
const U32c = 2147483648 // +1
const U32d = 2863311530
const U32e = 3579139413 // +1

func TestMult32(t *testing.T) {
	v00, _ := bits.Mul32(U32a-10, 6)
	v01, _ := bits.Mul32(U32a+17, 6)
	v10, _ := bits.Mul32(U32b-100, 6)
	v11, _ := bits.Mul32(U32b+9000, 6)
	v20, _ := bits.Mul32(U32c-192, 6)
	v21, _ := bits.Mul32(U32c+321, 6)
	v30, _ := bits.Mul32(U32d-1234, 6)
	v31, _ := bits.Mul32(U32d+4321, 6)
	v40, _ := bits.Mul32(U32e-5555, 6)
	v41, _ := bits.Mul32(U32e+7777, 6)

	fmt.Printf("TestMult32() 00: %x\n", v00)
	fmt.Printf("TestMult32() 01: %x\n", v01)
	fmt.Printf("TestMult32() 10: %x\n", v10)
	fmt.Printf("TestMult32() 11: %x\n", v11)
	fmt.Printf("TestMult32() 20: %x\n", v20)
	fmt.Printf("TestMult32() 21: %x\n", v21)
	fmt.Printf("TestMult32() 30: %x\n", v30)
	fmt.Printf("TestMult32() 31: %x\n", v31)
	fmt.Printf("TestMult32() 40: %x\n", v40)
	fmt.Printf("TestMult32() 41: %x\n", v41)
}
