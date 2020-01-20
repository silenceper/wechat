package context

import (
	"sync"

	"github.com/silenceper/wechat/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//jsAPITicket 读写锁 同一个AppID一个
	jsAPITicketLock *sync.RWMutex

	//accessTokenFunc 自定义获取 access token 的方法
	accessTokenFunc GetAccessTokenFunc
}

// SetJsAPITicketLock 设置jsAPITicket的lock
func (ctx *Context) SetJsAPITicketLock(lock *sync.RWMutex) {
	ctx.jsAPITicketLock = lock
}

// GetJsAPITicketLock 获取jsAPITicket 的lock
func (ctx *Context) GetJsAPITicketLock() *sync.RWMutex {
	return ctx.jsAPITicketLock
}
