// Usage:
// > cd .../pkg/fast
// > go test -bench .

package fast

import (
	"testing"

	"sea9.org/go/egs/randBench/pkg/bytedance/fastrand"
)

func Benchmark2(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_ = fastRand(2)
	}
}

func Benchmark4(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_ = fastRand(4)
	}
}

func Benchmark6(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_ = fastRand(6)
	}
}

func BenchmarkIntn(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_ = fastrand.Intn(6)
	}
}

func BenchmarkUint32n(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		fastrand.Uint32n(uint32(6))
	}
}
