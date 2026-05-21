package cache

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
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
			Host:         server.Addr(),
			Password:     "",
			Database:     0,
			PoolSize:     10,
			MinIdleConns: 5,
			DialTimeout:  5,
			ReadTimeout:  5,
			WriteTimeout: 5,
			PoolTimeout:  5,
			IdleTimeout:  300,
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

// setupRedisServer 创建并返回一个 miniredis 服务器实例
func setupRedisServer(t *testing.T) *miniredis.Miniredis {
	server, err := miniredis.Run()
	if err != nil {
		t.Fatal("miniredis.Run Error", err)
	}
	t.Cleanup(server.Close)
	return server
}

// TestRedisMaxIdleMapping 测试只设置MaxIdle应该映射到MinIdleConns
func TestRedisMaxIdleMapping(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:     server.Addr(),
		Database: 0,
		MaxIdle:  10,
	}
	r := NewRedis(ctx, opts)

	// 获取底层的 UniversalClient 并断言为 *redis.Client
	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	// 注意：MinIdleConns 表示期望的最小空闲连接数，但实际空闲连接数可能不同
	// 我们需要通过 Options() 来验证配置是否正确应用
	clientOpts := client.Options()
	if clientOpts.MinIdleConns != 10 {
		t.Errorf("期望 MinIdleConns = 10, 实际 = %d", clientOpts.MinIdleConns)
	}
}

// TestRedisMaxActiveMapping 测试只设置MaxActive应该映射到PoolSize
func TestRedisMaxActiveMapping(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:      server.Addr(),
		Database:  0,
		MaxActive: 20,
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	if clientOpts.PoolSize != 20 {
		t.Errorf("期望 PoolSize = 20, 实际 = %d", clientOpts.PoolSize)
	}
}

// TestRedisNewFieldsPriority 测试新字段应该优先于旧字段
func TestRedisNewFieldsPriority(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:         server.Addr(),
		Database:     0,
		MaxIdle:      5,
		MinIdleConns: 15,
		MaxActive:    10,
		PoolSize:     30,
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	if clientOpts.MinIdleConns != 15 {
		t.Errorf("期望 MinIdleConns = 15 (新字段优先), 实际 = %d", clientOpts.MinIdleConns)
	}
	if clientOpts.PoolSize != 30 {
		t.Errorf("期望 PoolSize = 30 (新字段优先), 实际 = %d", clientOpts.PoolSize)
	}
}

// TestRedisPositiveTimeouts 测试正值超时应该正确应用
func TestRedisPositiveTimeouts(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:         server.Addr(),
		Database:     0,
		DialTimeout:  10,
		ReadTimeout:  20,
		WriteTimeout: 30,
		PoolTimeout:  40,
		IdleTimeout:  50,
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	if clientOpts.DialTimeout != 10*time.Second {
		t.Errorf("期望 DialTimeout = 10s, 实际 = %v", clientOpts.DialTimeout)
	}
	if clientOpts.ReadTimeout != 20*time.Second {
		t.Errorf("期望 ReadTimeout = 20s, 实际 = %v", clientOpts.ReadTimeout)
	}
	if clientOpts.WriteTimeout != 30*time.Second {
		t.Errorf("期望 WriteTimeout = 30s, 实际 = %v", clientOpts.WriteTimeout)
	}
	if clientOpts.PoolTimeout != 40*time.Second {
		t.Errorf("期望 PoolTimeout = 40s, 实际 = %v", clientOpts.PoolTimeout)
	}
	if clientOpts.IdleTimeout != 50*time.Second {
		t.Errorf("期望 IdleTimeout = 50s, 实际 = %v", clientOpts.IdleTimeout)
	}
}

// TestRedisNegativeTimeouts 测试-1值应该禁用超时
func TestRedisNegativeTimeouts(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:         server.Addr(),
		Database:     0,
		DialTimeout:  -1,
		ReadTimeout:  -1,
		WriteTimeout: -1,
		PoolTimeout:  -1,
		IdleTimeout:  -1,
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	// -1 应该被设置为负值表示禁用超时
	// DialTimeout, PoolTimeout, IdleTimeout 会被设置为 -1ns
	if clientOpts.DialTimeout != -1 {
		t.Errorf("期望 DialTimeout = -1ns (禁用), 实际 = %v", clientOpts.DialTimeout)
	}
	// ReadTimeout 和 WriteTimeout 在 go-redis 中有特殊处理
	// 当设置为负值时，会被规范化为 0，这也表示无超时
	t.Logf("ReadTimeout = %v (设置为-1后的值)", clientOpts.ReadTimeout)
	t.Logf("WriteTimeout = %v (设置为-1后的值)", clientOpts.WriteTimeout)

	if clientOpts.PoolTimeout != -1 {
		t.Errorf("期望 PoolTimeout = -1ns (禁用), 实际 = %v", clientOpts.PoolTimeout)
	}
	if clientOpts.IdleTimeout != -1 {
		t.Errorf("期望 IdleTimeout = -1ns (禁用), 实际 = %v", clientOpts.IdleTimeout)
	}
}

// TestRedisZeroTimeouts 测试0值应该使用go-redis默认值
func TestRedisZeroTimeouts(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:         server.Addr(),
		Database:     0,
		DialTimeout:  0,
		ReadTimeout:  0,
		WriteTimeout: 0,
		PoolTimeout:  0,
		IdleTimeout:  0,
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	// 0值应该保持为0，由 go-redis 使用默认值
	// go-redis 的默认值：
	// DialTimeout: 5s
	// ReadTimeout: 3s
	// WriteTimeout: ReadTimeout
	// PoolTimeout: ReadTimeout + 1s
	// IdleTimeout: 5min

	if clientOpts.DialTimeout == 0 {
		t.Error("期望 DialTimeout 使用 go-redis 默认值 (5s), 实际为 0")
	}
	if clientOpts.ReadTimeout == 0 {
		t.Error("期望 ReadTimeout 使用 go-redis 默认值 (3s), 实际为 0")
	}
	if clientOpts.WriteTimeout == 0 {
		t.Error("期望 WriteTimeout 使用 go-redis 默认值 (ReadTimeout), 实际为 0")
	}
	if clientOpts.PoolTimeout == 0 {
		t.Error("期望 PoolTimeout 使用 go-redis 默认值 (ReadTimeout + 1s), 实际为 0")
	}
	if clientOpts.IdleTimeout == 0 {
		t.Error("期望 IdleTimeout 使用 go-redis 默认值 (5min), 实际为 0")
	}
}

// TestRedisMixedTimeouts 测试混合超时配置
func TestRedisMixedTimeouts(t *testing.T) {
	server := setupRedisServer(t)
	ctx := context.Background()

	opts := &RedisOpts{
		Host:         server.Addr(),
		Database:     0,
		DialTimeout:  5,  // 正值
		ReadTimeout:  -1, // 禁用
		WriteTimeout: 0,  // 使用默认值
		PoolTimeout:  10, // 正值
		IdleTimeout:  -1, // 禁用
	}
	r := NewRedis(ctx, opts)

	client, ok := r.conn.(*redis.Client)
	if !ok {
		t.Fatal("无法转换为 *redis.Client")
	}

	clientOpts := client.Options()
	if clientOpts.DialTimeout != 5*time.Second {
		t.Errorf("期望 DialTimeout = 5s, 实际 = %v", clientOpts.DialTimeout)
	}
	// ReadTimeout 设置为 -1，会被 go-redis 处理为 0（无超时）
	t.Logf("ReadTimeout = %v (设置为-1后的值)", clientOpts.ReadTimeout)

	// WriteTimeout 设置为 0，应该使用 go-redis 的默认值
	// 默认值通常是 ReadTimeout 的值
	t.Logf("WriteTimeout = %v (设置为0后使用的默认值)", clientOpts.WriteTimeout)

	if clientOpts.PoolTimeout != 10*time.Second {
		t.Errorf("期望 PoolTimeout = 10s, 实际 = %v", clientOpts.PoolTimeout)
	}

	// IdleTimeout 设置为 -1，应该被设置为 -1ns（禁用空闲超时）
	if clientOpts.IdleTimeout != -1 {
		t.Errorf("期望 IdleTimeout = -1ns (禁用), 实际 = %v", clientOpts.IdleTimeout)
	}
}
