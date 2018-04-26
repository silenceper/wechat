package js

import (
	"github.com/swxctx/wechat/context"
	"github.com/swxctx/wechat/util"
)

const getTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

// Js struct
type Js struct {
	*context.Context
}

// Config 返回给用户jssdk配置信息
type Config struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

// JsSign JsSign
type JsSign struct {
	Appid     string `json:"appid"`
	Noncestr  string `json:"noncestr"`
	Timestamp int64  `json:"timestamp"`
	Url       string `json:"url"`
	Signature string `json:"signature"`
}

// resTicket 请求jsapi_tikcet返回结果
type resTicket struct {
	util.CommonError

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

// NewJs init
func NewJs(context *context.Context) *Js {
	js := new(Js)
	js.Context = context
	return js
}
