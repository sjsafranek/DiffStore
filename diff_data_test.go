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

// TestGetPreviousByIndexSuccess
func TestGetPreviousByIndexSuccess(t *testing.T) {
	// requirements
	generateTestRandomData("TestGetPreviousByIndexSuccess", 100)
	idx := int(TestDiffStore.Length() / 2)
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
	idx := int(TestDiffStore.Length() / 2)
	// benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestDiffStore.GetPreviousByIndex(idx)
	}
}

/*

Encode() ([]byte, error)

Decode(data []byte) error

*/
