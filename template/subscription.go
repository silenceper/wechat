package template

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/util"
)

const (
	//发送订阅消息
	//https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
	subscriptionMessageSendURL = "https://api.weixin.qq.com/cgi-bin/message/subscribe/send"
)

// SubscriptionMessage 订阅消息
type SubscriptionMessage struct {
	ToUser     string               `json:"touser"`      //必选，接收者（用户）的 openid
	TemplateID string               `json:"template_id"` //必选，所需下发的订阅模板id
	Page       string               `json:"page"`        //可选，点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data       map[string]*DataItem `json:"data"`        // 必须, 模板内容
}

// SendSubscriptionMessage 发送订阅消息
func (tpl *Template) SendSubscriptionMessage(msg *SubscriptionMessage) (err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", subscriptionMessageSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)

	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("subscription msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
