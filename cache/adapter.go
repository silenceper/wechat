package cache

import (
	"context"
	"time"

	v8redis "github.com/go-redis/redis/v8"
	v9redis "github.com/redis/go-redis/v9"
)

// redisOps 抽象不同版本 go-redis 的底层操作
type redisOps interface {
	get(ctx context.Context, key string) (string, error)
	set(ctx context.Context, key string, val interface{}, timeout time.Duration) error
	exists(ctx context.Context, key string) (int64, error)
	del(ctx context.Context, key string) error
}

// redisAdapter 基于 redisOps 实现 Cache 和 ContextCache 接口
type redisAdapter struct {
	ctx context.Context
	ops redisOps
}

// 编译时接口检查
var (
	_ Cache        = (*redisAdapter)(nil)
	_ ContextCache = (*redisAdapter)(nil)
)

// NewRedisAdapter 基于已有的 go-redis 客户端创建 Cache 适配器。
// 同时支持 go-redis v8 (github.com/go-redis/redis/v8) 和 v9 (github.com/redis/go-redis/v9)，
// 传入的 conn 参数应为对应版本的 redis.Cmdable 实现
// （如 *redis.Client、*redis.ClusterClient、redis.UniversalClient 等）。
func NewRedisAdapter(ctx context.Context, conn interface{}) ContextCache {
	var ops redisOps
	switch c := conn.(type) {
	case v8redis.Cmdable:
		ops = &v8RedisOps{conn: c}
	case v9redis.Cmdable:
		ops = &v9RedisOps{conn: c}
	default:
		panic("cache: 不支持的 Redis 客户端类型，请传入 go-redis v8 或 v9 的 Cmdable 实例")
	}
	return &redisAdapter{ctx: ctx, ops: ops}
}

// Get 获取一个值
func (a *redisAdapter) Get(key string) interface{} {
	return a.GetContext(a.ctx, key)
}

// GetContext 获取一个值
func (a *redisAdapter) GetContext(ctx context.Context, key string) interface{} {
	result, err := a.ops.get(ctx, key)
	if err != nil {
		return nil
	}
	return result
}

// Set 设置一个值
func (a *redisAdapter) Set(key string, val interface{}, timeout time.Duration) error {
	return a.SetContext(a.ctx, key, val, timeout)
}

// SetContext 设置一个值
func (a *redisAdapter) SetContext(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	return a.ops.set(ctx, key, val, timeout)
}

// IsExist 判断key是否存在
func (a *redisAdapter) IsExist(key string) bool {
	return a.IsExistContext(a.ctx, key)
}

// IsExistContext 判断key是否存在
func (a *redisAdapter) IsExistContext(ctx context.Context, key string) bool {
	result, _ := a.ops.exists(ctx, key)
	return result > 0
}

// Delete 删除
func (a *redisAdapter) Delete(key string) error {
	return a.DeleteContext(a.ctx, key)
}

// DeleteContext 删除
func (a *redisAdapter) DeleteContext(ctx context.Context, key string) error {
	return a.ops.del(ctx, key)
}

// v8RedisOps 封装 go-redis v8 的操作
type v8RedisOps struct {
	conn v8redis.Cmdable
}

func (o *v8RedisOps) get(ctx context.Context, key string) (string, error) {
	return o.conn.Get(ctx, key).Result()
}

func (o *v8RedisOps) set(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	return o.conn.SetEX(ctx, key, val, timeout).Err()
}

func (o *v8RedisOps) exists(ctx context.Context, key string) (int64, error) {
	return o.conn.Exists(ctx, key).Result()
}

func (o *v8RedisOps) del(ctx context.Context, key string) error {
	return o.conn.Del(ctx, key).Err()
}

// v9RedisOps 封装 go-redis v9 的操作
type v9RedisOps struct {
	conn v9redis.Cmdable
}

func (o *v9RedisOps) get(ctx context.Context, key string) (string, error) {
	return o.conn.Get(ctx, key).Result()
}

func (o *v9RedisOps) set(ctx context.Context, key string, val interface{}, timeout time.Duration) error {
	return o.conn.Set(ctx, key, val, timeout).Err()
}

func (o *v9RedisOps) exists(ctx context.Context, key string) (int64, error) {
	return o.conn.Exists(ctx, key).Result()
}

func (o *v9RedisOps) del(ctx context.Context, key string) error {
	return o.conn.Del(ctx, key).Err()
}
