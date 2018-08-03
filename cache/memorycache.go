package cache

import (
	"time"
	"sync"
)
//MemoryCache struct
type MemoryCache struct {
	mu                sync.RWMutex
	Data 			  map[string]Item
	gcInterval        time.Duration   //gc周期
	stopGc            chan bool       //停止gc管道标识
}

// chche item
type Item struct {
	Object 			interface{}
	Expiration  	int64

}
//new memory cache
func NewMemoryCache() *MemoryCache {
	data := make(map[string]Item)
	gcInterval, _ := time.ParseDuration("3s")
	mem := &MemoryCache{Data:data,gcInterval:gcInterval}
	go mem.gcLoop()
	return mem
}

//Get return cached value
func (mem *MemoryCache) Get(key string) interface{} {
	mem.mu.RLock()
	defer mem.mu.RUnlock()
	item, ok := mem.Data[key]
	if  !ok {
		return nil
	}
	return item.Object
}

// IsExist check value exists in memorycache.
func (mem *MemoryCache) IsExist(key string) bool {
	if _, ok := mem.Data[key]; !ok {
		return false
	}
	return true
}

//Set cached value with key and expire time.
func (mem *MemoryCache) Set(key string, val interface{}, timeout time.Duration)(error) {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	var e =  time.Now().Add(timeout).UnixNano()
	mem.Data[key] = Item{Object:val,Expiration:e}
	return nil
}

//Delete delete value in memorycache.
func (mem *MemoryCache) Delete(key string)(error) {
	delete(mem.Data, key)
	return nil
}
//check isExpired
func (item Item) IsExpired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration //如果当前时间超则过期
}

//循环gc
func (mem *MemoryCache) gcLoop() {
	ticker := time.NewTicker(mem.gcInterval) //初始化一个定时器
	for {
		select {
		case <-ticker.C:
			mem.clearItems()
		case <-mem.stopGc:
			ticker.Stop()
			return
		}
	}
}

// remove keys
func (mem *MemoryCache) clearItems() {
	mem.mu.Lock()
	defer mem.mu.Unlock()
	for key, itm := range mem.Data {
		if itm.IsExpired() {
			mem.Delete(key)
		}
	}
}
