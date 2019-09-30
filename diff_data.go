package diffstore

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// NewDiffStore creates and returns DiffStore struct
func New() DiffStore {
	var ddata DiffStore
	ddata.CurrentValue = ""
	ddata.Diffs = []string{}
	ddata.CreateAt = time.Now()
	return ddata
}

// Encode marshals struct into json.
func (self *DiffStore) Encode() ([]byte, error) {
	enc, err := json.Marshal(self)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

// Decode unmarshals struct from json.
func (self *DiffStore) Decode(data []byte) error {
	err := json.Unmarshal(data, &self)
	if err != nil {
		return err
	}
	return nil
}

// diffRebuildtexts rebuilds text value from list of diff changes.
func (self *DiffStore) diffRebuildtexts(diffs []diffmatchpatch.Diff) []string {
	text := []string{"", ""}
	for _, diff := range diffs {
		if diff.Type != diffmatchpatch.DiffInsert {
			text[0] += diff.Text
		}
		if diff.Type != diffmatchpatch.DiffDelete {
			text[1] += diff.Text
		}
	}
	return text
}

// rebuildTextsToDiffN rebuilds text value until the supplied update index is reached
func (self *DiffStore) rebuildTextsToDiffN(index int) (string, error) {
	dmp := diffmatchpatch.New()
	lastText := ""
	self.lock.Lock()
	defer self.lock.Unlock()

	for i, diff := range self.Diffs {
		seq1, _ := dmp.DiffFromDelta(lastText, diff)
		textsLinemode := self.diffRebuildtexts(seq1)
		rebuilt := textsLinemode[len(textsLinemode)-1]

		if i == index {
			return rebuilt, nil
		}
		lastText = rebuilt
	}

	return "", fmt.Errorf("Could not rebuild from diffs")
}

// Update adds new text change.
func (self *DiffStore) Update(newText string) {

	// check for changes
	if self.GetCurrent() == newText {
		return
	}

	self.lock.RLock()
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(self.CurrentValue, newText, true)
	delta := dmp.DiffToDelta(diffs)
	self.CurrentValue = newText

	if nil == self.Diffs {
		self.Diffs = []string{}
	}

	self.Diffs = append(self.Diffs, delta)

	self.UpdateAt = time.Now()
	self.lock.RUnlock()
}

// GetCurrent returns current text value.
func (self *DiffStore) GetCurrent() string {
	return self.CurrentValue
}

// GetPreviousByIndex builds and returns previous text value by update index
func (self *DiffStore) GetPreviousByIndex(idx int) (string, error) {
	// check inputs
	if 0 > idx {
		return "", fmt.Errorf("Index must be positive integer")
	}

	// use timestamp to find value
	oldValue, err := self.rebuildTextsToDiffN(idx)
	return oldValue, err
}

// Length returns length of text updates
func (self *DiffStore) Length() int {
	return len(self.Diffs)
}
