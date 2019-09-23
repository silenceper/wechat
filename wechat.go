package wechat

import (
	"net/http"
	"sync"
	"github.com/machao520/wechat/wechat"
	"github.com/machao520/wechat/cache"
	"github.com/machao520/wechat/context"
	"github.com/machao520/wechat/js"
	"github.com/machao520/wechat/material"
	"github.com/machao520/wechat/menu"
	"github.com/machao520/wechat/miniprogram"
	"github.com/machao520/wechat/oauth"
	"github.com/machao520/wechat/pay"
	"github.com/machao520/wechat/qr"
	"github.com/machao520/wechat/server"
	"github.com/machao520/wechat/template"
	"github.com/machao520/wechat/user"
)
//
//// Wechat struct
type Wechat struct {
	*wechat.Wechat
}

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string //支付 - 商户 ID
	PayNotifyURL   string //支付 - 接受微信支付结果通知的接口地址
	PayKey         string //支付 - 商户后台设置的支付 key
	Cache          cache.Cache
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	wc := new(wechat.Wechat)
	copyConfigToContext(cfg, wc)
	return &Wechat{wc}
}

func copyConfigToContext(cfg *Config, wc *wechat.Wechat) {
	wc.AppID = cfg.AppID
	wc.AppSecret = cfg.AppSecret
	wc.Token = cfg.Token
	wc.EncodingAESKey = cfg.EncodingAESKey
	wc.PayMchID = cfg.PayMchID
	wc.PayKey = cfg.PayKey
	wc.PayNotifyURL = cfg.PayNotifyURL
	wc.Cache = cfg.Cache
	wc.SetAccessTokenLock(new(sync.RWMutex))
	wc.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	context := new(context.Context)
	context.Request = req
	context.Writer = writer
	return server.NewServer(context, wc.Wechat)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Wechat)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Wechat)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Wechat)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Wechat)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Wechat)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Wechat)
}

// GetPay 返回支付消息的实例
func (wc *Wechat) GetPay() *pay.Pay {
	return pay.NewPay(wc.Wechat)
}

// GetQR 返回二维码的实例
func (wc *Wechat) GetQR() *qr.QR {
	return qr.NewQR(wc.Wechat)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram() *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(wc.Wechat)
}
