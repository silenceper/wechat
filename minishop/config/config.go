//Package config 小商店config配置
package config

import (
	"github.com/silenceper/wechat/v2/cache"
)

//Config config for 小商店
type Config struct {
	AppID           string `json:"app_id"`     //appid(小程序)
	AppSecret       string `json:"app_secret"` //appsecret(小程序)
	ServiceID       string
	SpecificationID string
	Cache           cache.Cache
}
