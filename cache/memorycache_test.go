package cache

import (
	"testing"
	"time"
)

func TestMemorycache(t *testing.T) {
	mem := NewMemoryCache()
	var err error
	timeoutDuration := 10 * time.Second
	if err = mem.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !mem.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := mem.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = mem.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
