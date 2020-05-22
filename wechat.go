package wechat

import (
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	payConfig "github.com/silenceper/wechat/v2/pay/config"

	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/pay"
)

// Wechat struct
type Wechat struct {
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

//GetOfficialAccount 获取微信公众号实例
func (wc *Wechat) GetOfficialAccount(cfg *offConfig.Config) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(cfg)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram(cfg *miniConfig.Config) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(cfg)
}

// GetPay 获取微信支付的实例
func (wc *Wechat) GetPay(cfg *payConfig.Config) *pay.Pay {
	return pay.NewPay(cfg)
}
