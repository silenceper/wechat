package cache

import "time"

const (
	COMPONENT_VERIFY_TICKET = "component_verify_ticket_%s"
	COMPONENT_ACCESS_TOKEN = "component_access_token_%s"
)

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}
