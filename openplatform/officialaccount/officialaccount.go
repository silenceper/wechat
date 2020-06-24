package officialaccount

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	opContext "github.com/silenceper/wechat/v2/openplatform/context"
)

//OfficialAccount 代公众号实现业务
type OfficialAccount struct {
	//授权的公众号的appID
	appID string
	*officialaccount.OfficialAccount
}

//NewOfficialAccount 实例化
//appID :为授权方公众号 APPID，非开放平台第三方平台 APPID
func NewOfficialAccount(opCtx *opContext.Context, appID string) *OfficialAccount {
	officialAccount := officialaccount.NewOfficialAccount(&offConfig.Config{
		AppID:          opCtx.AppID,
		EncodingAESKey: opCtx.EncodingAESKey,
		Token:          opCtx.Token,
		Cache:          opCtx.Cache,
	})
	//设置获取access_token的函数
	officialAccount.SetAccessTokenHandle(NewDefaultAuthrAccessToken(opCtx, appID))
	return &OfficialAccount{appID: appID, OfficialAccount: officialAccount}
}

//DefaultAuthrAccessToken 默认获取授权ak的方法
type DefaultAuthrAccessToken struct {
	opCtx *opContext.Context
	appID string
}

//NewDefaultAuthrAccessToken New
func NewDefaultAuthrAccessToken(opCtx *opContext.Context, appID string) credential.AccessTokenHandle {
	return &DefaultAuthrAccessToken{
		opCtx: opCtx,
		appID: appID,
	}
}

//GetAccessToken 获取ak
func (ak *DefaultAuthrAccessToken) GetAccessToken() (string, error) {
	return ak.opCtx.GetAuthrAccessToken(ak.appID)
}
