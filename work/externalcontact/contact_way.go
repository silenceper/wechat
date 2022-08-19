package externalcontact

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// AddContactWayURL 配置客户联系「联系我」方式
	AddContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=%s"
	// GetContactWayURL 获取企业已配置的「联系我」方式
	GetContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=%s"
)

type (
	// ConclusionsRequest 结束语请求
	ConclusionsRequest struct {
		Text        ConclusionsText         `json:"text"`
		Image       ConclusionsImageRequest `json:"image"`
		Link        ConclusionsLink         `json:"link"`
		MiniProgram ConclusionsMiniProgram  `json:"miniprogram"`
	}
	// ConclusionsText 文本格式结束语
	ConclusionsText struct {
		Content string `json:"content"`
	}
	// ConclusionsImageRequest 图片格式结束语请求
	ConclusionsImageRequest struct {
		MediaID string `json:"media_id"`
	}
	// ConclusionsLink 链接格式结束语
	ConclusionsLink struct {
		Title  string `json:"title"`
		PicURL string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	}
	// ConclusionsMiniProgram 小程序格式结束语
	ConclusionsMiniProgram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		AppID      string `json:"appid"`
		Page       string `json:"page"`
	}
	// ConclusionsResponse 结束语响应
	ConclusionsResponse struct {
		Text        ConclusionsText          `json:"text"`
		Image       ConclusionsImageResponse `json:"image"`
		Link        ConclusionsLink          `json:"link"`
		MiniProgram ConclusionsMiniProgram   `json:"miniprogram"`
	}
	// ConclusionsImageResponse 图片格式结束语响应
	ConclusionsImageResponse struct {
		PicURL string `json:"pic_url"`
	}
)

type (
	// AddContactWayRequest 配置客户联系「联系我」方式请求
	AddContactWayRequest struct {
		Type          int                `json:"type"`
		Scene         int                `json:"scene"`
		Style         int                `json:"style"`
		Remark        string             `json:"remark"`
		SkipVerify    bool               `json:"skip_verify"`
		State         string             `json:"state"`
		User          []string           `json:"user"`
		Party         []int              `json:"party"`
		IsTemp        bool               `json:"is_temp"`
		ExpiresIn     int                `json:"expires_in"`
		ChatExpiresIn int                `json:"chat_expires_in"`
		UnionID       string             `json:"unionid"`
		Conclusions   ConclusionsRequest `json:"conclusions"`
	}
	// AddContactWayResponse 配置客户联系「联系我」方式响应
	AddContactWayResponse struct {
		util.CommonError
		ConfigID string `json:"config_id"`
		QrCode   string `json:"qr_code"`
	}
)

// AddContactWay 配置客户联系「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) AddContactWay(req *AddContactWayRequest) (*AddContactWayResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(AddContactWayURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &AddContactWayResponse{}
	if err = util.DecodeWithError(response, result, "AddContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}

type (
	// GetContactWayRequest 获取企业已配置的「联系我」方式请求
	GetContactWayRequest struct {
		ConfigID string `json:"config_id"`
	}
	// GetContactWayResponse 获取企业已配置的「联系我」方式响应
	GetContactWayResponse struct {
		util.CommonError
		ContactWay ContactWay `json:"contact_way"`
	}
	// ContactWay 「联系我」配置
	ContactWay struct {
		ConfigID      string              `json:"config_id"`
		Type          int                 `json:"type"`
		Scene         int                 `json:"scene"`
		Style         int                 `json:"style"`
		Remark        string              `json:"remark"`
		SkipVerify    bool                `json:"skip_verify"`
		State         string              `json:"state"`
		QrCode        string              `json:"qr_code"`
		User          []string            `json:"user"`
		Party         []int               `json:"party"`
		IsTemp        bool                `json:"is_temp"`
		ExpiresIn     int                 `json:"expires_in"`
		ChatExpiresIn int                 `json:"chat_expires_in"`
		UnionID       string              `json:"unionid"`
		Conclusions   ConclusionsResponse `json:"conclusions"`
	}
)

// GetContactWay 获取企业已配置的「联系我」方式
// see https://developer.work.weixin.qq.com/document/path/92228
func (r *Client) GetContactWay(req *GetContactWayRequest) (*GetContactWayResponse, error) {
	var (
		accessToken string
		err         error
	)
	if accessToken, err = r.GetAccessToken(); err != nil {
		return nil, err
	}
	var response []byte
	if response, err = util.PostJSON(fmt.Sprintf(GetContactWayURL, accessToken), req); err != nil {
		return nil, err
	}
	result := &GetContactWayResponse{}
	if err = util.DecodeWithError(response, result, "GetContactWay"); err != nil {
		return nil, err
	}
	return result, nil
}
