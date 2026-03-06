// Usage:
// > cd .../pkg/rand
// > go test -bench .

package rand

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark2(b *testing.B) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + int64(1)))
	b.ResetTimer()
	for b.Loop() {
		_ = rndRand(rnd, 2)
	}
}

func Benchmark4(b *testing.B) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + int64(1)))
	b.ResetTimer()
	for b.Loop() {
		_ = rndRand(rnd, 4)
	}
}

func Benchmark6(b *testing.B) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano() + int64(1)))
	b.ResetTimer()
	for b.Loop() {
		_ = rndRand(rnd, 6)
	}
}
