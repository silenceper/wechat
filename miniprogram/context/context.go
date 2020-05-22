package context

import (
	"sync"

	"github.com/silenceper/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config

	//accessTokenLock 读写锁 同一个AppID一个
	accessTokenLock *sync.RWMutex

	//accessTokenFunc 自定义获取 access token 的方法
	accessTokenFunc GetAccessTokenFunc
}
