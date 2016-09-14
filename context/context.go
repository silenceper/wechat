package context

import (
	"net/http"
	"sync"

	"github.com/silenceper/wechat/cache"
)

//Context struct
type Context struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string

	Cache cache.Cache

	Writer  http.ResponseWriter
	Request *http.Request

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//jsapiTicket 读写锁 同一个AppID一个
	jsApiTicketLock *sync.RWMutex
}

// Query returns the keyed url query value if it exists
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//SetJsApiTicket 设置jsApiTicket的lock
func (ctx *Context) SetJsApiTicketLock(lock *sync.RWMutex) {
	ctx.jsApiTicketLock = lock
}

func (ctx *Context) GetJsApiTicketLock() *sync.RWMutex {
	return ctx.jsApiTicketLock
}
