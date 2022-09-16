package auth

import (
	context2 "context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram"
	miniprogramAuth "github.com/silenceper/wechat/v2/miniprogram/auth"
	"github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
)

// Auth 登录/用户信息
type Auth struct {
	*context.Context
	appID string
}

// NewAuth new auth (授权方appID)
func NewAuth(ctx *context.Context, appID string) *Auth {
	return &Auth{ctx, appID}
}

// Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string) (result miniprogramAuth.ResCode2Session, err error) {
	return auth.Code2SessionContext(context2.Background(), jsCode)
}

// Code2SessionContext 登录凭证校验。
func (auth *Auth) Code2SessionContext(ctx context2.Context, jsCode string) (result miniprogramAuth.ResCode2Session, err error) {
	var response []byte
	var componentAccessToken string
	componentAccessToken, err = auth.GetComponentAccessToken()
	if err != nil {
		return
	}
	if response, err = util.HTTPGetContext(ctx, fmt.Sprintf(code2SessionURL, auth.appID, jsCode, auth.Context.AppID, componentAccessToken)); err != nil {
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

// CheckEncryptedData .检查加密信息是否由微信生成（当前只支持手机号加密数据），只能检测最近3天生成的加密数据
func (auth *Auth) CheckEncryptedData(encryptedMsgHash string) (result miniprogramAuth.RspCheckEncryptedData, err error) {
	var miniProgram = miniprogram.MiniProgram{}
	miniProgram.SetAccessTokenHandle(auth.AccessTokenHandle)
	return miniProgram.GetAuth().CheckEncryptedDataContext(context2.Background(), encryptedMsgHash)
}
