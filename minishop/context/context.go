package context

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	ServiceID         string
	SpecificationID   string
	Config            *config.Config
	AccessTokenHandle credential.AccessTokenHandle
}
