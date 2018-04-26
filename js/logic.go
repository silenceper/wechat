package js

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/swxctx/wechat/util"
)

// GetJsSign 获取js请求所需的签名
func (js *Js) GetJsSign(url string) (sign string, err error) {
	jsSign := &JsSign{}
	config, err := js.GetConfig(url)
	if err != nil {
		return
	}
	jsSign.Appid = config.AppID
	jsSign.Noncestr = config.NonceStr
	jsSign.Signature = config.Signature
	jsSign.Timestamp = config.Timestamp

	//url 取出# 之后的部分
	urlArr := strings.Split(url, "#")
	jsSign.Url = urlArr[0]
	sign, err = marshalJsSign(jsSign)
	return
}

/*
	1. 从缓存获取
	2. 缓存则重新请求微信服务器获取
*/

// GetConfig 获取jssdk需要的签名信息
func (js *Js) GetConfig(url string) (config *Config, err error) {
	var (
		ticket string
	)
	config = new(Config)
	ticket, err = js.GetTicket()
	if err != nil {
		return
	}

	nonce := util.RandomStr(16)
	timestamp := util.GetCurrTs()
	str := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s", ticket, nonce, timestamp, url)
	sigStr := util.Signature(str)

	config.AppID = js.AppID
	config.NonceStr = nonce
	config.Timestamp = timestamp
	config.Signature = sigStr
	return
}

/*
	{
		"errcode":0,
		"errmsg":"ok",
		"ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
		"expires_in":7200
	}
*/

// GetTicket 获取jsapi_ticket
func (js *Js) GetTicket() (ticket string, err error) {
	js.GetJsAPITicketLock().Lock()
	defer js.GetJsAPITicketLock().Unlock()

	// 从redis中获取
	jsAPITicketKey := fmt.Sprintf("jsapi_ticket_%s", js.AppID)
	val := js.Cache.Get(jsAPITicketKey)
	if val != nil {
		ticket = val.(string)
		return
	}

	var ticketInfo resTicket
	ticketInfo, err = js.GetTicketFromServer()
	if err != nil {
		return
	}
	ticket = ticketInfo.Ticket
	return
}

// GetTicketFromServer 从微信服务器获取ticket
func (js *Js) GetTicketFromServer() (ticket resTicket, err error) {
	var (
		accessToken string
		resp        []byte
	)
	// 获取accessToken用于api请求
	accessToken, err = js.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf(getTicketURL, accessToken)
	resp, err = util.HTTPGet(url)
	err = json.Unmarshal(resp, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("getTicket Error : errcode=%d , errmsg=%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}

	jsAPITicketKey := fmt.Sprintf("jsapi_ticket_%s", js.AppID)
	expire := ticket.ExpiresIn
	if expire < 1 {
		expire = 3600
	} else {
		// 缓冲
		expire = expire - 100
	}
	err = js.Cache.Set(jsAPITicketKey, ticket.Ticket, time.Duration(expire)*time.Second)
	return
}

func marshalJsSign(data *JsSign) (string, error) {
	resource, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(resource), nil
}
