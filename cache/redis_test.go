package cache

import (
	"time"
	"testing"
)

func TestRedis(t *testing.T) {
	red := NewRedis("192.168.118.174:6379","")
	var err error
	timeoutDuration := 10 * time.Second
	if err = red.Set("username", "silenceper", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !red.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := red.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = red.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
