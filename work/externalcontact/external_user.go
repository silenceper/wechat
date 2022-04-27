package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

type ExternalUserListResponse struct {
	util.CommonError
	ExternalUserId []string `json:"external_userid"`
}

// 获取客户列表
// @see https://developer.work.weixin.qq.com/document/path/92113
func (r *Client) GetExternalUserList(userId string) ([]string, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%v&userid=%v"
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.HTTPGet(fmt.Sprintf(requestUrl, accessToken, userId))
	if err != nil {
		return nil, err
	}
	var result ExternalUserListResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get_external_user_list error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return result.ExternalUserId, nil
}

type ExternalUserDetailResponse struct {
	util.CommonError
	ExternalUser
}

type ExternalUser struct {
	ExternalUserId  string       `json:"external_userid"`
	Name            string       `json:"name"`
	Avatar          string       `json:"avatar"`
	Type            int64        `json:"type"`
	Gender          int64        `json:"gender"`
	UnionId         string       `json:"unionid"`
	Position        string       `json:"position"`
	CorpName        string       `json:"corp_name"`
	CorpFullName    string       `json:"corp_full_name"`
	ExternalProfile string       `json:"external_profile"`
	FollowUser      []FollowUser `json:"follow_user"`
	NextCursor      string       `json:"next_cursor"`
}

type FollowUser struct {
	UserId         string        `json:"userid"`
	Remark         string        `json:"remark"`
	Description    string        `json:"description"`
	CreateTime     string        `json:"create_time"`
	Tags           []Tag         `json:"tags"`
	RemarkCorpName string        `json:"remark_corp_name"`
	RemarkMobiles  []string      `json:"remark_mobiles"`
	OperUserId     string        `json:"oper_userid"`
	AddWay         int64         `json:"add_way"`
	WeChatChannels WechatChannel `json:"wechat_channels"`
	State          string        `json:"state"`
}

type Tag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	Type      int64  `json:"type"`
	TagId     string `json:"tag_id"`
}

type WechatChannel struct {
	NickName string `json:"nickname"`
	Source   string `json:"source"`
}

func (r *Client) GetExternalUserDetail(externalUserId string, nextCursor ...string) (*ExternalUser, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%v&external_userid=%v&cursor=%v"
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	response, err = util.HTTPGet(fmt.Sprintf(requestUrl, accessToken, externalUserId, nextCursor))
	if err != nil {
		return nil, err
	}
	var result ExternalUserDetailResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get_external_user_detail error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return &result.ExternalUser, nil
}

type BatchGetExternalUserDetailsRequest struct {
	UserIdList []string `json:"userid_list"`
	Cursor     string   `json:"cursor"`
}

type ExternalUserDetailListResponse struct {
	util.CommonError
	ExternalContactList []ExternalUser `json:"external_contact_list"`
}

func (r *Client) BatchGetExternalUserDetails(request BatchGetExternalUserDetailsRequest) ([]ExternalUser, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=%v"
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, _ := json.Marshal(request)
	response, err = util.HTTPPost(fmt.Sprintf(requestUrl, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	var result ExternalUserDetailListResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get_external_user_detail error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return result.ExternalContactList, nil
}

type UpdateUserRemarkRequest struct {
	UserId           string   `json:"userid"`
	ExternalUserId   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaid string   `json:"remark_pic_mediaid"`
}

func (r *Client) UpdateUserRemark(request UpdateUserRemarkRequest) error {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=%v"
	accessToken, err := r.ctx.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(request)
	response, err = util.HTTPPost(fmt.Sprintf(requestUrl, accessToken), string(jsonData))
	if err != nil {
		return err
	}
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get_external_user_detail error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err
	}
	return nil
}
