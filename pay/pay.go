package pay

import (
	"github.com/silenceper/wechat/pay/config"
	"github.com/silenceper/wechat/pay/notify"
	"github.com/silenceper/wechat/pay/order"
)

//Pay 微信支付相关API
type Pay struct {
	cfg *config.Config
}

//NewPay 实例化微信支付相关API
func NewPay(cfg *config.Config) *Pay {
	return &Pay{cfg}
}

// GetOrder  下单
func (pay *Pay) GetOrder() *order.Order {
	return order.NewOrder(pay.cfg)
}

// GetNotify  通知
func (pay *Pay) GetNotify() *notify.Notify {
	return notify.NewNotify(pay.cfg)
}
