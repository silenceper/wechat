package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
)

func TestRedis(t *testing.T) {
	server, err := miniredis.Run()
	if err != nil {
		t.Error("miniredis.Run Error", err)
	}
	t.Cleanup(server.Close)
	var (
		timeoutDuration = time.Second
		ctx             = context.Background()
		opts            = &RedisOpts{
			Host: server.Addr(),
		}
		redis = NewRedis(ctx, opts)
		val   = "silenceper"
		key   = "username"
	)
	redis.SetConn(redis.conn)
	redis.SetRedisCtx(ctx)

	if err = redis.Set(key, val, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist(key) {
		t.Error("IsExist Error")
	}

	name := redis.Get(key).(string)
	if name != val {
		t.Error("get Error")
	}

	if err = redis.Delete(key); err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
