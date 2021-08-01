// Package config 小程序config配置
package config

import (
	"github.com/silenceper/wechat/v2/cache"
)

// Config config for 小程序
type Config struct {
	AppID          string `json:"app_id"`           // app_id
	AppSecret      string `json:"app_secret"`       // app_secret
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // encoding_aes_key
	Cache          cache.Cache
}
