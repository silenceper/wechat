package cache

import (
	"errors"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

//Memcache struct contains *memcache.Client
type Memcache struct {
	conn *memcache.Client
}

//NewMemcache create new memcache
func NewMemcache(server ...string) *Memcache {
	mc := memcache.New(server...)
	return &Memcache{mc}
}

//Get return cached value
func (mem *Memcache) Get(key string) interface{} {
	if item, err := mem.conn.Get(key); err == nil {
		return string(item.Value)
	}
	return nil
}

// IsExist check value exists in memcache.
func (mem *Memcache) IsExist(key string) bool {
	_, err := mem.conn.Get(key)
	if err != nil {
		return false
	}
	return true
}

//Set cached value with key and expire time.
func (mem *Memcache) Set(key string, val interface{}, timeout time.Duration) error {
	v, ok := val.(string)
	if !ok {
		return errors.New("val must string")
	}
	item := &memcache.Item{Key: key, Value: []byte(v), Expiration: int32(timeout / time.Second)}
	return mem.conn.Set(item)
}

//Delete delete value in memcache.
func (mem *Memcache) Delete(key string) error {
	return mem.conn.Delete(key)
}
