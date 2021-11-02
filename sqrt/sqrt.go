package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt error from square root function
// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string {
// 	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
// }

// Sqrt return squre root of the input value.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := Err("Value is -ve") //ErrNegativeSqrt(x)
		return 0, err
	}

	z := x / 2
	for l0, l1, l2 := 0.0, 0.0, 0.0; (l0 != z) && (l1 != z) && (l2 != z); z -= (z*z - x) / (2 * z) {
		l2, l1, l0 = l1, l0, z
	}
	return z, nil
}

func main() {
	vals := []float64{2.0, 4.0, 7.0, 10.0, 49.0, -5.0, 100.0, 121.0, 128.0}

	for _, v := range vals {
		r, err := Sqrt(v)
		if err != nil {
			fmt.Printf("√%v - %v\n", v, err)
		} else {
			fmt.Printf("√%v - %v (%v)\n", v, r, math.Sqrt(v))
		}
	}
}

type Err string

func (e Err) Error() string {
	return fmt.Sprintf("[game]%v", string(e))
}
