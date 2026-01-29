// Usage:
// > cd .../pkg/ver2
// > go test -bench .

package ver2

import (
	"math/rand/v2"
	"testing"
	"time"
)

func Benchmark2(b *testing.B) {
	rnd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v2Rand(rnd, 2)
	}
}

func Benchmark4(b *testing.B) {
	rnd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v2Rand(rnd, 4)
	}
}

func Benchmark6(b *testing.B) {
	rnd := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(1)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v2Rand(rnd, 6)
	}
}
