package main

import "fmt"

// MyReader test reader
type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (n int, err error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 17)
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
}
