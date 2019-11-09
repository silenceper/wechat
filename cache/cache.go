package cache

import "time"

const (
	// ComponentVerifyTicket 票据
	ComponentVerifyTicket = "component_verify_ticket_%s"
	// ComponentAccessToken 开放平台apitoken
	ComponentAccessToken = "component_access_token_%s"
)

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}
