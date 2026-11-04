package kf

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	customerBatchGetAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget?access_token=%s"
)

// CustomerBatchGetOptions 客户基本信息获取请求参数
type CustomerBatchGetOptions struct {
	ExternalUserIDList []string `json:"external_userid_list"` // external_userid列表
	// NeedEnterSessionContext 是否需要返回客户48小时内最后一次进入会话的上下文信息
	// 0-不返回 1-返回。默认不返回
	NeedEnterSessionContext int `json:"need_enter_session_context"`
}

// CustomerWechatChannelsSchema 进入会话的视频号信息，从视频号进入会话才有值
type CustomerWechatChannelsSchema struct {
	Nickname     string `json:"nickname"`      // 视频号名称，视频号场景值为1、2、3时返回此项
	ShopNickname string `json:"shop_nickname"` // 视频号小店名称，视频号场景值为4、5时返回此项
	Scene        uint32 `json:"scene"`         // 视频号场景值。1：视频号主页，2：视频号直播间商品列表页，3：视频号商品橱窗页，4：视频号小店商品详情页，5：视频号小店订单页
}

// CustomerEnterSessionContextSchema 48小时内最后一次进入会话的上下文信息
// 请求的need_enter_session_context参数设置为1才返回
type CustomerEnterSessionContextSchema struct {
	Scene          string                       `json:"scene"`           // 进入会话的场景值，获取客服账号链接开发者自定义的场景值
	SceneParam     string                       `json:"scene_param"`     // 进入会话的自定义参数，获取客服账号链接返回的url，开发者按规范拼接的scene_param参数
	WechatChannels CustomerWechatChannelsSchema `json:"wechat_channels"` // 进入会话的视频号信息，从视频号进入会话才有值
}

// CustomerSchema 微信客户基本资料
type CustomerSchema struct {
	ExternalUserID string `json:"external_userid"` // 微信客户的external_userid
	NickName       string `json:"nickname"`        // 微信昵称
	Avatar         string `json:"avatar"`          // 微信头像。第三方不可获取
	Gender         int    `json:"gender"`          // 性别
	UnionID        string `json:"unionid"`         // unionid，需要绑定微信开发者帐号才能获取到，查看绑定方法: https://open.work.weixin.qq.com/kf/doc/92512/93143/94769#%E5%A6%82%E4%BD%95%E8%8E%B7%E5%8F%96%E5%BE%AE%E4%BF%A1%E5%AE%A2%E6%88%B7%E7%9A%84unionid
	// EnterSessionContext 48小时内最后一次进入会话的上下文信息
	// 请求的need_enter_session_context参数设置为1才返回
	EnterSessionContext CustomerEnterSessionContextSchema `json:"enter_session_context"`
}

// CustomerBatchGetSchema 获取客户基本信息响应内容
type CustomerBatchGetSchema struct {
	util.CommonError
	CustomerList          []CustomerSchema `json:"customer_list"`           // 微信客户信息列表
	InvalidExternalUserID []string         `json:"invalid_external_userid"` // 无效的微信客户ID
}

// CustomerBatchGet 客户基本信息获取
func (r *Client) CustomerBatchGet(options CustomerBatchGetOptions) (info CustomerBatchGetSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	if accessToken, err = r.ctx.GetAccessToken(); err != nil {
		return
	}
	if data, err = util.PostJSON(fmt.Sprintf(customerBatchGetAddr, accessToken), options); err != nil {
		return
	}
	if err = json.Unmarshal(data, &info); err != nil {
		return
	}
	if info.ErrCode != 0 {
		return info, NewSDKErr(info.ErrCode, info.ErrMsg)
	}
	return info, nil
}
