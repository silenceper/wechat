package miniprogram

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	openContext "github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/auth"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/basic"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/component"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/encryptor"
)

//MiniProgram 代小程序实现业务
type MiniProgram struct {
	AppID       string
	openContext *openContext.Context
	*miniprogram.MiniProgram
}

//NewMiniProgram 实例化
func NewMiniProgram(opCtx *openContext.Context, appID string) *MiniProgram {
	mini := miniprogram.NewMiniProgram(&miniConfig.Config{
		AppID:          opCtx.AppID,
		EncodingAESKey: opCtx.EncodingAESKey,
		Token:          opCtx.Token,
		Cache:          opCtx.Cache,
	})
	//设置获取access_token的函数
	mini.SetAccessTokenHandle(NewDefaultAuthrAccessToken(opCtx, appID))
	return &MiniProgram{
		openContext: opCtx,
		AppID:       appID,
		MiniProgram: mini,
	}
}

// PlatformAuth
func (miniProgram *MiniProgram) PlatformAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.GetContext())
}

// PlatformEncryptor
func (miniProgram *MiniProgram) PlatformEncryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(miniProgram.GetContext())
}

//GetComponent get component
//快速注册小程序相关
func (miniProgram *MiniProgram) GetComponent() *component.Component {
	return component.NewComponent(miniProgram.openContext)
}

//GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetBasic() *basic.Basic {
	return basic.NewBasic(miniProgram.openContext, miniProgram.AppID)
}

//DefaultAuthrAccessToken 默认获取授权ak的方法
type DefaultAuthrAccessToken struct {
	opCtx *openContext.Context
	appID string
}

//NewDefaultAuthrAccessToken New
func NewDefaultAuthrAccessToken(opCtx *openContext.Context, appID string) credential.AccessTokenHandle {
	return &DefaultAuthrAccessToken{
		opCtx: opCtx,
		appID: appID,
	}
}

//GetAccessToken 获取ak
func (ak *DefaultAuthrAccessToken) GetAccessToken() (string, error) {
	return ak.opCtx.GetAuthrAccessToken(ak.appID)
}
