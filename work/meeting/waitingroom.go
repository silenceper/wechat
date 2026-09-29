package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingWaitingRoomGetCurrentUserListURL 获取实时等候室成员列表
	meetingWaitingRoomGetCurrentUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/waitingroom/get_current_user_list?access_token=%s"
	// meetingWaitingRoomGetUserListURL 获取等候室成员记录
	meetingWaitingRoomGetUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/waitingroom/get_user_list?access_token=%s"
)

type (
	// WaitingRoomGetCurrentUserListRequest 获取实时等候室成员列表请求
	WaitingRoomGetCurrentUserListRequest struct {
		MeetingID string `json:"meetingid"`
		Limit     uint32 `json:"limit,omitempty"`
		Cursor    string `json:"cursor,omitempty"`
	}

	// WaitingRoomGetCurrentUserListResponse 获取实时等候室成员列表响应
	WaitingRoomGetCurrentUserListResponse struct {
		util.CommonError
		HasMore    bool               `json:"has_more"`
		NextCursor string             `json:"next_cursor"`
		UserList   []*WaitingRoomUser `json:"user_list"`
	}

	// WaitingRoomGetUserListRequest 获取等候室成员记录请求
	WaitingRoomGetUserListRequest struct {
		MeetingID string `json:"meetingid"`
		Limit     uint32 `json:"limit,omitempty"`
		Cursor    string `json:"cursor,omitempty"`
	}

	// WaitingRoomGetUserListResponse 获取等候室成员记录响应
	WaitingRoomGetUserListResponse struct {
		util.CommonError
		HasMore    bool                 `json:"has_more"`
		NextCursor string               `json:"next_cursor"`
		UserList   []*WaitingRoomRecord `json:"user_list"`
	}

	// WaitingRoomUser 实时等候室成员
	WaitingRoomUser struct {
		UserID       string `json:"userid"`
		InstanceID   uint32 `json:"instance_id"`
		CustomerData string `json:"customer_data"`
		TmpOpenID    string `json:"tmp_openid"`
	}

	// WaitingRoomRecord 等候室成员记录
	WaitingRoomRecord struct {
		UserID     string `json:"userid"`
		TmpOpenID  string `json:"tmp_openid"`
		InstanceID uint32 `json:"instance_id"`
		JoinTime   int64  `json:"join_time"`
		QuitTime   int64  `json:"quit_time"`
	}
)

// WaitingRoomGetCurrentUserList 获取实时等候室成员列表
// see https://developer.work.weixin.qq.com/document/path/98163
func (r *Client) WaitingRoomGetCurrentUserList(req *WaitingRoomGetCurrentUserListRequest) (*WaitingRoomGetCurrentUserListResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingWaitingRoomGetCurrentUserListURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &WaitingRoomGetCurrentUserListResponse{}
	err = util.DecodeWithError(response, result, "WaitingRoomGetCurrentUserList")
	return result, err
}

// WaitingRoomGetUserList 获取等候室成员记录
// see https://developer.work.weixin.qq.com/document/path/98164
func (r *Client) WaitingRoomGetUserList(req *WaitingRoomGetUserListRequest) (*WaitingRoomGetUserListResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingWaitingRoomGetUserListURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &WaitingRoomGetUserListResponse{}
	err = util.DecodeWithError(response, result, "WaitingRoomGetUserList")
	return result, err
}
