package js

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/machao520/wechat/util"
	"github.com/machao520/wechat/wechat"
)

const getTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

// Js struct
type Js struct {
	*wechat.Wechat
}

// Config 返回给用户jssdk配置信息
type Config struct {
	AppID     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Signature string `json:"signature"`
}

// resTicket 请求jsapi_tikcet返回结果
type resTicket struct {
	util.CommonError

	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
}

//NewJs init
func NewJs( wc *wechat.Wechat) *Js {
	js := new(Js)
	js.Wechat = wc
	return js
}

//GetConfig 获取jssdk需要的配置参数
//uri 为当前网页地址
func (js *Js) GetConfig(uri string) (config *Config, err error) {
	config = new(Config)
	var ticketStr string
	ticketStr, err = js.GetTicket()
	if err != nil {
		return
	}

	nonceStr := util.RandomStr(16)
	timestamp := util.GetCurrTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticketStr, nonceStr, timestamp, uri)
	sigStr := util.Signature(str)

	config.AppID = js.AppID
	config.NonceStr = nonceStr
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}

//GetTicket 获取jsapi_ticket
func (js *Js) GetTicket() (ticketStr string, err error) {
	js.Wechat.GetJsAPITicketLock().Lock()
	defer js.GetJsAPITicketLock().Unlock()

	//先从cache中取
	jsAPITicketCacheKey := fmt.Sprintf("jsapi_ticket_%s", js.AppID)
	val := js.Cache.Get(jsAPITicketCacheKey)
	if val != nil {
		ticketStr = val.(string)
		return
	}
	var ticket resTicket
	ticket, err = js.getTicketFromServer()
	if err != nil {
		return
	}
	ticketStr = ticket.Ticket
	return
}

//getTicketFromServer 强制从服务器中获取ticket
func (js *Js) getTicketFromServer() (ticket resTicket, err error) {
	var accessToken string
	accessToken, err = js.GetAccessToken()
	if err != nil {
		return
	}

	var response []byte
	url := fmt.Sprintf(getTicketURL, accessToken)
	response, err = util.HTTPGet(url)
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}

	jsAPITicketCacheKey := fmt.Sprintf("jsapi_ticket_%s", js.AppID)
	expires := ticket.ExpiresIn - 1500
	err = js.Cache.Set(jsAPITicketCacheKey, ticket.Ticket, time.Duration(expires)*time.Second)
	return
}
