package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/swxctx/wechat/util"
)

/*
	1. 通过授权url到指定地址授权后获取到code
*/

// GetRedirectURL 获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state string) string {
	return fmt.Sprintf(redirectOauthURL, oauth.AppID, url.QueryEscape(redirectURI), scope, state)
}

// Redirect 跳转到网页授权
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state string) {
	location := oauth.GetRedirectURL(redirectURI, scope, state)
	// 重定向
	http.Redirect(writer, req, location, 302)
}

/*
	2. 区分于接口请求的全局accessToken
*/

// ResAccessToken 获取用户授权access_token的返回结果
type ResAccessToken struct {
	util.CommonError

	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// GetUserAccessToken 通过网页授权的code 换取access_token(区别于context中的access_token)
func (oauth *Oauth) GetUserAccessToken(code string) (result ResAccessToken, err error) {
	var (
		response []byte
	)
	urlStr := fmt.Sprintf(accessTokenURL, oauth.AppID, oauth.AppSecret, code)
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//RefreshAccessToken 刷新access_token
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error) {
	var (
		response []byte
	)

	urlStr := fmt.Sprintf(refreshAccessTokenURL, oauth.AppID, refreshToken)
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//CheckAccessToken 检验access_token是否有效
func (oauth *Oauth) CheckAccessToken(accessToken, openID string) (b bool, err error) {
	var (
		response []byte
	)

	urlStr := fmt.Sprintf(checkAccessTokenURL, accessToken, openID)
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		b = false
		return
	}
	b = true
	return
}

/*
	3. 通过openId及accessToken获取用户基本信息
*/

//UserInfo 用户授权获取到用户信息
type UserInfo struct {
	util.CommonError

	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgURL string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

//GetUserInfo 如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信息
func (oauth *Oauth) GetUserInfo(accessToken, openID string) (result UserInfo, err error) {
	var (
		response []byte
	)

	urlStr := fmt.Sprintf(userInfoURL, accessToken, openID)
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
