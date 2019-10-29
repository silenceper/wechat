package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"gitee.com/zhimiao/wechat-sdk/util"
)

var (
	qyRedirectOauthURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&agentid=%s&state=%s#wechat_redirect"
	qyUserInfoURL      = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
	qyUserDetailURL    = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserdetail"
)

//GetQyRedirectURL 获取企业微信跳转的url地址
func (oauth *Oauth) GetQyRedirectURL(redirectURI, agentid, scope, state string) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(qyRedirectOauthURL, oauth.AppID, urlStr, scope, agentid, state), nil
}

//QyUserInfo 用户授权获取到用户信息
type QyUserInfo struct {
	util.CommonError

	UserID     string `json:"UserId"`
	DeviceID   string `json:"DeviceId"`
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int64  `json:"expires_in"`
}

//GetQyUserInfoByCode 根据code获取企业user_info
func (oauth *Oauth) GetQyUserInfoByCode(code string) (result QyUserInfo, err error) {
	qyAccessToken, e := oauth.GetQyAccessToken()
	if e != nil {
		err = e
		return
	}
	urlStr := fmt.Sprintf(qyUserInfoURL, qyAccessToken, code)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyUserInfoByCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//QyUserDetail 到用户详情
type QyUserDetail struct {
	util.CommonError

	UserID string `json:"UserId"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Gender string `json:"gender"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
	QrCode string `json:"qr_code"`
}

//GetQyUserDetailUserTicket 根据user_ticket获取到用户详情
func (oauth *Oauth) GetQyUserDetailUserTicket(userTicket string) (result QyUserDetail, err error) {
	var qyAccessToken string
	qyAccessToken, err = oauth.GetQyAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", qyUserDetailURL, qyAccessToken)
	var response []byte
	response, err = util.PostJSON(uri, map[string]string{
		"user_ticket": userTicket,
	})
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyUserDetailUserTicket Error , errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
