package cache

import "time"

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeput time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}
