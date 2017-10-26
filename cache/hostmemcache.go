// Note:
//   1. Line 21 in this file, 5 seconds may less
//   2. Line 42 in memstore.go, need change to resize memory uses. (how many items do you need to store)

package cache

import (
	"encoding/json"
	"time"
)

//HostmemCache struct contains *HostmemCache.Client
type HostmemCache struct {
}

// Init clear keys
func init() {
	go func() { // watcher cleans database keys every 500 Milliseconds
		for {
			select {
			case <-time.After(5 * time.Second): // for debug == 500ms
				CleanDB()
			}
		}
	}()
}

//NewHostmemCache create new HostmemCache
func NewHostmemCache() *HostmemCache {
	return &HostmemCache{}
}

//Get return cached value
func (mem *HostmemCache) Get(key string) interface{} {
	item, ok := DataBase[key]
	if !ok {
		return nil
	}

	var result interface{}
	if err := json.Unmarshal(item.Val, &result); err != nil {
		return nil
	}
	return result
}

// IsExist check value exists in HostmemCache.
func (mem *HostmemCache) IsExist(key string) bool {
	_, ok := DataBase[key]
	return ok
}

//Set cached value with key and expire time.
func (mem *HostmemCache) Set(key string, val interface{}, timeout time.Duration) (err error) {
	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return err
	}

	DataBase[key] = &DBItem{time.Now().Add(timeout).Unix(), TList, data}
	return nil
}

//Delete delete value in HostmemCache.
func (mem *HostmemCache) Delete(key string) error {
	delete(DataBase, key)
	return nil
}
