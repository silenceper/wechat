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

//CreateOpenRes 新增开放平台返回
type CreateOpenRes struct {
	OpenAppid string `json:"open_appid"`
	util.CommonError
}

//GetOpenRes 获取公众号/小程序的开放平台返回
type GetOpenRes struct {
	OpenAppid string `json:"open_appid"`
	util.CommonError
}

//BindRes 获取公众号/小程序的绑定开放平台返回
type BindRes struct {
	util.CommonError
}

//UnbindRes 获取公众号/小程序的解绑开放平台返回
type UnbindRes struct {
	util.CommonError
}

type Account struct {
	*context.Context
}

//NewAccount new
func NewAccount(ctx *context.Context) *Account {
	return &Account{ctx}
}

//Create 创建开放平台帐号并绑定公众号/小程序
func (account *Account) Create(appID string) (string, error) {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	if err != nil {
		return "", err
	}
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
	if ret.ErrCode != 0 {
		err = util.DecodeWithError(body, ret, "Create")
		return "", err
	}
	return ret.OpenAppid, nil
}

//Bind 将公众号/小程序绑定到开放平台帐号下
func (account *Account) Bind(appID string, openAppID string) error {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	if err != nil {
		return err
	}
	req := map[string]string{
		"appid":      appID,
		"open_appid": openAppID,
	}
	uri := fmt.Sprintf(bindOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	ret := &BindRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return err
	}
	if ret.ErrCode != 0 {
		err = util.DecodeWithError(body, ret, "Bind")
		return err
	}
	return nil
}

//Unbind 将公众号/小程序从开放平台帐号下解绑
func (account *Account) Unbind(appID string, openAppID string) error {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	if err != nil {
		return err
	}
	req := map[string]string{
		"appid":      appID,
		"open_appid": openAppID,
	}
	uri := fmt.Sprintf(unbindOpenURL, accessToken)
	body, err := util.PostJSON(uri, req)
	if err != nil {
		return err
	}
	ret := &UnbindRes{}
	if err := json.Unmarshal(body, ret); err != nil {
		return err
	}
	if ret.ErrCode != 0 {
		err = util.DecodeWithError(body, ret, "Unbind")
		return err
	}
	return nil
}

//Get 获取公众号/小程序所绑定的开放平台帐号
func (account *Account) Get(appID string) (string, error) {
	accessToken, err := account.Context.GetAuthrAccessToken(appID)
	if err != nil {
		return "", err
	}
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
	if ret.ErrCode != 0 {
		err = util.DecodeWithError(body, ret, "Get")
		return "", err
	}
	return ret.OpenAppid, nil
}
