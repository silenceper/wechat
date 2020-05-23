package config

import (
	"github.com/silenceper/wechat/v2/cache"
)

//Config config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     //appid
	AppSecret string `json:"app_secret"` //appsecret
	Cache     cache.Cache
}
