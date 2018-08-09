package diffstore

import (
	"sync"
	"time"
)

// DiffStore struct for storing diff data
type DiffStore struct {
	CurrentValue string `json:"value"`
	// Diffs        map[int64]string `json:"diffs"`
	Diffs []string `json:"diffs"`
	// Shards       map[int64]DiffShard
	lock     sync.RWMutex
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

// // DiffShard struct for storing pieces of diff data
// type DiffShard struct {
// 	CurrentValue string
// 	Diffs        map[int64]string
// 	lock         sync.RWMutex
// }
