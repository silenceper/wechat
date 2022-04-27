package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

// GetCropTagRequest 获取企业标签请求
type GetCropTagRequest struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
}

// GetCropTagListResponse 获取企业标签列表响应
type GetCropTagListResponse struct {
	util.CommonError
	TagGroup []TagGroup `json:"tag_group"`
}

// TagGroup 企业标签组
type TagGroup struct {
	GroupID    string            `json:"group_id"`
	GroupName  string            `json:"group_name"`
	CreateTime string            `json:"create_time"`
	GroupOrder int               `json:"group_order"`
	Deleted    bool              `json:"deleted"`
	Tag        []TagGroupTagItem `json:"tag"`
}

// TagGroupTagItem 企业标签内的子项
type TagGroupTagItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int    `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

// GetCropTagList 获取企业标签库
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) GetCropTagList(req GetCropTagRequest) ([]TagGroup, error) {
	var accessToken string
	var requestURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestURL, accessToken), string(jsonData))
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

// AddCropTagRequest 添加企业标签请求
type AddCropTagRequest struct {
	GroupID   string           `json:"group_id"`
	GroupName string           `json:"group_name"`
	Order     int              `json:"order"`
	Tag       []AddCropTagItem `json:"tag"`
	AgentID   int              `json:"agentid"`
}

// AddCropTagItem 添加企业标签子项
type AddCropTagItem struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}

// AddCropTagResponse 添加企业标签响应
type AddCropTagResponse struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group"`
}

// AddCropTag 添加企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) AddCropTag(req AddCropTagRequest) (*TagGroup, error) {
	var accessToken string
	var requestURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return nil, err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestURL, accessToken), string(jsonData))
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

// EditCropTagRequest 编辑客户企业标签请求
type EditCropTagRequest struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	AgentID string `json:"agent_id"`
}

// EditCropTag 修改企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) EditCropTag(req EditCropTagRequest) error {
	var accessToken string
	var requestURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestURL, accessToken), string(jsonData))
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

// DeleteCropTagRequest 删除企业标签请求
type DeleteCropTagRequest struct {
	TagID   []string `json:"tag_id"`
	GroupID []string `json:"group_id"`
	AgentID string   `json:"agent_id"`
}

// DeleteCropTag 删除企业客户标签
// @see https://developer.work.weixin.qq.com/document/path/92117
func (r *Client) DeleteCropTag(req DeleteCropTagRequest) error {
	var accessToken string
	var requestURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(req)
	response, err = util.HTTPPost(fmt.Sprintf(requestURL, accessToken), string(jsonData))
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

// MarkTagRequest 给客户打标签请求
type MarkTagRequest struct {
	UserID         string   `json:"user_id"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag"`
	RemoveTag      []string `json:"remove_tag"`
}

// MarkTag 为客户打上标签
// @see https://developer.work.weixin.qq.com/document/path/92118
func (r *Client) MarkTag(request MarkTagRequest) error {
	var accessToken string
	var requestURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=%v"
	accessToken, err := r.GetAccessToken()
	if err != nil {
		return err
	}
	var response []byte
	jsonData, _ := json.Marshal(request)
	response, err = util.HTTPPost(fmt.Sprintf(requestURL, accessToken), string(jsonData))
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
