package robot

import (
	"github.com/silenceper/wechat/v2/work/context"
)

// Client 群聊机器人接口实例
type Client struct {
	*context.Context
}

// NewClient 初始化实例
func NewClient(ctx *context.Context) *Client {
	return &Client{
		ctx,
	}
}
