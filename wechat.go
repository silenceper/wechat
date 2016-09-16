package wechat

import (
	"net/http"
	"sync"

	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/js"
	"github.com/silenceper/wechat/material"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/oauth"
	"github.com/silenceper/wechat/server"
)

//Wechat struct
type Wechat struct {
	Context *context.Context
}

//Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Cache          cache.Cache
}

//NewWechat init
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
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsApiTicketLock(new(sync.RWMutex))
}

//GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

//GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

//GetOauth oauth2网页授权
func (wc *Wechat) GetOauth(req *http.Request, writer http.ResponseWriter) *oauth.Oauth {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return oauth.NewOauth(wc.Context)
}

//GetJs js-sdk配置
func (wc *Wechat) GetJs(req *http.Request, writer http.ResponseWriter) *js.Js {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return js.NewJs(wc.Context)
}

//GetMenu 菜单管理接口
func (wc *Wechat) GetMenu(req *http.Request, writer http.ResponseWriter) *menu.Menu {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return menu.NewMenu(wc.Context)
}
