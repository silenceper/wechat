// 运行测试：go test -race -v ./cache/ -run "TestMemory" -count=1
package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemoryGet(t *testing.T) {
	mem := NewMemory()

	err := mem.Set("username", "silenceper", 10*time.Second)
	assert.NoError(t, err)

	val := mem.Get("username")
	assert.Equal(t, "silenceper", val)

	val = mem.Get("unknown-key")
	assert.Nil(t, val)
}

func TestMemoryIsExist(t *testing.T) {
	mem := NewMemory()

	err := mem.Set("username", "silenceper", 10*time.Second)
	assert.NoError(t, err)

	assert.True(t, mem.IsExist("username"))
	assert.False(t, mem.IsExist("unknown-key"))
}

func TestMemoryDelete(t *testing.T) {
	mem := NewMemory()

	err := mem.Set("username", "silenceper", 10*time.Second)
	assert.NoError(t, err)

	err = mem.Delete("username")
	assert.NoError(t, err)

	// delete 不存在的 key 不应报错
	err = mem.Delete("unknown-key")
	assert.NoError(t, err)
}

func TestMemoryExpire(t *testing.T) {
	mem := NewMemory()

	err := mem.Set("username", "silenceper", 10*time.Millisecond)
	assert.NoError(t, err)

	assert.True(t, mem.IsExist("username"))

	time.Sleep(20 * time.Millisecond)

	assert.False(t, mem.IsExist("username"))
	assert.Nil(t, mem.Get("username"))
}

// TestMemoryConcurrentOps 验证 Get/IsExist/Set/Delete 并发执行不产生数据竞争
func TestMemoryConcurrentOps(t *testing.T) {
	mem := NewMemory()
	_ = mem.Set("key", "value", time.Minute)

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(4)
		// 并发读
		go func() {
			defer wg.Done()
			_ = mem.Get("key")
		}()
		// 并发 IsExist
		go func() {
			defer wg.Done()
			_ = mem.IsExist("key")
		}()
		// 并发写
		go func(i int) {
			defer wg.Done()
			_ = mem.Set("key", i, time.Minute)
		}(i)
		// 并发删除
		go func() {
			defer wg.Done()
			_ = mem.Delete("key")
		}()
	}
	wg.Wait()
}

// TestMemoryConcurrentExpireAndRead 验证过期惰性删除在并发场景下不会死锁
func TestMemoryConcurrentExpireAndRead(t *testing.T) {
	mem := NewMemory()

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			key := "expire-key"
			_ = mem.Set(key, i, 5*time.Millisecond)
		}(i)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Millisecond)
			// 此时 key 可能已过期，触发 deleteKey，验证不死锁
			_ = mem.Get("expire-key")
			_ = mem.IsExist("expire-key")
		}()
	}
	wg.Wait()
}
