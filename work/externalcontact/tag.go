package externalcontact

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

type GetCropTagRequest struct {
	TagId   []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
}

type GetCropTagListResponse struct {
	util.CommonError
	TagGroup []TagGroup `json:"tag_group"`
}

type TagGroup struct {
	GroupId    string            `json:"group_id"`
	GroupName  string            `json:"group_name"`
	CreateTime string            `json:"create_time"`
	GroupOrder int               `json:"group_order"`
	Deleted    bool              `json:"deleted"`
	Tag        []TagGroupTagItem `json:"tag"`
}

type TagGroupTagItem struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

// 获取企业标签库
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) GetCropTagList(req GetCropTagRequest) ([]TagGroup, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestUrl, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	var result GetCropTagListResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCropTagList error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return result.TagGroup, nil
}

type AddCropTagRequest struct {
	GroupId   string           `json:"group_id"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	Tag       []AddCropTagItem `json:"tag"`
	AgentId   int              `json:"agentid"`
}

type AddCropTagItem struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

type AddCropTagResponse struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group"`
}

// 添加企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) AddCropTag(req AddCropTagRequest) (*TagGroup, error) {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestUrl, accessToken), string(jsonData))
	if err != nil {
		return nil, err
	}
	var result AddCropTagResponse
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add_corp_tag error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return nil, err
	}
	return &result.TagGroup, nil
}

type EditCropTagRequest struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	AgentId string `json:"agent_id"`
}

// 修改企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) EditCropTag(req EditCropTagRequest) error {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
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
		err = fmt.Errorf("edit_corp_tag error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err

	}
	return nil
}

type DeleteCropTagRequest struct {
	TagId   []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
	AgentId string   `json:"agent_id"`
}

// 删除企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) DeleteCropTag(req DeleteCropTagRequest) error {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
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
		err = fmt.Errorf("del_corp_tag error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err

	}
	return nil
}

type MarkTagRequest struct {
	UserId         string   `json:"user_id"`
	ExternalUserId string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

// 为客户打上标签
// @see https://developer.work.weixin.qq.com/document/path/92118
func (r *Client) MarkTag(request MarkTagRequest) error {
	var accessToken string
	var requestUrl = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
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
		err = fmt.Errorf("mark_tag error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err

	}
	return nil
}
