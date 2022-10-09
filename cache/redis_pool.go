package cache

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

// RedisPool rediscache 池
type RedisPool struct {
	sync.Pool
}

// NewRedisPool 实例化redis cache 池
func NewRedisPool(conn redis.UniversalClient) *RedisPool {
	return &RedisPool{
		Pool: sync.Pool{
			New: func() interface{} {
				redisCache := &Redis{ctx: context.TODO(), conn: conn}
				return redisCache
			},
		},
	}
}

// Borrow 传入 context.Context, 从池中借走一个 Redis cache
func (p *RedisPool) Borrow(ctx context.Context) *Redis {
	redisCache, _ := p.Pool.Get().(*Redis)
	redisCache.SetRedisCtx(ctx)
	return redisCache
}

// Return 归还一个 Redis cache 到池中
func (p *RedisPool) Return(redisCache *Redis) {
	redisCache.SetRedisCtx(context.TODO())
	p.Pool.Put(redisCache)
}
