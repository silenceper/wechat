package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingCreateURL 创建预约会议
	meetingCreateURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/create?access_token=%s"
	// meetingUpdateURL 修改预约会议
	meetingUpdateURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/update?access_token=%s"
	// meetingCancelURL 取消预约会议
	meetingCancelURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/cancel?access_token=%s"
	// meetingGetInfoURL 获取会议详情
	meetingGetInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_info?access_token=%s"
	// meetingGetUserMeetingIDURL 获取成员会议ID列表
	meetingGetUserMeetingIDURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_user_meetingid?access_token=%s"
)

type (
	// MeetingCreateRequest 创建预约会议请求
	MeetingCreateRequest struct {
		AdminUserID     string     `json:"admin_userid"`
		Title           string     `json:"title"`
		MeetingStart    uint32     `json:"meeting_start,omitempty"`
		MeetingDuration uint32     `json:"meeting_duration,omitempty"`
		Description     string     `json:"description,omitempty"`
		Location        string     `json:"location,omitempty"`
		AgentID         uint32     `json:"agentid,omitempty"`
		Invitees        *Invitees  `json:"invitees,omitempty"`
		Guests          []*Guest   `json:"guests,omitempty"`
		Settings        *Settings  `json:"settings,omitempty"`
		CalID           string     `json:"cal_id,omitempty"`
		Reminders       *Reminders `json:"reminders,omitempty"`
	}

	// MeetingCreateResponse 创建预约会议响应
	MeetingCreateResponse struct {
		util.CommonError
		MeetingID   string   `json:"meetingid"`
		ExcessUsers []string `json:"excess_users"`
		MeetingCode string   `json:"meeting_code"`
		MeetingLink string   `json:"meeting_link"`
	}

	// MeetingUpdateRequest 修改预约会议请求
	MeetingUpdateRequest struct {
		MeetingID       string     `json:"meetingid"`
		Title           string     `json:"title,omitempty"`
		MeetingStart    uint32     `json:"meeting_start,omitempty"`
		MeetingDuration uint32     `json:"meeting_duration,omitempty"`
		Description     string     `json:"description,omitempty"`
		Location        string     `json:"location,omitempty"`
		Invitees        *Invitees  `json:"invitees,omitempty"`
		Settings        *Settings  `json:"settings,omitempty"`
		CalID           string     `json:"cal_id,omitempty"`
		Reminders       *Reminders `json:"reminders,omitempty"`
	}

	// MeetingUpdateResponse 修改预约会议响应
	MeetingUpdateResponse struct {
		util.CommonError
		ExcessUsers []string `json:"excess_users"`
	}

	// MeetingCancelRequest 取消预约会议请求
	MeetingCancelRequest struct {
		MeetingID    string `json:"meetingid"`
		SubMeetingID string `json:"sub_meetingid,omitempty"`
	}

	// MeetingGetInfoRequest 获取会议详情请求
	MeetingGetInfoRequest struct {
		MeetingID    string `json:"meetingid,omitempty"`
		MeetingCode  string `json:"meeting_code,omitempty"`
		SubMeetingID string `json:"sub_meetingid,omitempty"`
	}

	// MeetingGetInfoResponse 获取会议详情响应
	MeetingGetInfoResponse struct {
		util.CommonError
		AdminUserID         string           `json:"admin_userid"`
		Title               string           `json:"title"`
		MeetingStart        uint32           `json:"meeting_start"`
		MeetingDuration     uint32           `json:"meeting_duration"`
		Description         string           `json:"description"`
		Location            string           `json:"location"`
		MainDepartment      uint32           `json:"main_department"`
		Status              uint32           `json:"status"`
		MeetingType         uint32           `json:"meeting_type"`
		Attendees           *Attendees       `json:"attendees"`
		Guests              []*Guest         `json:"guests"`
		Settings            *Settings        `json:"settings"`
		CalID               string           `json:"cal_id"`
		Reminders           *Reminders       `json:"reminders"`
		MeetingCode         string           `json:"meeting_code"`
		MeetingLink         string           `json:"meeting_link"`
		HasVote             bool             `json:"has_vote"`
		SubMeetings         []*SubMeeting    `json:"sub_meetings"`
		HasMoreSubMeeting   uint32           `json:"has_more_sub_meeting"`
		RemainSubMeetings   uint32           `json:"remain_sub_meetings"`
		CurrentSubMeetingID string           `json:"current_sub_meetingid"`
		SubRepeatList       []*SubRepeatInfo `json:"sub_repeat_list"`
	}

	// GetUserMeetingIDRequest 获取成员会议ID列表请求
	GetUserMeetingIDRequest struct {
		UserID    string `json:"userid"`
		Cursor    string `json:"cursor,omitempty"`
		Limit     uint32 `json:"limit,omitempty"`
		BeginTime uint32 `json:"begin_time,omitempty"`
		EndTime   uint32 `json:"end_time,omitempty"`
	}

	// GetUserMeetingIDResponse 获取成员会议ID列表响应
	GetUserMeetingIDResponse struct {
		util.CommonError
		NextCursor    string   `json:"next_cursor"`
		MeetingIDList []string `json:"meetingid_list"`
	}
)

// MeetingCreate 创建预约会议
// see https://developer.work.weixin.qq.com/document/path/98148
func (r *Client) MeetingCreate(req *MeetingCreateRequest) (*MeetingCreateResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingCreateURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &MeetingCreateResponse{}
	err = util.DecodeWithError(response, result, "MeetingCreate")
	return result, err
}

// MeetingUpdate 修改预约会议
// see https://developer.work.weixin.qq.com/document/path/98154
func (r *Client) MeetingUpdate(req *MeetingUpdateRequest) (*MeetingUpdateResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingUpdateURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &MeetingUpdateResponse{}
	err = util.DecodeWithError(response, result, "MeetingUpdate")
	return result, err
}

// MeetingCancel 取消预约会议
// see https://developer.work.weixin.qq.com/document/path/98153
func (r *Client) MeetingCancel(req *MeetingCancelRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingCancelURL, accessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "MeetingCancel")
}

// MeetingGetInfo 获取会议详情
// see https://developer.work.weixin.qq.com/document/path/98149
func (r *Client) MeetingGetInfo(req *MeetingGetInfoRequest) (*MeetingGetInfoResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetInfoURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &MeetingGetInfoResponse{}
	err = util.DecodeWithError(response, result, "MeetingGetInfo")
	return result, err
}

// GetUserMeetingID 获取成员会议ID列表
// see https://developer.work.weixin.qq.com/document/path/98714
func (r *Client) GetUserMeetingID(req *GetUserMeetingIDRequest) (*GetUserMeetingIDResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetUserMeetingIDURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetUserMeetingIDResponse{}
	err = util.DecodeWithError(response, result, "GetUserMeetingID")
	return result, err
}
