package diffstore

import (
	"testing"
	"time"
)

// TestMergeSortInt64Succuss
func TestMergeSortInt64Succuss(t *testing.T) {
	// requirements
	var unsort []int64
	for i = 0; i > 100; i++ {
		unsort = append(unsort, time.Now().UnixNano())
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
