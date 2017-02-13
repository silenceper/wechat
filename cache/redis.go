package cache

import (
	"errors"
	"time"

	"menteslibres.net/gosexy/redis"
)

//Redis struct contains *Rediscache.Client
type Rediscache struct {
	client *redis.Client
}

//NewRediscache create new Redis
func NewRediscache(host string, port uint) (*Rediscache, error) {
	client := redis.New()
	err := client.Connect(host, port)
	return &Rediscache{client}, err
}

//Get return cached value
func (res *Rediscache) Get(key string) interface{} {
	if item, err := res.client.Get(key); err == nil {
		return item
	}
	return nil
}

// IsExist check value exists in Redis.
func (res *Rediscache) IsExist(key string) bool {
	_, err := res.client.Exists(key)
	if err != nil {
		return false
	}
	return true
}

//Set cached value with key and expire time.
func (res *Rediscache) Set(key string, val interface{}, timeout time.Duration) error {
	_, ok := val.(string)
	if !ok {
		return errors.New("val must string")
	}
	_, err := res.client.SetEx(key, int64(timeout/time.Second), val)
	return err
}

//Delete delete value in Redis.
func (res *Rediscache) Delete(key string) error {
	_, err := res.client.Del(key)
	return err
}
