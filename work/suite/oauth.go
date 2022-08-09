package suite

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetUserinfo3rdURL 获取访问用户身份
	GetUserinfo3rdURL = "https://qyapi.weixin.qq.com/cgi-bin/service/getuserinfo3rd?suite_access_token=%s&code=%s"
	// GetUserDetail3rdURL 获取访问用户敏感信息
	GetUserDetail3rdURL = "https://qyapi.weixin.qq.com/cgi-bin/service/getuserdetail3rd?suite_access_token=%s"
)

type (
	// GetUserinfo3rdResponse 获取访问用户身份响应
	GetUserinfo3rdResponse struct {
		util.CommonError
		CorpID     string `json:"CorpId"`
		UserID     string `json:"UserId"`
		DeviceID   string `json:"DeviceId"`
		UserTicket string `json:"user_ticket"`
		ExpiresIn  int    `json:"expires_in"`
		OpenUserid string `json:"open_userid"`
		OpenID     string `json:"OpenId"`
	}
)

// GetUserinfo3rd 获取访问用户身份
// @see https://developer.work.weixin.qq.com/document/path/91121
func (r *Client) GetUserinfo3rd(code string) (*GetUserinfo3rdResponse, error) {
	var (
		response []byte
		err      error
	)
	response, err = util.HTTPGet(fmt.Sprintf(GetUserinfo3rdURL, r.SuiteAccessToken, code))
	if err != nil {
		return nil, err
	}
	result := &GetUserinfo3rdResponse{}
	err = util.DecodeWithError(response, result, "GetUserinfo3rd")
	if err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetUserDetail3rdRequest 获取访问用户敏感信息请求
	GetUserDetail3rdRequest struct {
		UserTicket string `json:"user_ticket"`
	}
	// GetUserDetail3rdResponse 获取访问用户敏感信息响应
	GetUserDetail3rdResponse struct {
		util.CommonError
		CorpID string `json:"corpid"`
		UserID string `json:"userid"`
		Name   string `json:"name"`
		Gender string `json:"gender"`
		Avatar string `json:"avatar"`
		QrCode string `json:"qr_code"`
	}
)

// GetUserDetail3rd 获取访问用户敏感信息
// see https://developer.work.weixin.qq.com/document/path/91122
func (r *Client) GetUserDetail3rd(request *GetUserDetail3rdRequest) (*GetUserDetail3rdResponse, error) {
	var (
		response []byte
		err      error
	)
	jsonData, _ := json.Marshal(request)
	response, err = util.HTTPPost(fmt.Sprintf(GetUserDetail3rdURL, r.SuiteAccessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	result := &GetUserDetail3rdResponse{}
	err = util.DecodeWithError(response, &result, "GetUserDetail3rd")
	if err != nil {
		return nil, err
	}
	return result, nil
}
