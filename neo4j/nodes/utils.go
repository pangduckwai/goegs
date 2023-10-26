package nodes

import (
	"fmt"
	"strings"
)

///////////////////////////////////////////////////////////////////
//  3322 2222 2222 1111 1111 1100 0000 0000
//  1098 7654 3210 9876 5432 1098 7654 3210
//  .... .... .... .... .... .... .... .111 Turn     0x00000007
//  ..11 1100 0000 0000 0000 0000 0000 0000 Action   0x3C000000
//  .... ..11 1111 0000 0000 0000 0000 0000 Value1   0x03F00000
//  .... .... .... 1111 1100 0000 0000 0000 Value2   0x000FC000
//  .... .... .... .... ..11 1111 0000 0000 Value3   0x00003F00
//  .... .... .... .... .... .... 1100 0000 Value4   0x000000C0 (Dice R)
//  .... .... .... .... .... .... ..11 0000 Value5   0x00000030 (Dice W)
//  .... .... .... .... .... .... .... 1000 End game 0x00000008
//  .100 0000 0000 0000 0000 0000 0000 0000 Current (transient)
//  1000 0000 0000 0000 0000 0000 0000 0000 Busy (transient)

const mskPsis = 0x3FFFFFFF // 0011 1111 1111 1111 1111 1111 1111 1111 Mask for persisting value (ignore transient ones)
const mFdBusy = 0x80000000
const mIvBusy = 0x7FFFFFFF // 0111 1111 1111 1111 1111 1111 1111 1111
const mFdCurr = 0x40000000
const mIvCurr = 0xBFFFFFFF // 1011 1111 1111 1111 1111 1111 1111 1111
const mFdGame = 0x00000008 // .... .... .... .... .... .... .... 1000
const mIvGame = 0xFFFFFFF7 // 1111 1111 1111 1111 1111 1111 1111 0111
const mFdTurn = 0x00000007 // .... .... .... .... .... .... .... .111
const mIvTurn = 0xFFFFFFF8 // 1111 1111 1111 1111 1111 1111 1111 1000
const mFdActn = 0x3C000000 // ..11 1100 0000 0000 0000 0000 0000 0000
const mIvActn = 0xC3FFFFFF // 1100 0011 1111 1111 1111 1111 1111 1111
const offActn = 26
const mFdVal1 = 0x03F00000 // .... ..11 1111 0000 0000 0000 0000 0000
const mIvVal1 = 0xFC0FFFFF // 1111 1100 0000 1111 1111 1111 1111 1111
const offVal1 = 20
const mFdVal2 = 0x000FC000 // .... .... .... 1111 1100 0000 0000 0000
const mIvVal2 = 0xFFF03FFF // 1111 1111 1111 0000 0011 1111 1111 1111
const offVal2 = 14
const mFdVal3 = 0x00003F00 // .... .... .... .... ..11 1111 0000 0000
const mIvVal3 = 0xFFFFC0FF // 1111 1111 1111 1111 1100 0000 1111 1111
const offVal3 = 8
const mFdVal4 = 0x000000C0 // .... .... .... .... .... .... 1100 0000
const mIvVal4 = 0xFFFFFF3F // 1111 1111 1111 1111 1111 1111 0011 1111
const offVal4 = 6
const mFdVal5 = 0x00000030 // .... .... .... .... .... .... ..11 0000
const mIvVal5 = 0xFFFFFFCF // 1111 1111 1111 1111 1111 1111 1100 1111
const offVal5 = 4

// ClearTransient clear transient fields
func ClearTransient(data uint32) uint32 {
	return data & mskPsis
}

// Busy check if node is busy
func Busy(data uint32) bool {
	return ((data & mFdBusy) != 0)
}

// SetBusy mark the node as busy
func SetBusy(data uint32) uint32 {
	return data | mFdBusy
}

// ClearBusy clear the busy flag
func ClearBusy(data uint32) uint32 {
	return data & mIvBusy
}

// Current check if node is current or not
func Current(data uint32) bool {
	return ((data & mFdCurr) != 0)
}

// SetCurrent mark the node as current
func SetCurrent(data uint32) uint32 {
	return data | mFdCurr
}

// ClearCurrent clear the current flag
func ClearCurrent(data uint32) uint32 {
	return data & mIvCurr
}

// EndGame indicate a game ending node
func EndGame(data uint32) bool {
	return ((data & mFdGame) != 0)
}

// SetEndGame mark the node as game ending
func SetEndGame(data uint32) uint32 {
	return data | mFdGame
}

// ClearEndGame clear the end-game flag
func ClearEndGame(data uint32) uint32 {
	return data & mIvGame
}

// Turn get turn value
func Turn(data uint32) uint8 {
	return uint8(data & mFdTurn)
}

// SetTurn set turn value
func SetTurn(data uint32, t uint8) uint32 {
	return data&mIvTurn | uint32(t)&mFdTurn
}

// Action get action value
func Action(data uint32) uint8 {
	return uint8((data & mFdActn) >> offActn)
}

// SetAction set action value
func SetAction(data uint32, a uint8) uint32 {
	return data&mIvActn | ((uint32(a) << offActn) & mFdActn)
}

// Param get short value
func Param(data uint32, i int) uint8 {
	switch i {
	case 1:
		return uint8((data & mFdVal1) >> offVal1)
	case 2:
		return uint8((data & mFdVal2) >> offVal2)
	case 3:
		return uint8((data & mFdVal3) >> offVal3)
	case 4:
		return uint8((data & mFdVal4) >> offVal4)
	case 5:
		return uint8((data & mFdVal5) >> offVal5)
	default:
		return uint8(15)
	}
}

// SetParam set short value
func SetParam(data uint32, i int, v uint8) uint32 {
	switch i {
	case 1:
		return data&mIvVal1 | ((uint32(v) << offVal1) & mFdVal1)
	case 2:
		return data&mIvVal2 | ((uint32(v) << offVal2) & mFdVal2)
	case 3:
		return data&mIvVal3 | ((uint32(v) << offVal3) & mFdVal3)
	case 4:
		return data&mIvVal4 | ((uint32(v) << offVal4) & mFdVal4)
	case 5:
		return data&mIvVal5 | ((uint32(v) << offVal5) & mFdVal5)
	default:
		return 0xFFFFFFFF
	}
}

const dvFlag = 0x8 // 1000

// HasDerived child nodes are derived nodes
// Node has derived child when action equals 'TurnEnded', 'PositionFortified', 'TerritoryAttacked' and 'TerritoryDefended'
func HasDerived(d uint32) bool {
	return ((Action(d) & dvFlag) != 0)
}

// UseUcb next node selected using UCB
func UseUcb(d uint32) bool {
	return (Action(d) <= 9)
}

// ToString return a string representation of a Node. The first part is the node ID, 'p' <space> characters are
// added after the node ID and before the rest of the contents.
func ToString(
	d uint32, // node data
	v uint16, // node value
	w uint64, // number of wins
	r uint64, // number of runs
	p int, // padding size
	n bool, // true if parent is nil
	cnt int, // number of child nodes
	lvl int, // levels
) string {
	pad := strings.Repeat(" ", p)
	val := fmt.Sprintf("%2v,%2v,%2v,%v,%v,%4v", Param(d, 1), Param(d, 2), Param(d, 3), Param(d, 4), Param(d, 5), v)
	rt := "±"
	if n || (Current(d) && EndGame(d)) {
		rt = "≡"
	} else if Current(d) {
		rt = "»"
	} else if EndGame(d) {
		rt = "«"
	}

	if lvl > 0 { // Borrow node.Level (a transient field) to store the number of child nodes during tracking
		return fmt.Sprintf(
			"%v[%11v/%11v]%d{%x}%-17v|%4v%v %vL",
			pad, w, r, Turn(d), Action(d), val, cnt, rt, lvl,
		)
	}
	return fmt.Sprintf(
		"%v[%11v/%11v]%d{%x}%-17v|%4v%v TEMP: {r:1, d:%v, v:%v}", // TODO TEMP
		pad, w, r, Turn(d), Action(d), val, cnt, rt, d, v,
	)
}

// Same chekc if the data of 2 nodes are the same
func Same(d0 uint32, v0 uint16, d1 uint32, v1 uint16) bool {
	return (d0&mskPsis) == (d1&mskPsis) && (v0 == v1)
}

// Less for sorting the child array in Node
func Less(dn uint32, vn uint16, dl uint32, vl uint16) bool {
	if dn < dl {
		return true
	} else if dn > dl {
		return false
	}

	return vn < vl
}

// Less for sorting the child array in Node
// func Less(dn uint32, vn uint16, dl uint32, vl uint16) bool {
// 	ai := Action(dn)
// 	aj := Action(dl)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	ai = Param(dn, 1)
// 	aj = Param(dl, 1)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	ai = Param(dn, 2)
// 	aj = Param(dl, 2)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	ai = Param(dn, 3)
// 	aj = Param(dl, 3)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	ai = Param(dn, 4)
// 	aj = Param(dl, 4)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	ai = Param(dn, 5)
// 	aj = Param(dl, 5)
// 	if ai < aj {
// 		return true
// 	} else if ai > aj {
// 		return false
// 	}

// 	return vn < vl
// }
