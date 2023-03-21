package addresslist

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// createTagURL 创建标签
	createTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s"
	// updateTagURL 更新标签名字
	updateTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s"
	// deleteTagURL 删除标签
	deleteTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d"
)

type (
	// CreateTagRequest 创建标签请求
	CreateTagRequest struct {
		TagName string `json:"tagname"`
		TagID   int    `json:"tagid,omitempty"`
	}
	// CreateTagResponse 创建标签响应
	CreateTagResponse struct {
		util.CommonError
		TagID int `json:"tagid"`
	}
)

// CreateTag 创建标签
// see https://developer.work.weixin.qq.com/document/path/90210
func (r *Client) CreateTag(req *CreateTagRequest) (*CreateTagResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(createTagURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &CreateTagResponse{}
	if err = util.DecodeWithError(response, result, "CreateTag"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// UpdateTagRequest 更新标签名字请求
	UpdateTagRequest struct {
		TagID   int    `json:"tagid"`
		TagName string `json:"tagname"`
	}
)

// UpdateTag 更新标签名字
// see https://developer.work.weixin.qq.com/document/path/90211
func (r *Client) UpdateTag(req *UpdateTagRequest) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(updateTagURL, accessToken), req); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "UpdateTag")
}

// DeleteTag 删除标签
// @see https://developer.work.weixin.qq.com/document/path/90212
func (r *Client) DeleteTag(tagID int) error {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return err
	}
	var response []byte
	if response, err = util.HTTPGet(fmt.Sprintf(deleteTagURL, accessToken, tagID)); err != nil {
		return err
	}
	return util.DecodeWithCommonError(response, "DeleteTag")
}
