package miniprogram

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/util"
)

const (
	templateSendURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send"
)


//Message 发送的模板消息内容
type TemplateMessage struct {
	ToUser     			string              `json:"touser"`          // 必须, 接受者OpenID
	TemplateID 			string              `json:"template_id"`     // 必须, 模版ID
	Page       			string              `json:"page,omitempty"`   // 可选, 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	FormID       		string             	`json:"form_id,omitempty"`   // 必须, 表单提交场景下，为 submit 事件带上的 formId；支付场景下，为本次支付的 prepay_id
	EmphasisKeyword     string     			`json:"emphasis_keyword,omitempty"` // 可选, 模板需要放大的关键词，不填则默认无放大
	Data       map[string]*DataItem 		`json:"data"`            // 必须, 模板数据
}

//DataItem 模版内某个 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type resTemplateSend struct {
	util.CommonError

	MsgID int64 `json:"msgid"`
}

//Send 发送模板消息
func (tpl *MiniProgram) SendTemplate(msg *TemplateMessage) (msgID int64, err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return 0,err
	}
	uri := fmt.Sprintf("%s?access_token=%s", templateSendURL, accessToken)
	response, err := util.PostJSON(uri, msg)

	var result resTemplateSend
	err = json.Unmarshal(response, &result)
	if err != nil {
		return 0,err
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return 0,err
	}
	msgID = result.MsgID
	return msgID,nil
}
