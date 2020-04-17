package wechat

import (
	"net/http"
	"sync"

	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/delivery"
	"github.com/silenceper/wechat/device"
	"github.com/silenceper/wechat/js"
	"github.com/silenceper/wechat/material"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/miniprogram"
	"github.com/silenceper/wechat/oauth"
	"github.com/silenceper/wechat/pay"
	"github.com/silenceper/wechat/qr"
	"github.com/silenceper/wechat/server"
	"github.com/silenceper/wechat/tcb"
	"github.com/silenceper/wechat/user"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
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
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.PayMchID = cfg.PayMchID
	context.PayKey = cfg.PayKey
	context.PayNotifyURL = cfg.PayNotifyURL
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *message.Template {
	return message.NewTemplate(wc.Context)
}

// GetPay 返回支付消息的实例
func (wc *Wechat) GetPay() *pay.Pay {
	return pay.NewPay(wc.Context)
}

// GetQR 返回二维码的实例
func (wc *Wechat) GetQR() *qr.QR {
	return qr.NewQR(wc.Context)
}

// GetMiniProgram 获取小程序的实例
func (wc *Wechat) GetMiniProgram() *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(wc.Context)
}

// GetDevice 获取智能设备的实例
func (wc *Wechat) GetDevice() *device.Device {
	return device.NewDevice(wc.Context)
}

// GetTcb 获取小程序-云开发的实例
func (wc *Wechat) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(wc.Context)
}

// GetDelivery 物流管理
func (wc *Wechat) GetDelivery() *delivery.Delivery {
	return delivery.NewDelivery(wc.Context)
}
