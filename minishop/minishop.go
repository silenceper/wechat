package minishop

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/minishop/order"
)

// MiniShop 微信小商店相关
type MiniShop struct {
	ctx *context.Context
}

// NewMiniShop 实例化小商店api
func NewMiniShop(cfg *config.Config) *MiniShop {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.AppID, cfg.AppSecret, credential.CacheKeyMiniProgramPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &MiniShop{ctx}
}

// GetOrder 获取订单
func (ms *MiniShop) GetOrder() *order.Order {
	return order.NewOrder(ms.ctx)
}
