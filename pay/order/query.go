package order

import (
	"encoding/xml"

	"github.com/silenceper/wechat/v2/pay/notify"
	"github.com/silenceper/wechat/v2/util"
)

var queryGateway = "https://api.mch.weixin.qq.com/pay/orderquery"

// 传入的参数
type QueryParams struct {
	OutTradeNo    string // 商户订单号
	SignType      string // 签名类型
	TransactionId string // 微信订单号
}

// queryRequest 接口请求参数
type queryRequest struct {
	AppID         string `xml:"appid"`               // 公众账号ID
	MchID         string `xml:"mch_id"`              // 商户号
	NonceStr      string `xml:"nonce_str"`           // 随机字符串
	Sign          string `xml:"sign"`                // 签名
	SignType      string `xml:"sign_type,omitempty"` // 签名类型
	TransactionId string `xml:"transaction_id"`      // 微信订单号
	OutTradeNo    string `xml:"out_trade_no"`        // 商户订单号
}

// 查询订单
func (o *Order) QueryOrder(p *QueryParams) (paidResult notify.PaidResult, err error) {
	nonceStr := util.RandomStr(32)
	// 签名类型
	if p.SignType == "" {
		p.SignType = "MD5"
	}

	param := make(map[string]string)
	param["appid"] = o.AppID
	param["mch_id"] = o.MchID
	param["nonce_str"] = nonceStr
	param["out_trade_no"] = p.OutTradeNo
	param["sign_type"] = p.SignType
	param["transaction_id"] = p.TransactionId

	sign, err := util.ParamSign(param, o.Key)
	if err != nil {
		return
	}
	request := queryRequest{
		AppID:         o.AppID,
		MchID:         o.MchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		OutTradeNo:    p.OutTradeNo,
		TransactionId: p.TransactionId,
		SignType:      p.SignType,
	}

	rawRet, err := util.PostXML(queryGateway, request)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &paidResult)
	if err != nil {
		return
	}

	return
}
