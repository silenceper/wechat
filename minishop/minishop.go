package minishop

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/minishop/config"
	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/minishop/order"
	"github.com/silenceper/wechat/v2/minishop/service"
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
	ms := &MiniShop{ctx}
	return ms
}

// SetServeice 设置服务商
func (ms *MiniShop) SetServeice(serviceID, specificationID string) {
	ms.ctx.Config.ServiceID = serviceID
	ms.ctx.Config.SpecificationID = specificationID
}

// GetOrder 获取订单
func (ms *MiniShop) GetOrder() *order.Order {
	return order.NewOrder(ms.ctx)
}

// GetService 获取服务商
func (ms *MiniShop) GetService() *service.Service {
	return service.NewService(ms.ctx)
}
