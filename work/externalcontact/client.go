package externalcontact

import (
	"github.com/silenceper/wechat/v2/work/context"
)

// Client 实例
type Client struct {
	*context.Context
}

// NewClient
func NewClient(ctx *context.Context) (client *Client, err error) {

	client = &Client{
		ctx,
	}

	return client, nil
}
