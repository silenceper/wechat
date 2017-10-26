package cache

import (
	"testing"
	"time"
)

func TestHostmemCache(t *testing.T) {
	mem := NewHostmemCache()
	var err error
	timeoutDuration := 3 * time.Second
	t.Logf("timeoutDuration:%d\n", timeoutDuration)
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

	if err = mem.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	time.Sleep(timeoutDuration * 3)

	if mem.IsExist("username") {
		t.Error("expired Error")
	}

}
