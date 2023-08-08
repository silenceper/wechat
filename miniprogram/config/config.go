// Package config 小程序config配置
package config

import (
	"github.com/silenceper/wechat/v2/cache"
)

// Config .config for 小程序
type Config struct {
	AppID     string `json:"app_id"`     // appid
	AppSecret string `json:"app_secret"` // appSecret
	Cache     cache.Cache

	//使用云托管开放接口方式访问接口，此时API调用不附着AccessToken
	//参考文档：https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/guide/weixin/open.html
	NoAccessToken bool `json:"cloud_run_access"`
	//是否使用云托管开放接口方式访问接口，此时API调用可以开启使用HTTP API
	//参考文档：https://developers.weixin.qq.com/miniprogram/dev/wxcloudrun/src/guide/weixin/open.html
	UsingHTTP bool `json:"using_http"`
}
