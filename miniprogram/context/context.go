package context

import (
	"github.com/linyken/wechat/v2/credential"
	"github.com/linyken/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenContextHandle
}
