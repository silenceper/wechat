package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingGetInviteesURL 获取会议受邀成员列表
	meetingGetInviteesURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_invitees?access_token=%s"
	// meetingSetInviteesURL 更新会议受邀成员列表
	meetingSetInviteesURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/set_invitees?access_token=%s"
)

type (
	// GetInviteesRequest 获取会议受邀成员列表请求
	GetInviteesRequest struct {
		MeetingID string `json:"meetingid"`
		Cursor    string `json:"cursor,omitempty"`
	}

	// GetInviteesResponse 获取会议受邀成员列表响应
	GetInviteesResponse struct {
		util.CommonError
		HasMore    bool       `json:"has_more"`
		NextCursor string     `json:"next_cursor"`
		Invitees   []*Invitee `json:"invitees"`
	}

	// Invitee 受邀成员
	Invitee struct {
		UserID string `json:"userid"`
	}

	// SetInviteesRequest 更新会议受邀成员列表请求
	SetInviteesRequest struct {
		MeetingID string     `json:"meetingid"`
		Invitees  []*Invitee `json:"invitees"`
	}
)

// GetInvitees 获取会议受邀成员列表
// see https://developer.work.weixin.qq.com/document/path/98160
func (r *Client) GetInvitees(req *GetInviteesRequest) (*GetInviteesResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetInviteesURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetInviteesResponse{}
	err = util.DecodeWithError(response, result, "GetInvitees")
	return result, err
}

// SetInvitees 更新会议受邀成员列表
// see https://developer.work.weixin.qq.com/document/path/98162
func (r *Client) SetInvitees(req *SetInviteesRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingSetInviteesURL, accessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "SetInvitees")
}
