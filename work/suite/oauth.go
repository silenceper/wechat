package suite

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// GetUserinfo3rdURL 获取访问用户身份
	GetUserinfo3rdURL = "https://qyapi.weixin.qq.com/cgi-bin/service/getuserinfo3rd?suite_access_token=%s&code=%s"
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
		result   *GetUserinfo3rdResponse
	)
	response, err = util.HTTPGet(fmt.Sprintf(GetUserinfo3rdURL, r.SuiteAccessToken, code))
	if err != nil {
		return nil, err
	}
	err = util.DecodeWithError(response, result, "GetUserinfo3rd")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Client) TestFunc() string {
	return r.SuiteAccessToken + "success"
}
