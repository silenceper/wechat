package cache

import (
	"time"
)

// Noncache is not a cache service, do nothing just for test
type Noncache struct {
}

// NewNoncache return a new Noncahe
func NewNoncache(server ...string) *Noncache {
	return &Noncache{}
}

// Get Allways return nil
func (c *Noncache) Get(key string) interface{} {
	return nil
}

// Set Do not think it will save something for u
func (c *Noncache) Set(key string, val interface{}, timeout time.Duration) error {
	return nil
}

// IsExist Allways not found any key-value
func (c *Noncache) IsExist(key string) bool {
	return false
}

// Delete There no storage, need not delete any thing
func (c *Noncache) Delete(key string) error {
	return nil
}
