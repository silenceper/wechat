package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingGetGuestsURL 获取会议嘉宾列表
	meetingGetGuestsURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_guests?access_token=%s"
	// meetingSetGuestsURL 更新会议嘉宾列表
	meetingSetGuestsURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/set_guests?access_token=%s"
)

type (
	// GetGuestsRequest 获取会议嘉宾列表请求
	GetGuestsRequest struct {
		MeetingID string `json:"meetingid"`
	}

	// GetGuestsResponse 获取会议嘉宾列表响应
	GetGuestsResponse struct {
		util.CommonError
		MeetingID   string   `json:"meetingid"`
		MeetingCode string   `json:"meeting_code"`
		Title       string   `json:"title"`
		Guests      []*Guest `json:"guests"`
	}

	// SetGuestsRequest 更新会议嘉宾列表请求
	SetGuestsRequest struct {
		MeetingID string   `json:"meetingid"`
		Guests    []*Guest `json:"guests"`
	}
)

// GetGuests 获取会议嘉宾列表
// see https://developer.work.weixin.qq.com/document/path/99039
func (r *Client) GetGuests(req *GetGuestsRequest) (*GetGuestsResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetGuestsURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetGuestsResponse{}
	err = util.DecodeWithError(response, result, "GetGuests")
	return result, err
}

// SetGuests 更新会议嘉宾列表
// see https://developer.work.weixin.qq.com/document/path/99040
func (r *Client) SetGuests(req *SetGuestsRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingSetGuestsURL, accessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "SetGuests")
}
