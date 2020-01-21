package wechat

import (
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/miniprogram"
	miniConfig "github.com/silenceper/wechat/miniprogram/config"
	"github.com/silenceper/wechat/officialaccount"
	offConfig "github.com/silenceper/wechat/officialaccount/config"
	"github.com/silenceper/wechat/openplatform"
	opConfig "github.com/silenceper/wechat/openplatform/config"
	"github.com/silenceper/wechat/pay"
	payConfig "github.com/silenceper/wechat/pay/config"
)

// Wechat struct
type Wechat struct {
	cache cache.Cache
}

// Config for user
type Config struct {
	PayMchID     string //支付 - 商户 ID
	PayNotifyURL string //支付 - 接受微信支付结果通知的接口地址
	PayKey       string //支付 - 商户后台设置的支付 key
}

// NewWechat init
func NewWechat() *Wechat {
	return &Wechat{}
}

//SetCache 设置cache
func (wc *Wechat) SetCache(cahce cache.Cache) {
	wc.cache = cahce
}

//GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	if cfg.Cache == nil {
		cfg.Cache = wc.cache
	}
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}

// GetOpenPlatform 获取微信开放平台的实例
func (wc *Wechat) GetOpenPlatform(cfg *opConfig.Config) *openplatform.OpenPlatform {
	return openplatform.NewOpenPlatform(cfg)
}
