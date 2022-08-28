// Package user blacklist 公众号用户黑名单管理
// 参考文档：https://developers.weixin.qq.com/doc/offiaccount/User_Management/Manage_blacklist.html
package user

import (
	"log"
)

const (
	// 获取公众号的黑名单列表
	getblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"
	// 拉黑用户
	batchblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"
	// 取消拉黑用户
	batchunblacklistURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN"
)

// TODO: // GetBlackList 获取公众号的黑名单列表
// func (user *User) GetBlackList() (err error) {
// 	var accessToken string
// 	accessToken, err = user.GetAccessToken()
// 	if err != nil {
// 		return
// 	}
// 	uri := fmt.Sprintf(getblacklistURL, accessToken)
// }

// BatchBlackList 拉黑用户
// 参数 openidList：需要拉入黑名单的用户的openid，一次拉黑最多允许20个
func (user *User) BatchBlackList(openidList ...string) (err error) {
	// var accessToken string
	// accessToken, err = user.GetAccessToken()
	// if err != nil {
	// 	return
	// }

	request := map[string][]string{
		"openid_list": openidList,
	}

	log.Println(request)
	// uri := fmt.Sprintf(batchblacklistURL, accessToken)
	return
}

// BatchunBlackList 取消拉黑用户
// func (user *User) BatchunBlackList() (err error) {
// 	var accessToken string
// 	accessToken, err = user.GetAccessToken()
// 	if err != nil {
// 		return
// 	}

// 	uri := fmt.Sprintf(batchblacklistURL, accessToken)
// }
