package context

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/minishop/config"
)

// Context struct
type Context struct {
	Config            *config.Config
	AccessTokenHandle credential.AccessTokenHandle
}
