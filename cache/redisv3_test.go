package cache

import (
	"testing"
	"time"

	redis "gopkg.in/redis.v3"
)

func TestRdsV3(t *testing.T) {
	rdsv3 := RdsV3ClientInit(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}))
	rdsv3.Set("test", "redisv3", time.Duration(time.Minute*10))
	val := rdsv3.Get("test")
	t.Logf("val:%s", val.(string))
}
