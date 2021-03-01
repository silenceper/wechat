package auth

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	code2SessionURL = "https://api.weixin.qq.com/sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s"
)

//Auth 登录/用户信息
type Auth struct {
	*context.Context
}

//NewAuth new auth
func NewAuth(ctx *context.Context) *Auth {
	return &Auth{ctx}
}

// ResCode2Session 登录凭证校验的返回结果
type ResCode2Session struct {
	util.CommonError

	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

//Code2Session 登录凭证校验。
func (auth *Auth) Code2Session(jsCode string, appID string, componentAccessToken string) (result ResCode2Session, err error) {
	urlStr := fmt.Sprintf(code2SessionURL, appID, jsCode, auth.AppID, componentAccessToken)
	fmt.Println(urlStr)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	fmt.Println(string(response), err)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("Code2Session error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//GetPaidUnionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func (auth *Auth) GetPaidUnionID() {
	//TODO
}
