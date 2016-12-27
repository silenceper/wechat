package cache

import (
	"testing"
	"time"
)

func TestRedies(t *testing.T) {
	// create redis client defalut port is 6379
	redis, err := NewRediscache("127.0.0.1", uint(6379))
	if err != nil {
		t.Error("connect Error", err)
	}

	if err = redis.Set("username", "wechat", 10*time.Second); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "silenceper" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
