package business

import "github.com/linyken/wechat/v2/miniprogram/context"

// Business 业务
type Business struct {
	*context.Context
}

// NewBusiness init
func NewBusiness(ctx *context.Context) *Business {
	return &Business{ctx}
}
