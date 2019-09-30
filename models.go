package diffstore

import (
	"sync"
	"time"
)

// DiffStore struct for storing diff data
type DiffStore struct {
	CurrentValue string   `json:"value"`
	Diffs        []string `json:"diffs"`
	lock         sync.RWMutex
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}
