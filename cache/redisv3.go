package cache

import (
	"time"

	redis "gopkg.in/redis.v3"
)

/*
	使用方式:
		1. 通过NewRdsV3Client建立连接后使用
		2. 通过RdsV3ClientInit初始化获得对象(与外部服务共用连接)
*/

// RedisV3 client
type RedisV3 struct {
	rdsClient *redis.Client
}

// RedisConfig redis连接配置
type RedisConfig struct {
	Address     string `yml:"address" json:"address"`
	Password    string `yml:"password" json:"password"`
	Timeout     int    `yml:"timeout" json:"timeout"`
	MaxIdle     int    `yml:"max_idle" json:"max_idle"`
	IdleTimeout int    `yml:"idle_timeout" json:"idle_timeout"`
}

// 初始化redisclient
func NewRdsV3Client(c *RedisConfig) *RedisV3 {
	client := redis.NewClient(&redis.Options{
		Addr:         c.Address,
		Password:     c.Password,
		PoolSize:     c.MaxIdle,
		IdleTimeout:  time.Duration(c.IdleTimeout) * time.Second,
		DialTimeout:  time.Duration(c.Timeout) * time.Second,
		ReadTimeout:  time.Duration(c.Timeout) * time.Second,
		WriteTimeout: time.Duration(c.Timeout) * time.Second,
	})
	return &RedisV3{client}
}

// RdsV3ClientInit 外部服务初始化(直接传入已初始化的redis client)
func RdsV3ClientInit(client *redis.Client) *RedisV3 {
	return &RedisV3{
		rdsClient: client,
	}
}

// Get
func (r *RedisV3) Get(key string) interface{} {
	return r.rdsClient.Get(key).Val()
}

// Set
func (r *RedisV3) Set(key string, val interface{}, duration time.Duration) (err error) {
	if r.IsExist(key) {
		err := r.Delete(key)
		if err != nil {
			return err
		}
	}
	return r.rdsClient.Set(key, val, duration).Err()
}

// IsExist false->non exist
func (r *RedisV3) IsExist(key string) bool {
	return r.rdsClient.Exists(key).Val()
}

// Delete
func (r *RedisV3) Delete(key string) error {
	return r.rdsClient.Del(key).Err()
}
