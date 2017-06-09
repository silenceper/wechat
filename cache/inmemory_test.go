package cache

import (
	"testing"
	"time"
)

func TestInMemoryCache(t *testing.T) {
	testKey, testVal := "username", "sunbenxin"

	defaultExpiration := 10 * time.Minute
	timeoutDuration := 10 * time.Second
	mem := NewInMemoryCache(defaultExpiration)

	if err := mem.Set(testKey, testVal, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !mem.IsExist(testKey) {
		t.Error("IsExist Error")
	}

	name := mem.Get(testKey).(string)
	if name != testVal {
		t.Error("get Error")
	}

	if err := mem.Delete(testKey); err != nil {
		t.Error("delete Error, err=%v", err)
	}
}
