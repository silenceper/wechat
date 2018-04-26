package cache

import (
	"time"

	"github.com/kataras/golog"
	redis "gopkg.in/redis.v3"
)

var rdsClient *redis.Client

// Init 初始化 redis client
func CacheInit(client *redis.Client) {
	if rdsClient == nil {
		rdsClient = client
	}
}

// checkRedis check
func checkRedis() {
	if rdsClient == nil {
		golog.Error("wx server cache: redis client is nil")
	}
}

// Get
func Get(key string) string {
	checkRedis()
	return rdsClient.Get(key).String()
}

// Set
func Set(key string, val interface{}, duration time.Duration) (err error) {
	checkRedis()
	if IsExist(key) {
		err := Delete(key)
		if err != nil {
			return err
		}
	}
	return rdsClient.Set(key, val, duration).Err()
}

// IsExist false->non exist
func IsExist(key string) bool {
	checkRedis()
	return rdsClient.Exists(key).Val()
}

// Delete
func Delete(key string) error {
	checkRedis()
	return rdsClient.Del(key).Err()
}
