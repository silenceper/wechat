package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	v8redis "github.com/go-redis/redis/v8"
	v9redis "github.com/redis/go-redis/v9"
)

func setupMiniRedis(t *testing.T) *miniredis.Miniredis {
	server, err := miniredis.Run()
	if err != nil {
		t.Fatal("miniredis.Run Error", err)
	}
	t.Cleanup(server.Close)
	return server
}

func TestNewRedisAdapterV8(t *testing.T) {
	server := setupMiniRedis(t)

	rdb := v8redis.NewClient(&v8redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(context.Background(), rdb)

	key := "v8_key"
	val := "v8_value"

	if err := c.Set(key, val, 10*time.Second); err != nil {
		t.Fatal("Set Error:", err)
	}

	got := c.Get(key)
	if got == nil {
		t.Fatal("Get 返回 nil，期望有值")
	}
	if got.(string) != val {
		t.Fatalf("Get 返回 %v，期望 %v", got, val)
	}

	if !c.IsExist(key) {
		t.Fatal("IsExist 应返回 true")
	}

	if err := c.Delete(key); err != nil {
		t.Fatal("Delete Error:", err)
	}

	if c.IsExist(key) {
		t.Fatal("Delete 后 IsExist 应返回 false")
	}
}

func TestNewRedisAdapterV9(t *testing.T) {
	server := setupMiniRedis(t)

	rdb := v9redis.NewClient(&v9redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(context.Background(), rdb)

	key := "v9_key"
	val := "v9_value"

	if err := c.Set(key, val, 10*time.Second); err != nil {
		t.Fatal("Set Error:", err)
	}

	got := c.Get(key)
	if got == nil {
		t.Fatal("Get 返回 nil，期望有值")
	}
	if got.(string) != val {
		t.Fatalf("Get 返回 %v，期望 %v", got, val)
	}

	if !c.IsExist(key) {
		t.Fatal("IsExist 应返回 true")
	}

	if err := c.Delete(key); err != nil {
		t.Fatal("Delete Error:", err)
	}

	if c.IsExist(key) {
		t.Fatal("Delete 后 IsExist 应返回 false")
	}
}

func TestNewRedisAdapterContext(t *testing.T) {
	server := setupMiniRedis(t)
	ctx := context.Background()

	rdb := v8redis.NewClient(&v8redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(ctx, rdb)

	key := "ctx_key"
	val := "ctx_value"

	if err := c.SetContext(ctx, key, val, 10*time.Second); err != nil {
		t.Fatal("SetContext Error:", err)
	}

	got := c.GetContext(ctx, key)
	if got == nil || got.(string) != val {
		t.Fatalf("GetContext 返回 %v，期望 %v", got, val)
	}

	if !c.IsExistContext(ctx, key) {
		t.Fatal("IsExistContext 应返回 true")
	}

	if err := c.DeleteContext(ctx, key); err != nil {
		t.Fatal("DeleteContext Error:", err)
	}

	if c.IsExistContext(ctx, key) {
		t.Fatal("DeleteContext 后应返回 false")
	}
}

func TestNewRedisAdapterGetMissing(t *testing.T) {
	server := setupMiniRedis(t)

	rdb := v8redis.NewClient(&v8redis.Options{Addr: server.Addr()})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(context.Background(), rdb)

	got := c.Get("nonexistent")
	if got != nil {
		t.Fatalf("Get 对不存在的 key 应返回 nil，实际返回 %v", got)
	}
}

func TestNewRedisAdapterPanicOnInvalidType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("传入不支持的类型应触发 panic")
		}
	}()
	NewRedisAdapter(context.Background(), "invalid")
}

func TestNewRedisAdapterUniversalClientV8(t *testing.T) {
	server := setupMiniRedis(t)

	rdb := v8redis.NewUniversalClient(&v8redis.UniversalOptions{
		Addrs: []string{server.Addr()},
	})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(context.Background(), rdb)

	key := "uni_v8"
	val := "uni_value"

	if err := c.Set(key, val, 10*time.Second); err != nil {
		t.Fatal("Set Error:", err)
	}

	got := c.Get(key)
	if got == nil || got.(string) != val {
		t.Fatalf("UniversalClient v8: Get 返回 %v，期望 %v", got, val)
	}
}

func TestNewRedisAdapterUniversalClientV9(t *testing.T) {
	server := setupMiniRedis(t)

	rdb := v9redis.NewUniversalClient(&v9redis.UniversalOptions{
		Addrs: []string{server.Addr()},
	})
	t.Cleanup(func() { rdb.Close() })

	c := NewRedisAdapter(context.Background(), rdb)

	key := "uni_v9"
	val := "uni_value"

	if err := c.Set(key, val, 10*time.Second); err != nil {
		t.Fatal("Set Error:", err)
	}

	got := c.Get(key)
	if got == nil || got.(string) != val {
		t.Fatalf("UniversalClient v9: Get 返回 %v，期望 %v", got, val)
	}
}
