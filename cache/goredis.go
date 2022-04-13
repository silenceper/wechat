package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// GoRedis .redis cache
type GoRedis struct {
	ctx  context.Context
	conn redis.UniversalClient
}

// NewGoRedis 实例化
func NewGoRedis(ctx context.Context, opts *RedisOpts) *GoRedis {
	conn := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        []string{opts.Host},
		DB:           opts.Database,
		Password:     opts.Password,
		IdleTimeout:  time.Second * time.Duration(opts.IdleTimeout),
		MinIdleConns: opts.MaxIdle,
	})
	return &GoRedis{ctx: ctx, conn: conn}
}

// SetConn 设置conn
func (r *GoRedis) SetConn(conn redis.UniversalClient) {
	r.conn = conn
}

// SetRedisCtx 设置redis ctx 参数
func (r *GoRedis) SetRedisCtx(ctx context.Context) {
	r.ctx = ctx
}

// Get 获取一个值
func (r *GoRedis) Get(key string) interface{} {
	result, err := r.conn.Do(r.ctx, "GET", key).Result()
	if err != nil {
		return nil
	}
	return result
}

// Set 设置一个值
func (r *GoRedis) Set(key string, val interface{}, timeout time.Duration) error {
	return r.conn.SetEX(r.ctx, key, val, timeout).Err()
}

// IsExist 判断key是否存在
func (r *GoRedis) IsExist(key string) bool {
	result, _ := r.conn.Exists(r.ctx, key).Result()

	return result > 0
}

// Delete 删除
func (r *GoRedis) Delete(key string) error {
	return r.conn.Del(r.ctx, key).Err()
}
