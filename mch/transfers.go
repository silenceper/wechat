package mch

import "github.com/silenceper/wechat/context"

//https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
var transfersGateway = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"

type Transfers struct {
	*context.Context
}

type TransfersParams struct {
}

type TransfersConfig struct {
}
