package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingGetRealtimeAttendeeListURL 获取实时会中成员列表
	meetingGetRealtimeAttendeeListURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_realtime_attendee_list?access_token=%s"
	// meetingGetAttendeeListURL 获取已参会成员列表
	meetingGetAttendeeListURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_attendee_list?access_token=%s"
	// meetingCheckDeviceInMeetingURL 获取成员设备是否入会
	meetingCheckDeviceInMeetingURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/check_device_in_meeting?access_token=%s"
	// meetingGetQualityURL 获取会议健康度
	meetingGetQualityURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/get_quality?access_token=%s"
)

type (
	// GetRealtimeAttendeeListRequest 获取实时会中成员列表请求
	GetRealtimeAttendeeListRequest struct {
		MeetingID    string `json:"meetingid"`
		SubMeetingID string `json:"sub_meetingid,omitempty"`
		Cursor       string `json:"cursor,omitempty"`
		Limit        uint32 `json:"limit,omitempty"`
	}

	// GetRealtimeAttendeeListResponse 获取实时会中成员列表响应
	GetRealtimeAttendeeListResponse struct {
		util.CommonError
		Status     string              `json:"status"`
		HasMore    bool                `json:"has_more"`
		NextCursor string              `json:"next_cursor"`
		Attendees  []*RealtimeAttendee `json:"attendees"`
	}

	// RealtimeAttendee 实时会中成员
	RealtimeAttendee struct {
		UserID            string `json:"userid"`
		TmpOpenID         string `json:"tmp_openid"`
		JoinTime          string `json:"join_time"`
		InstanceID        uint32 `json:"instance_id"`
		Role              uint32 `json:"role"`
		JoinType          uint32 `json:"join_type"`
		AudioState        bool   `json:"audio_state"`
		VideoState        bool   `json:"video_state"`
		ScreenSharedState bool   `json:"screen_shared_state"`
	}

	// GetAttendeeListRequest 获取已参会成员列表请求
	GetAttendeeListRequest struct {
		MeetingID    string `json:"meetingid"`
		SubMeetingID string `json:"sub_meetingid,omitempty"`
		StartTime    uint32 `json:"start_time,omitempty"`
		EndTime      uint32 `json:"end_time,omitempty"`
		Cursor       string `json:"cursor,omitempty"`
		Limit        uint32 `json:"limit,omitempty"`
	}

	// GetAttendeeListResponse 获取已参会成员列表响应
	GetAttendeeListResponse struct {
		util.CommonError
		HasMore    bool              `json:"has_more"`
		NextCursor string            `json:"next_cursor"`
		Attendees  []*AttendedMember `json:"attendees"`
	}

	// AttendedMember 已参会成员
	AttendedMember struct {
		UserID            string `json:"userid"`
		TmpOpenID         string `json:"tmp_openid"`
		JoinTime          string `json:"join_time"`
		QuitTime          string `json:"quit_time"`
		InstanceID        uint32 `json:"instance_id"`
		Role              uint32 `json:"role"`
		WebinarRole       uint32 `json:"webinar_role"`
		JoinType          uint32 `json:"join_type"`
		Net               string `json:"net"`
		AudioState        bool   `json:"audio_state"`
		VideoState        bool   `json:"video_state"`
		ScreenSharedState bool   `json:"screen_shared_state"`
		CustomerData      string `json:"customer_data"`
	}

	// CheckDeviceInMeetingRequest 获取成员设备是否入会请求
	CheckDeviceInMeetingRequest struct {
		UserID         string   `json:"userid"`
		InstanceIDList []uint32 `json:"instance_id_list,omitempty"`
		MeetingIDList  []string `json:"meetingid_list"`
	}

	// CheckDeviceInMeetingResponse 获取成员设备是否入会响应
	CheckDeviceInMeetingResponse struct {
		util.CommonError
		ResultList []*DeviceMeetingResult `json:"result_list"`
	}

	// DeviceMeetingResult 设备会议结果
	DeviceMeetingResult struct {
		MeetingID  string `json:"meetingid"`
		InstanceID uint32 `json:"instance_id"`
	}

	// GetQualityRequest 获取会议健康度请求
	GetQualityRequest struct {
		MeetingID    string `json:"meetingid"`
		SubMeetingID string `json:"sub_meetingid,omitempty"`
		StartTime    uint32 `json:"start_time,omitempty"`
		Cursor       string `json:"cursor,omitempty"`
		Limit        uint32 `json:"limit,omitempty"`
	}

	// GetQualityResponse 获取会议健康度响应
	GetQualityResponse struct {
		util.CommonError
		Quality            int32              `json:"quality"`
		AudioQuality       int32              `json:"audio_quality"`
		VideoQuality       int32              `json:"video_quality"`
		ScreenShareQuality int32              `json:"screen_share_quality"`
		NetworkQuality     int32              `json:"network_quality"`
		Problems           []string           `json:"problems"`
		Attendees          []*QualityAttendee `json:"attendees"`
		NextCursor         string             `json:"next_cursor"`
		HasMore            bool               `json:"has_more"`
	}

	// QualityAttendee 参会人员健康度
	QualityAttendee struct {
		UserID             string   `json:"userid"`
		TmpOpenID          string   `json:"tmp_openid"`
		InstanceID         uint32   `json:"instance_id"`
		Quality            int32    `json:"quality"`
		AudioQuality       int32    `json:"audio_quality"`
		VideoQuality       int32    `json:"video_quality"`
		ScreenShareQuality int32    `json:"screen_share_quality"`
		NetworkQuality     int32    `json:"network_quality"`
		Problems           []string `json:"problems"`
	}
)

// GetRealtimeAttendeeList 获取实时会中成员列表
// see https://developer.work.weixin.qq.com/document/path/98157
func (r *Client) GetRealtimeAttendeeList(req *GetRealtimeAttendeeListRequest) (*GetRealtimeAttendeeListResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetRealtimeAttendeeListURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetRealtimeAttendeeListResponse{}
	err = util.DecodeWithError(response, result, "GetRealtimeAttendeeList")
	return result, err
}

// GetAttendeeList 获取已参会成员列表
// see https://developer.work.weixin.qq.com/document/path/98156
func (r *Client) GetAttendeeList(req *GetAttendeeListRequest) (*GetAttendeeListResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetAttendeeListURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetAttendeeListResponse{}
	err = util.DecodeWithError(response, result, "GetAttendeeList")
	return result, err
}

// CheckDeviceInMeeting 获取成员设备是否入会
// see https://developer.work.weixin.qq.com/document/path/98165
func (r *Client) CheckDeviceInMeeting(req *CheckDeviceInMeetingRequest) (*CheckDeviceInMeetingResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingCheckDeviceInMeetingURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &CheckDeviceInMeetingResponse{}
	err = util.DecodeWithError(response, result, "CheckDeviceInMeeting")
	return result, err
}

// GetQuality 获取会议健康度
// see https://developer.work.weixin.qq.com/document/path/98821
func (r *Client) GetQuality(req *GetQualityRequest) (*GetQualityResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingGetQualityURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetQualityResponse{}
	err = util.DecodeWithError(response, result, "GetQuality")
	return result, err
}
