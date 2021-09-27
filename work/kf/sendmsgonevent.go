package kf

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// 发送事件响应消息
	sendMsgOnEventAddr = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg_on_event?access_token=%s"
)

// SendMsgOnEventSchema 发送事件响应消息
type SendMsgOnEventSchema struct {
	util.CommonError
	MsgID string `json:"msgid"` // 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节, 字符串取值范围(正则表达式)：[0-9a-zA-Z_-]*
}

// SendMsgOnEvent 发送事件响应消息
//「进入会话事件」响应消息：
// 如果满足通过API下发欢迎语条件（条件为：1. 企业没有在管理端配置了原生欢迎语；2. 用户在过去48小时里未收过欢迎语，且未向该用户发过消息），则用户进入会话事件会额外返回一个welcome_code，开发者以此为凭据调用接口（填到该接口code参数），即可向客户发送客服欢迎语。
// 为了保证用户体验以及避免滥用，开发者仅可在收到相关事件后20秒内调用，且只可调用一次。
func (r *Client) SendMsgOnEvent(options interface{}) (info SendMsgOnEventSchema, err error) {
	var (
		accessToken string
		data        []byte
	)
	accessToken, err = r.ctx.GetAccessToken()
	if err != nil {
		return
	}
	data, err = util.PostJSON(fmt.Sprintf(sendMsgOnEventAddr, accessToken), options)
	if err != nil {
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
