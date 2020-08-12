package account

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/util"
)

//Account 开放平台张哈管理
const (
	createOpenURL = "https://api.weixin.qq.com/cgi-bin/open/create?access_token=%s"
	bindOpenURL   = "https://api.weixin.qq.com/cgi-bin/open/bind?access_token=%s"
	unbindOpenURL = "https://api.weixin.qq.com/cgi-bin/open/unbind?access_token=%s"
	getOpenURL    = "https://api.weixin.qq.com/cgi-bin/open/get?access_token=%s"
)

type CreateOpenRes struct {
	OpenAppid string `json:"open_appid"`
	BaseRes
}
type GetOpenRes struct {
	OpenAppid string `json:"open_appid"`
	BaseRes
}
type BaseRes struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type Account struct {
	*context.Context
}

//NewAccount new
func NewAccount(ctx *context.Context) *Account {
	return &Account{ctx}
}

//Create 创建开放平台帐号并绑定公众号/小程序
func (account *Account) Create(appID string) (openAppId string, err error) {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	req := map[string]string{
		"appid": appID,
	}
	uri := fmt.Sprintf(createOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return "", err
	}
	ret := &CreateOpenRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return "", err
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf("Create error : errcode=%v , errmsg=%v", ret.Errcode, ret.Errmsg)
		return "", err
	}
	return ret.OpenAppid, nil
}

//Bind 将公众号/小程序绑定到开放平台帐号下
func (account *Account) Bind(appID string, openAppId string) error {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	req := map[string]string{
		"appid":      appID,
		"open_appid": openAppId,
	}
	uri := fmt.Sprintf(bindOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	ret := &BaseRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return err
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf("Bind error : errcode=%v , errmsg=%v", ret.Errcode, ret.Errmsg)
		return err
	}
	return nil
}

//Unbind 将公众号/小程序从开放平台帐号下解绑
func (account *Account) Unbind(appID string, openAppID string) error {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	req := map[string]string{
		"appid":      appID,
		"open_appid": openAppID,
	}
	uri := fmt.Sprintf(unbindOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	ret := &BaseRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return err
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf("Unbind error : errcode=%v , errmsg=%v", ret.Errcode, ret.Errmsg)
		return err
	}
	return nil
}

//Get 获取公众号/小程序所绑定的开放平台帐号
func (account *Account) Get(appID string) (string, error) {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	req := map[string]string{
		"appid": appID,
	}
	uri := fmt.Sprintf(getOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return "", err
	}
	ret := &GetOpenRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return "", err
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf("Get error : errcode=%v , errmsg=%v", ret.Errcode, ret.Errmsg)
		return "", err
	}
	return ret.OpenAppid, nil
}
