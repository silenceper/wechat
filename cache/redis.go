package cache

import (
	"time"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn *redis.Pool
}

func NewRedis(rediscon, redispass string) *Redis {
	conn := &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rediscon)
			if err != nil {
				fmt.Println(err)
			}
			if redispass != "" {
				if _, err := c.Do("AUTH", redispass); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, nil
		},
	}
	return &Redis{conn}
}

func (c *Redis) Get(key string) interface{} {
	C := c.conn.Get()
	defer C.Close()
	if value, err := C.Do("GET", key); err == nil {
		a, _ := redis.String(value, err)
		return a
	}
	return nil
}

func (c *Redis) IsExist(key string) bool {
	C := c.conn.Get()
	defer C.Close()
	a, _ := C.Do("EXISTS", key)
	i:=a.(int64)
	if i >0 {
		return true
	}
	return false
}

func (c *Redis) Set(key string, val interface{}, timeout time.Duration) error {
	C := c.conn.Get()
	defer C.Close()
	_, err := C.Do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

func (c *Redis) Delete(key string) error {
	C := c.conn.Get()
	defer C.Close()
	_, err := C.Do("DEL", key)
	return err
}
