// +build !linux

//Package msgaudit for unsupport platform
package msgaudit

import (
	"fmt"

	"github.com/silenceper/wechat/v2/work/config"
)

// Client 会话存档
type Client struct {
}

// NewClient new
func NewClient(cfg *config.Config) (*Client, error) {
	return nil, fmt.Errorf("会话存档功能目前只支持Linux平台运行")
}
