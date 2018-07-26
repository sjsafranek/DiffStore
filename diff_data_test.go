package diffstore

import (
	"fmt"
	"testing"
)

var TestDiffStore DiffStore

func init() {
	TestDiffStore = New()
}

func generateTestRandomData(testname string, num int) {
	TestDiffStore = New()
	for i := 0; i < num; i++ {
		new_value := fmt.Sprintf("%v_%v", testname, i)
		TestDiffStore.Update(new_value)
	}
}

// TestUpdateAndGetCurrentSuccuss
func TestUpdateAndGetCurrentSuccuss(t *testing.T) {
	generateTestRandomData("TestUpdateAndGetCurrentSuccuss", 100)
	TestDiffStore.Update("TestUpdateAndGetCurrentSuccuss")
	if "TestUpdateAndGetCurrentSuccuss" != TestDiffStore.GetCurrent() {
		t.Error("Values do not match")
	}
}

// BenchmarkUpdate
func BenchmarkUpdate(b *testing.B) {
	generateTestRandomData("BenchmarkUpdate", 100)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		new_value := fmt.Sprintf("BenchmarkUpdate%v", i)
		TestDiffStore.Update(new_value)
	}
}

// BenchmarkGetCurrent
func BenchmarkGetCurrent(b *testing.B) {
	generateTestRandomData("BenchmarkGetCurrent", 100)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetCurrent()
	}
}

// TestGetSnapshotsSuccess
func TestGetSnapshotsSuccess(t *testing.T) {
	generateTestRandomData("TestGetSnapshotsSuccess", 100)
	snapshots := TestDiffStore.GetSnapshots()
	if 0 == len(snapshots) {
		t.Error("No timestamps found")
	}
}

// BenchmarkGetSnapshots
func BenchmarkGetSnapshots(b *testing.B) {
	generateTestRandomData("BenchmarkGetSnapshots", 100)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetSnapshots()
	}
}

// TestGetPreviousByTimestampSuccess
func TestGetPreviousByTimestampSuccess(t *testing.T) {
	generateTestRandomData("TestGetPreviousByTimestampSuccess", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx := int(len(snapshots) / 2)
	_, err := TestDiffStore.GetPreviousByTimestamp(snapshots[idx])
	if nil != err {
		t.Error(err)
	}
}

// BenchmarkGetPreviousByTimestamp
func BenchmarkGetPreviousByTimestamp(b *testing.B) {
	// requirements
	generateTestRandomData("BenchmarkGetPreviousByTimestamp", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx := int(len(snapshots) / 2)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetPreviousByTimestamp(snapshots[idx])
	}
}

// TestGetPreviousByIndexSuccess
func TestGetPreviousByIndexSuccess(t *testing.T) {
	// requirements
	generateTestRandomData("TestGetPreviousByIndexSuccess", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx := int(len(snapshots) / 2)
	// test
	_, err := TestDiffStore.GetPreviousByIndex(idx)
	if nil != err {
		t.Error(err)
	}
}

// BenchmarkGetPreviousByIndex
func BenchmarkGetPreviousByIndex(b *testing.B) {
	// requirements
	generateTestRandomData("BenchmarkGetPreviousByIndex", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx := int(len(snapshots) / 2)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetPreviousByIndex(idx)
	}
}

// TestGetPreviousWithinTimestampRangeSuccess
func TestGetPreviousWithinTimestampRangeSuccess(t *testing.T) {
	// requirements
	generateTestRandomData("TestGetPreviousWithinTimestampRangeSuccess", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx1 := int(len(snapshots) / 5)
	idx2 := int(len(snapshots) / 2)
	// test
	_, err := TestDiffStore.GetPreviousWithinTimestampRange(snapshots[idx1], snapshots[idx2])
	if nil != err {
		t.Error(err)
	}
}

// BenchmarkGetPreviousWithinTimestampRange
func BenchmarkGetPreviousWithinTimestampRange(b *testing.B) {
	// requirements
	generateTestRandomData("BenchmarkGetPreviousWithinTimestampRange", 100)
	snapshots := TestDiffStore.GetSnapshots()
	idx1 := int(len(snapshots) / 5)
	idx2 := int(len(snapshots) / 2)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetPreviousWithinTimestampRange(snapshots[idx1], snapshots[idx2])
	}
}

/*

Encode() ([]byte, error)

Decode(data []byte) error

*/
