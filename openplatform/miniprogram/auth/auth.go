package auth

import (
	stdContext "context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
)

// Auth 登录/用户信息
type Auth struct {
	*context.Context
	authorizerAppID string
}

// NewAuth new auth (授权方appID)
func NewAuth(ctx *context.Context, appID string) *Auth {
	return &Auth{ctx, appID}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

// Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string) (result ResCode2Session, err error) {
	return auth.Code2SessionContext(stdContext.Background(), jsCode)
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx stdContext.Context, jsCode string) (result ResCode2Session, err error) {
	var response []byte
	var componentAccessToken string
	componentAccessToken, err = auth.GetComponentAccessToken()
	if err != nil {
		return
	}
	parse := fmt.Sprintf(code2SessionURL, auth.authorizerAppID, jsCode, auth.Context.AppID, componentAccessToken)
	if response, err = util.HTTPGetContext(ctx, parse); err != nil {
		return
	}
	if err = json.Unmarshal(response, &result); err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
