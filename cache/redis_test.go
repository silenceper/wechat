package cache

import (
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host: "127.0.0.1:6379",
	}
	redis := NewRedis(opts)
	var err error
	timeoutDuration := 1 * time.Second

	if err = redis.Set("username", "machao520r", timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	name := redis.Get("username").(string)
	if name != "machao520r" {
		t.Error("get Error")
	}

	if err = redis.Delete("username"); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
