package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

type followerUserResponse struct {
	util.CommonError
	FollowUser []string `json:"follow_user"`
}

//获取配置了客户联系功能的成员列表
//@see https://developer.work.weixin.qq.com/document/path/92571
func (r *Client) GetFollowUserList() ([]string, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=%s"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.HTTPGet(fmt.Sprintf(requestUrl, accessToken))
	if err != nil {
		return nil, err
	}
	var result followerUserResponse
	err = json.Unmarshal(response, &result)
	if result.ErrCode != 0 {
		err = fmt.Errorf("get_follow_user_list error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return result.FollowUser, nil
}
