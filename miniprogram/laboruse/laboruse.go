package laboruse

import (
	context2 "context"
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	// sendEmployeeRelationMsgURL 推送用工消息
	// see https://developers.weixin.qq.com/miniprogram/dev/server/API/laboruse/api_sendemployeerelationmsg.html
	sendEmployeeRelationMsgURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/employeerelationmsg/send?access_token=%s"
	// unbindUserB2CAuthInfoURL 解绑用工关系
	// see https://developers.weixin.qq.com/miniprogram/dev/server/API/laboruse/api_unbinduserb2cauthinfo.html
	unbindUserB2CAuthInfoURL = "https://api.weixin.qq.com/wxa/business/unbinduserb2cauthinfo?access_token=%s"
)

// LaborUse 用工关系
type LaborUse struct {
	*context.Context
}

// NewLaborUse 实例化用工关系接口
func NewLaborUse(ctx *context.Context) *LaborUse {
	return &LaborUse{ctx}
}

// DataItem 模版消息中某个关键词的值
type DataItem struct {
	Value string `json:"value"`           // 关键词对应的值
	Color string `json:"color,omitempty"` // 关键词展示颜色，可选
}

// employeeRelationMsgContent 用工消息内容的内层结构
type employeeRelationMsgContent struct {
	Data map[string]*DataItem `json:"data"`
}

// SendEmployeeRelationMsgRequest 推送用工消息请求参数
type SendEmployeeRelationMsgRequest struct {
	TemplateID string               `json:"template_id"` // 必选，消息模版id
	Page       string               `json:"page"`        // 必选，点击模版卡片后的跳转路径
	ToUser     string               `json:"touser"`      // 必选，被推送用户的openid
	Data       map[string]*DataItem `json:"-"`           // 必选，消息内容，由 SDK 按官方要求封装为 data 字符串
}

// sendEmployeeRelationMsgBody 实际请求体，官方要求 data 字段为字符串
type sendEmployeeRelationMsgBody struct {
	TemplateID string `json:"template_id"`
	Page       string `json:"page"`
	ToUser     string `json:"touser"`
	Data       string `json:"data"`
}

// SendEmployeeRelationMsg 推送用工消息，向已绑定用工关系的用户通过小程序商家助手发送模版消息
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/laboruse/api_sendemployeerelationmsg.html
func (labor *LaborUse) SendEmployeeRelationMsg(in *SendEmployeeRelationMsgRequest) error {
	return labor.SendEmployeeRelationMsgContext(context2.Background(), in)
}

// SendEmployeeRelationMsgContext 推送用工消息
func (labor *LaborUse) SendEmployeeRelationMsgContext(ctx context2.Context, in *SendEmployeeRelationMsgRequest) error {
	var accessToken string
	var err error
	if accessToken, err = labor.GetAccessTokenContext(ctx); err != nil {
		return err
	}

	// 官方要求 data 为字符串，内容为 {"data":{...}}，这里由 SDK 完成封装
	content, err := json.Marshal(&employeeRelationMsgContent{Data: in.Data})
	if err != nil {
		return err
	}
	body := &sendEmployeeRelationMsgBody{
		TemplateID: in.TemplateID,
		Page:       in.Page,
		ToUser:     in.ToUser,
		Data:       string(content),
	}

	uri := fmt.Sprintf(sendEmployeeRelationMsgURL, accessToken)
	var response []byte
	if response, err = util.PostJSONContext(ctx, uri, body); err != nil {
		return err
	}

	// 使用通用方法返回错误
	return util.DecodeWithCommonError(response, "laboruse.SendEmployeeRelationMsg")
}

// UnbindUserB2CAuthInfoRequest 解绑用工关系请求参数
type UnbindUserB2CAuthInfoRequest struct {
	OpenIDList []string `json:"openid_list"` // 必选，被解绑用户的openid列表
}

// UnbindUserB2CAuthInfo 解绑用工关系，批量解绑 B2C 模块已绑定的用工关系
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/laboruse/api_unbinduserb2cauthinfo.html
func (labor *LaborUse) UnbindUserB2CAuthInfo(in *UnbindUserB2CAuthInfoRequest) error {
	return labor.UnbindUserB2CAuthInfoContext(context2.Background(), in)
}

// UnbindUserB2CAuthInfoContext 解绑用工关系
func (labor *LaborUse) UnbindUserB2CAuthInfoContext(ctx context2.Context, in *UnbindUserB2CAuthInfoRequest) error {
	var accessToken string
	var err error
	if accessToken, err = labor.GetAccessTokenContext(ctx); err != nil {
		return err
	}

	uri := fmt.Sprintf(unbindUserB2CAuthInfoURL, accessToken)
	var response []byte
	if response, err = util.PostJSONContext(ctx, uri, in); err != nil {
		return err
	}

	// 使用通用方法返回错误
	return util.DecodeWithCommonError(response, "laboruse.UnbindUserB2CAuthInfo")
}
