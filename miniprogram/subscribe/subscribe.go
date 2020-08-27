package subscribe

import (
	"fmt"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	//发送订阅消息
	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	subscribeSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"

	// 获取当前帐号下的个人模板列表
	// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
	getTemplateURL = "https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate"
)

// Subscribe 订阅消息
type Subscribe struct {
	*context.Context
}

// NewSubscribe 实例化
func NewSubscribe(ctx *context.Context) *Subscribe {
	return &Subscribe{Context: ctx}
}

// Message 订阅消息请求参数
type Message struct {
	ToUser           string               `json:"touser"`            //必选，接收者（用户）的 openid
	TemplateID       string               `json:"template_id"`       //必选，所需下发的订阅模板id
	Page             string               `json:"page"`              //可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data             map[string]*DataItem `json:"data"`              //必选, 模板内容
	MiniprogramState string               `json:"miniprogram_state"` //可选，跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string               `json:"lang"`              //入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
}

//TemplateItem template item
type TemplateItem struct {
	PriTmplID string `json:"priTmplId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Example   string `json:"example"`
	Type      int64  `json:"type"`
}

//TemplateList template list
type TemplateList struct {
	util.CommonError
	Data []TemplateItem `json:"data"`
}

// Send 发送订阅消息
func (s *Subscribe) Send(msg *Message) (err error) {
	var accessToken string
	accessToken, err = s.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscribeSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	return util.DecodeWithCommonError(response, "Send")
}

//ListTemplates 获取当前帐号下的个人模板列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.getTemplateList.html
func (s *Subscribe) ListTemplates() (*TemplateList, error) {
	accessToken, err := s.GetAccessToken()
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("%s?access_token=%s", getTemplateURL, accessToken)
	response, err := util.HTTPGet(uri)
	if err != nil {
		return nil, err
	}
	templateList := TemplateList{}
	err = util.DecodeWithError(response, &templateList, "ListTemplates")
	if err != nil {
		return nil, err
	}
	return &templateList, nil
}
