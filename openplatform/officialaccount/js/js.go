package js

import (
	context2 "context"
	"fmt"

	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	officialJs "github.com/silenceper/wechat/v2/officialaccount/js"
	"github.com/silenceper/wechat/v2/util"
)

// Js wx jssdk
type Js struct {
	*context.Context
	credential.JsTicketHandle
	credential.JsTicketContextHandle
}

// NewJs init
func NewJs(context *context.Context, appID string) *Js {
	js := new(Js)
	js.Context = context
	jsTicketHandle := credential.NewDefaultJsTicket(appID, credential.CacheKeyOfficialAccountPrefix, context.Cache)
	js.SetJsTicketHandle(jsTicketHandle)
	return js
}

// SetJsTicketHandle 自定义js ticket取值方式
func (js *Js) SetJsTicketHandle(ticketHandle credential.JsTicketHandle) {
	js.JsTicketHandle = ticketHandle
}

// GetConfig 第三方平台 - 获取jssdk需要的配置参数
// uri 为当前网页地址
func (js *Js) GetConfig(uri, appid string) (config *officialJs.Config, err error) {
	return js.GetConfigContext(context2.Background(), uri, appid)
}

// GetConfigContext 新方法，允许传入上下文，避免协程泄漏
func (js *Js) GetConfigContext(ctx context2.Context, uri, appid string) (config *officialJs.Config, err error) {
	config = new(officialJs.Config)
	var accessToken string
	accessToken, err = js.GetAccessTokenContext(ctx)
	if err != nil {
		return
	}
	var ticketStr string
	ticketStr, err = js.GetTicketContext(ctx, accessToken)
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTS()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.AppID = appid
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}
