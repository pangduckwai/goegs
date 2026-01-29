// Usage:
// > cd .../pkg/fast
// > go test -bench .

package fast

import "testing"

func Benchmark2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fastRand(2)
	}
}

func Benchmark4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fastRand(4)
	}
}

func Benchmark6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fastRand(6)
	}
}
