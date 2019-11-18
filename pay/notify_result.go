package pay

import (
	"fmt"
	"github.com/silenceper/wechat/util"
	"sort"
)

// Base 公用参数
type Base struct {
	AppID    string `xml:"appid"`
	MchID    string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"`
}

// NotifyResult 下单回调
type NotifyResult struct {
	Base
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	ResultCode    string `xml:"result_code"`
	OpenID        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int    `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int    `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

// NotifyResp 消息通知返回
type NotifyResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

// VerifySign 验签
func (pcf *Pay) VerifySign(notifyRes NotifyResult) bool {
	// 封装map 请求过来的 map
	resMap := make(map[string]interface{})
	resMap["appid"] = notifyRes.AppID
	resMap["bank_type"] = notifyRes.BankType
	resMap["cash_fee"] = notifyRes.CashFee
	resMap["fee_type"] = notifyRes.FeeType
	resMap["is_subscribe"] = notifyRes.IsSubscribe
	resMap["mch_id"] = notifyRes.MchID
	resMap["nonce_str"] = notifyRes.NonceStr
	resMap["openid"] = notifyRes.OpenID
	resMap["out_trade_no"] = notifyRes.OutTradeNo
	resMap["result_code"] = notifyRes.ResultCode
	resMap["return_code"] = notifyRes.ReturnCode
	resMap["time_end"] = notifyRes.TimeEnd
	resMap["total_fee"] = notifyRes.TotalFee
	resMap["trade_type"] = notifyRes.TradeType
	resMap["transaction_id"] = notifyRes.TransactionID
	// 支付key
	sortedKeys := make([]string, 0, len(resMap))
	for k := range resMap {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	// STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {
		value := fmt.Sprintf("%v", resMap[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	// STEP3, 在键值对的最后加上key=API_KEY
	signStrings = signStrings + "key=" + pcf.PayKey
	// STEP4, 进行MD5签名并且将所有字符转为大写.
	sign := util.MD5Sum(signStrings)
	if sign != notifyRes.Sign {
		return false
	}
	return true
}
