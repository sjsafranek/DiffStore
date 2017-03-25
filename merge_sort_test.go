package diffstore

import (
	"math/rand"
	"testing"
	"time"
)

func randomInt64Between(min, max int) int64 {
	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	// Note that this is not safe to use for random numbers you
	// intend to be secret, use `crypto/rand` for those.
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min) + min)
}

// TestMergeSortInt64Succuss
func TestMergeSortInt64Succuss(t *testing.T) {
	// requirements
	var unsort []int64
	for i := 0; i < 100; i++ {
		num := time.Now().UnixNano()
		num += randomInt64Between(0, 1000000000000000000)
		unsort = append(unsort, num)
	}

	// test
	sorted := MergeSortInt64(unsort)
	var previous int64
	for i := range sorted {
		if 0 != i {
			if sorted[i] < previous {
				t.Error("int64 not sorted")
			}
		}
		previous = sorted[i]
	}
}

// BenchmarkMergeSortInt64100
func BenchmarkMergeSortInt64100(b *testing.B) {
	// requirements
	var unsort []int64
	for i := 0; i < 100; i++ {
		num := time.Now().UnixNano()
		num += randomInt64Between(0, 1000000000000000000)
		unsort = append(unsort, num)
	}
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortInt64(unsort)
	}
}

// BenchmarkMergeSortInt641000
func BenchmarkMergeSortInt641000(b *testing.B) {
	// requirements
	var unsort []int64
	for i := 0; i < 1000; i++ {
		num := time.Now().UnixNano()
		num += randomInt64Between(0, 1000000000000000000)
		unsort = append(unsort, num)
	}
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortInt64(unsort)
	}
}

// BenchmarkMergeSortInt6410000
func BenchmarkMergeSortInt6410000(b *testing.B) {
	// requirements
	var unsort []int64
	for i := 0; i < 10000; i++ {
		num := time.Now().UnixNano()
		num += randomInt64Between(0, 1000000000000000000)
		unsort = append(unsort, num)
	}
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortInt64(unsort)
	}
}

// BenchmarkMergeSortInt64100000
func BenchmarkMergeSortInt64100000(b *testing.B) {
	// requirements
	var unsort []int64
	for i := 0; i < 100000; i++ {
		num := time.Now().UnixNano()
		num += randomInt64Between(0, 1000000000000000000)
		unsort = append(unsort, num)
	}
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortInt64(unsort)
	}
}
