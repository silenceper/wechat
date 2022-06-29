package core

import (
	"github.com/imroc/req/v3"
	"github.com/silenceper/wechat/v2/cache"
)

// Config .config for 微信公众号
type Config struct {
	AppID     string `json:"app_id"`     // appid
	AppSecret string `json:"app_secret"` // appsecret
	ProxyUrl  string `json:"proxy_url"`  // 代理url
	Cache     cache.Cache
}

func (c *Config) Req() *req.Request {
	client := req.C()
	if c.ProxyUrl != "" {
		client.SetProxyURL(c.ProxyUrl)
	}
	return client.R()
}
