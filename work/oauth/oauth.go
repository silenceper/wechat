package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

// Oauth auth
type Oauth struct {
	*context.Context
}

var (
	// oauthTargetURL 企业微信内跳转地址
	oauthTargetURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect"
	// oauthUserInfoURL 获取用户信息地址
	oauthUserInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
	// oauthQrContentTargetURL 构造独立窗口登录二维码
	oauthQrContentTargetURL = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%s&redirect_uri=%s&state=%s"
)

// NewOauth new init oauth
func NewOauth(ctx *context.Context) *Oauth {
	return &Oauth{
		ctx,
	}
}

// GetTargetURL 获取授权地址
func (ctr *Oauth) GetTargetURL(callbackURL string) string {
	// url encode
	urlStr := url.QueryEscape(callbackURL)
	return fmt.Sprintf(
		oauthTargetURL,
		ctr.CorpID,
		urlStr,
	)
}

// GetQrContentTargetURL 构造独立窗口登录二维码
func (ctr *Oauth) GetQrContentTargetURL(callbackURL string) string {
	// url encode
	urlStr := url.QueryEscape(callbackURL)
	return fmt.Sprintf(
		oauthQrContentTargetURL,
		ctr.CorpID,
		ctr.AgentID,
		urlStr,
		util.RandomStr(16),
	)
}

// ResUserInfo 返回得用户信息
type ResUserInfo struct {
	util.CommonError
	// 当用户为企业成员时返回
	UserID   string `json:"UserId"`
	DeviceID string `json:"DeviceId"`
	// 非企业成员授权时返回
	OpenID string `json:"OpenId"`
}

// UserFromCode 根据code获取用户信息
func (ctr *Oauth) UserFromCode(code string) (result ResUserInfo, err error) {
	var accessToken string
	accessToken, err = ctr.GetAccessToken()
	if err != nil {
		return
	}
	var response []byte
	response, err = util.HTTPGet(
		fmt.Sprintf(oauthUserInfoURL, accessToken, code),
	)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
