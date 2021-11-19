package openplatform

import (
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/code"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/privacyconfig"
	"net/http"

	"github.com/silenceper/wechat/v2/officialaccount/server"
	"github.com/silenceper/wechat/v2/openplatform/account"
	"github.com/silenceper/wechat/v2/openplatform/config"
	"github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram"
	"github.com/silenceper/wechat/v2/openplatform/officialaccount"
)

//OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

//NewOpenPlatform new openplatform
func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

//GetServer get server
func (openPlatform *OpenPlatform) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	off := officialaccount.NewOfficialAccount(openPlatform.Context, "")
	return off.GetServer(req, writer)
}

//GetOfficialAccount 公众号代处理
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}

//GetMiniProgram 小程序代理
func (openPlatform *OpenPlatform) GetMiniProgram(appID string) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(openPlatform.Context, appID)
}

//GetCode 小程序代码
func (openPlatform *OpenPlatform) GetCode(appID string) *code.Code {
	return code.NewCode(openPlatform.Context, appID)
}

//GetPrivacyConfig 小程序用户隐私保护指引
func (openPlatform *OpenPlatform) GetPrivacyConfig(appID string) *privacyconfig.PrivacyConfig {
	return privacyconfig.NewPrivacyConfig(openPlatform.Context, appID)
}

//GetAccountManager 账号管理
//TODO
func (openPlatform *OpenPlatform) GetAccountManager() *account.Account {
	return account.NewAccount(openPlatform.Context)
}
