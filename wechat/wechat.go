package wechat

import (
	"github.com/machao520/wechat/cache"
	"sync"
)

type Wechat struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	PayMchID       string //支付 - 商户 ID
	PayNotifyURL   string //支付 - 接受微信支付结果通知的接口地址
	PayKey         string //支付 - 商户后台设置的支付 key
	Cache          cache.Cache


	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.RWMutex

	//accessTokenFunc 自定义获取 access token 的方法
	accessTokenFunc GetAccessTokenFunc
}


// SetJsAPITicketLock 设置jsAPITicket的lock
func (ctx *Wechat) SetJsAPITicketLock(lock *sync.RWMutex) {
	ctx.jsAPITicketLock = lock
}

// GetJsAPITicketLock 获取jsAPITicket 的lock
func (ctx *Wechat) GetJsAPITicketLock() *sync.RWMutex {
	return ctx.jsAPITicketLock
}