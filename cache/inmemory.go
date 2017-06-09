package cache

import (
	"errors"
	"time"

	"github.com/robfig/go-cache"
)

type InMemoryCache struct {
	cache.Cache
}

func NewInMemoryCache(defaultExpiration time.Duration) InMemoryCache {
	return InMemoryCache{*cache.New(defaultExpiration, time.Minute)}
}

func (c InMemoryCache) Get(key string) interface{} {
	value, found := c.Cache.Get(key)
	if !found {
		return nil
	}

	v, ok := value.(string)
	if !ok {
		return nil
	}

	return string(v)
}

func (c InMemoryCache) Set(key string, val interface{}, timeout time.Duration) error {
	v, ok := val.(string)
	if !ok {
		return errors.New("val must string")
	}

	c.Cache.Set(key, v, timeout)
	return nil
}

func (c InMemoryCache) IsExist(key string) bool {
	_, found := c.Cache.Get(key)
	if !found {
		return false
	}

	return true
}

func (c InMemoryCache) Delete(key string) error {
	if found := c.Cache.Delete(key); !found {
		return errors.New("key not found")
	}

	return nil
}
