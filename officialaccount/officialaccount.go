package officialaccount

import (
	"net/http"

	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/officialaccount/device"
	"github.com/silenceper/wechat/v2/officialaccount/js"
	"github.com/silenceper/wechat/v2/officialaccount/material"
	"github.com/silenceper/wechat/v2/officialaccount/menu"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/oauth"
	"github.com/silenceper/wechat/v2/officialaccount/server"
	"github.com/silenceper/wechat/v2/officialaccount/user"
)

//OfficialAccount 微信公众号相关API
type OfficialAccount struct {
	ctx *context.Context
}

//NewOfficialAccount 实例化公众号API
func NewOfficialAccount(cfg *config.Config) *OfficialAccount {
	defaultAK := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyOfficialAccountPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAK,
	}
	return &OfficialAccount{ctx: ctx}
}

//SetAccessTokenHandle 自定义access_token获取方式
func (officialAccount *OfficialAccount) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	officialAccount.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (officialAccount *OfficialAccount) GetContext() *context.Context {
	return officialAccount.ctx
}

// GetBasic qr/url 相关配置
func (officialAccount *OfficialAccount) GetBasic() *basic.Basic {
	return basic.NewBasic(officialAccount.ctx)
}

// GetMenu 菜单管理接口
func (officialAccount *OfficialAccount) GetMenu() *menu.Menu {
	return menu.NewMenu(officialAccount.ctx)
}

// GetServer 消息管理
func (officialAccount *OfficialAccount) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	srv := server.NewServer(officialAccount.ctx)
	srv.Request = req
	srv.Writer = writer
	return srv
}

//GetAccessToken 获取access_token
func (officialAccount *OfficialAccount) GetAccessToken() (string, error) {
	return officialAccount.ctx.GetAccessToken()
}

// GetOauth oauth2网页授权
func (officialAccount *OfficialAccount) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(officialAccount.ctx)
}

// GetMaterial 素材管理
func (officialAccount *OfficialAccount) GetMaterial() *material.Material {
	return material.NewMaterial(officialAccount.ctx)
}

// GetJs js-sdk配置
func (officialAccount *OfficialAccount) GetJs() *js.Js {
	return js.NewJs(officialAccount.ctx)
}

// GetUser 用户管理接口
func (officialAccount *OfficialAccount) GetUser() *user.User {
	return user.NewUser(officialAccount.ctx)
}

// GetTemplate 模板消息接口
func (officialAccount *OfficialAccount) GetTemplate() *message.Template {
	return message.NewTemplate(officialAccount.ctx)
}

// GetDevice 获取智能设备的实例
func (officialAccount *OfficialAccount) GetDevice() *device.Device {
	return device.NewDevice(officialAccount.ctx)
}
