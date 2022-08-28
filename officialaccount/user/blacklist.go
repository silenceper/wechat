// Package user blacklist 公众号用户黑名单管理
// 参考文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
package user

import (
	"errors"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 获取公众号的黑名单列表
	getblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"
	// 拉黑用户
	batchblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"
	// 取消拉黑用户
	batchunblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s"
)

// GetBlackList 获取公众号的黑名单列表
// 该接口每次调用最多可拉取 1000 个OpenID，当列表数较多时，可以通过多次拉取的方式来满足需求。
// 参数 beginOpenid：当 begin_openid 为空时，默认从开头拉取。
func (user *User) GetBlackList(beginOpenid ...string) (userlist *OpenidList, err error) {
	//* 获取 AccessToken
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	//* 调用接口
	var (
		resp    []byte
		request map[string]string
	)
	url := fmt.Sprintf(getblacklistURL, accessToken)
	if len(beginOpenid) == 1 {
		// 传入 begin_openid
		request["begin_openid"] = beginOpenid[0]
		resp, err = util.PostJSON(url, &request)
	} else {
		// 无requestBody
		resp, err = util.PostJSON(url, nil)
	}
	if err != nil {
		return nil, err
	}

	userlist = &OpenidList{}
	err = util.DecodeWithError(resp, userlist, "ListUserOpenIDs")
	if err != nil {
		return nil, err
	}

	return
}

// BatchBlackList 拉黑用户
// 参数 openidList：需要拉入黑名单的用户的openid，每次拉黑最多允许20个
func (user *User) BatchBlackList(openidList ...string) (err error) {
	//* 检查参数
	if len(openidList) == 0 || len(openidList) > 20 {
		return errors.New("参数 openidList 错误：每次拉黑用户数量为1-20个。")
	}

	//* 获取 AccessToken
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	//* 调用接口
	url := fmt.Sprintf(batchblacklistURL, accessToken)
	request := map[string][]string{"openid_list": openidList}
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}

	return util.DecodeWithCommonError(resp, "BatchBlackList")
}

// BatchunBlackList 取消拉黑用户
// 参数 openidList：需要取消拉入黑名单的用户的openid，每次拉黑最多允许20个
func (user *User) BatchunBlackList(openidList ...string) (err error) {
	//* 检查参数
	if len(openidList) == 0 || len(openidList) > 20 {
		return errors.New("参数 openidList 错误：每次取消拉黑用户数量为1-20个。")
	}

	//* 获取 AccessToken
	var accessToken string
	accessToken, err = user.GetAccessToken()
	if err != nil {
		return
	}

	//* 调用接口
	url := fmt.Sprintf(batchunblacklistURL, accessToken)
	request := map[string][]string{"openid_list": openidList}
	resp, err := util.PostJSON(url, &request)
	if err != nil {
		return
	}

	return util.DecodeWithCommonError(resp, "BatchunBlackList")
}
