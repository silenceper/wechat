package meeting

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// meetingEnrollSetConfigURL 修改会议报名配置
	meetingEnrollSetConfigURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/set_config?access_token=%s"
	// meetingEnrollGetConfigURL 获取会议报名配置
	meetingEnrollGetConfigURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/get_config?access_token=%s"
	// meetingEnrollQueryByTmpOpenIDURL 获取会议成员报名ID
	meetingEnrollQueryByTmpOpenIDURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/query_by_tmp_openid?access_token=%s"
	// meetingEnrollListURL 获取会议报名信息
	meetingEnrollListURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/list?access_token=%s"
	// meetingEnrollApproveURL 审批会议报名信息
	meetingEnrollApproveURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/approve?access_token=%s"
	// meetingEnrollImportURL 导入会议报名信息
	meetingEnrollImportURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/import?access_token=%s"
	// meetingEnrollDeleteURL 删除会议报名信息
	meetingEnrollDeleteURL = "https://qyapi.weixin.qq.com/cgi-bin/meeting/enroll/delete?access_token=%s"
)

type (
	// EnrollSetConfigRequest 修改会议报名配置请求
	EnrollSetConfigRequest struct {
		MeetingID                    string            `json:"meetingid"`
		ApproveType                  int32             `json:"approve_type"`
		IsCollectQuestion            int32             `json:"is_collect_question"`
		QuestionList                 []*EnrollQuestion `json:"question_list"`
		NoRegistrationNeededForStaff *bool             `json:"no_registration_needed_for_staff,omitempty"`
	}

	// EnrollQuestion 报名问题
	EnrollQuestion struct {
		IsRequired    int32           `json:"is_required,omitempty"`
		OptionList    []*EnrollOption `json:"option_list,omitempty"`
		QuestionTitle string          `json:"question_title,omitempty"`
		QuestionType  int32           `json:"question_type,omitempty"`
		SpecialType   int32           `json:"special_type,omitempty"`
	}

	// EnrollOption 报名问题选项
	EnrollOption struct {
		Content string `json:"content"`
	}

	// EnrollSetConfigResponse 修改会议报名配置响应
	EnrollSetConfigResponse struct {
		util.CommonError
		QuestionCount int32 `json:"question_count"`
	}

	// EnrollGetConfigRequest 获取会议报名配置请求
	EnrollGetConfigRequest struct {
		MeetingID string `json:"meetingid"`
	}

	// EnrollGetConfigResponse 获取会议报名配置响应
	EnrollGetConfigResponse struct {
		util.CommonError
		ApproveType                  int32             `json:"approve_type"`
		IsCollectQuestion            int32             `json:"is_collect_question"`
		NoRegistrationNeededForStaff bool              `json:"no_registration_needed_for_staff"`
		QuestionList                 []*EnrollQuestion `json:"question_list"`
	}

	// EnrollQueryByTmpOpenIDRequest 获取会议成员报名ID请求
	EnrollQueryByTmpOpenIDRequest struct {
		MeetingID     string   `json:"meetingid"`
		SortingRules  int32    `json:"sorting_rules,omitempty"`
		TmpOpenIDList []string `json:"tmp_openid_list"`
	}

	// EnrollQueryByTmpOpenIDResponse 获取会议成员报名ID响应
	EnrollQueryByTmpOpenIDResponse struct {
		util.CommonError
		EnrollIDList []*EnrollIDItem `json:"enroll_id_list"`
	}

	// EnrollIDItem 报名ID项
	EnrollIDItem struct {
		TmpOpenID string `json:"tmp_openid"`
		EnrollID  string `json:"enroll_id"`
	}

	// EnrollListRequest 获取会议报名信息请求
	EnrollListRequest struct {
		MeetingID string `json:"meetingid"`
		Status    int32  `json:"status,omitempty"`
		Cursor    string `json:"cursor,omitempty"`
		Limit     int32  `json:"limit,omitempty"`
	}

	// EnrollListResponse 获取会议报名信息响应
	EnrollListResponse struct {
		util.CommonError
		HasMore    bool          `json:"has_more"`
		NextCursor string        `json:"next_cursor"`
		EnrollList []*EnrollInfo `json:"enroll_list"`
	}

	// EnrollInfo 报名信息
	EnrollInfo struct {
		EnrollID         string          `json:"enroll_id"`
		EnrollTime       string          `json:"enroll_time"`
		EnrollSourceType int32           `json:"enroll_source_type"`
		NickName         string          `json:"nick_name"`
		Status           int32           `json:"status"`
		UserID           string          `json:"userid"`
		TmpOpenID        string          `json:"tmp_openid"`
		EnrollCode       string          `json:"enroll_code"`
		AnswerList       []*EnrollAnswer `json:"answer_list"`
	}

	// EnrollAnswer 报名答题
	EnrollAnswer struct {
		AnswerContent []string `json:"answer_content"`
		IsRequired    int32    `json:"is_required"`
		QuestionNum   int32    `json:"question_num"`
		QuestionTitle string   `json:"question_title"`
		QuestionType  int32    `json:"question_type"`
		SpecialType   int32    `json:"special_type"`
	}

	// EnrollApproveRequest 审批会议报名信息请求
	EnrollApproveRequest struct {
		MeetingID    string   `json:"meetingid"`
		Action       int32    `json:"action"`
		EnrollIDList []string `json:"enroll_id_list"`
	}

	// EnrollApproveResponse 审批会议报名信息响应
	EnrollApproveResponse struct {
		util.CommonError
		HandledCount int32 `json:"handled_count"`
	}

	// EnrollImportRequest 导入会议报名信息请求
	EnrollImportRequest struct {
		MeetingID  string           `json:"meetingid"`
		EnrollList []*EnrollRequest `json:"enroll_list"`
	}

	// EnrollRequest 报名成员请求
	EnrollRequest struct {
		UserID      string `json:"userid,omitempty"`
		Area        string `json:"area,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		NickName    string `json:"nick_name,omitempty"`
	}

	// EnrollImportResponse 导入会议报名信息响应
	EnrollImportResponse struct {
		util.CommonError
		TotalCount uint32            `json:"total_count"`
		EnrollList []*EnrollResponse `json:"enroll_list"`
	}

	// EnrollResponse 报名成员响应
	EnrollResponse struct {
		EnrollID    string `json:"enroll_id"`
		UserID      string `json:"userid"`
		Area        string `json:"area"`
		PhoneNumber string `json:"phone_number"`
		NickName    string `json:"nick_name"`
		EnrollCode  string `json:"enroll_code"`
	}

	// EnrollDeleteRequest 删除会议报名信息请求
	EnrollDeleteRequest struct {
		MeetingID    string            `json:"meetingid"`
		EnrollIDList []*EnrollDeleteID `json:"enroll_id_list"`
	}

	// EnrollDeleteID 报名ID
	EnrollDeleteID struct {
		EnrollID string `json:"enroll_id"`
	}

	// EnrollDeleteResponse 删除会议报名信息响应
	EnrollDeleteResponse struct {
		util.CommonError
		TotalCount uint32 `json:"total_count"`
	}
)

// EnrollSetConfig 修改会议报名配置
// see https://developer.work.weixin.qq.com/document/path/98797
func (r *Client) EnrollSetConfig(req *EnrollSetConfigRequest) (*EnrollSetConfigResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollSetConfigURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollSetConfigResponse{}
	err = util.DecodeWithError(response, result, "EnrollSetConfig")
	return result, err
}

// EnrollGetConfig 获取会议报名配置
// see https://developer.work.weixin.qq.com/document/path/98800
func (r *Client) EnrollGetConfig(req *EnrollGetConfigRequest) (*EnrollGetConfigResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollGetConfigURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollGetConfigResponse{}
	err = util.DecodeWithError(response, result, "EnrollGetConfig")
	return result, err
}

// EnrollQueryByTmpOpenID 获取会议成员报名ID
// see https://developer.work.weixin.qq.com/document/path/98794
func (r *Client) EnrollQueryByTmpOpenID(req *EnrollQueryByTmpOpenIDRequest) (*EnrollQueryByTmpOpenIDResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollQueryByTmpOpenIDURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollQueryByTmpOpenIDResponse{}
	err = util.DecodeWithError(response, result, "EnrollQueryByTmpOpenID")
	return result, err
}

// EnrollList 获取会议报名信息
// see https://developer.work.weixin.qq.com/document/path/98810
func (r *Client) EnrollList(req *EnrollListRequest) (*EnrollListResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollListURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollListResponse{}
	err = util.DecodeWithError(response, result, "EnrollList")
	return result, err
}

// EnrollApprove 审批会议报名信息
// see https://developer.work.weixin.qq.com/document/path/98807
func (r *Client) EnrollApprove(req *EnrollApproveRequest) (*EnrollApproveResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollApproveURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollApproveResponse{}
	err = util.DecodeWithError(response, result, "EnrollApprove")
	return result, err
}

// EnrollImport 导入会议报名信息
// see https://developer.work.weixin.qq.com/document/path/98816
func (r *Client) EnrollImport(req *EnrollImportRequest) (*EnrollImportResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollImportURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollImportResponse{}
	err = util.DecodeWithError(response, result, "EnrollImport")
	return result, err
}

// EnrollDelete 删除会议报名信息
// see https://developer.work.weixin.qq.com/document/path/98817
func (r *Client) EnrollDelete(req *EnrollDeleteRequest) (*EnrollDeleteResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(meetingEnrollDeleteURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &EnrollDeleteResponse{}
	err = util.DecodeWithError(response, result, "EnrollDelete")
	return result, err
}
