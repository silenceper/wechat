package refund

import (
	"encoding/xml"
	"fmt"

	"github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/util"
)

var refundGateway = "https://api.mch.weixin.qq.com/secapi/pay/refund"

// Refund struct extends context
type Refund struct {
	*config.Config
}

// NewRefund return an instance of refund package
func NewRefund(cfg *config.Config) *Refund {
	refund := Refund{cfg}
	return &refund
}

// Params 调用参数
type Params struct {
	TransactionID string // 微信支付订单号
	OutRefundNo   string // 商户订单号
	TotalFee      string // 订单金额
	SignType      string // 签名类型
	RefundFee     string // 退款金额
	RefundDesc    string // 退款原因
	RootCa        string // ca证书
	NotifyURL     string // 退款结果通知url
}

// request 接口请求参数
type request struct {
	AppID         string `xml:"appid"`                 // 公众账号ID
	MchID         string `xml:"mch_id"`                // 商户号
	NonceStr      string `xml:"nonce_str"`             // 随机字符串
	Sign          string `xml:"sign"`                  // 签名
	SignType      string `xml:"sign_type,omitempty"`   // 签名类型
	TransactionID string `xml:"transaction_id"`        // 微信支付订单号
	OutRefundNo   string `xml:"out_refund_no"`         // 商户订单号
	TotalFee      string `xml:"total_fee"`             // 订单金额
	RefundFee     string `xml:"refund_fee"`            // 退款金额
	RefundFeeType string `xml:"refund_fee_type"`       // 退款货币种类
	RefundDesc    string `xml:"refund_desc,omitempty"` // 退款原因
	NotifyURL     string `xml:"notify_url,omitempty"`  // 退款结果通知url
}

// Response 接口返回
type Response struct {
	ReturnCode          string `xml:"return_code"`
	ReturnMsg           string `xml:"return_msg"`
	AppID               string `xml:"appid,omitempty"`
	MchID               string `xml:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	TransactionID       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundID            string `xml:"refund_id,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             string `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`
}

// Refund 退款申请
func (refund *Refund) Refund(p *Params) (rsp Response, err error) {
	nonceStr := util.RandomStr(32)
	param := make(map[string]string)

	// 签名类型
	if p.SignType == "" {
		p.SignType = util.SignTypeMD5
	}

	param["mch_id"] = refund.MchID
	param["nonce_str"] = nonceStr
	param["out_refund_no"] = p.OutRefundNo
	param["refund_desc"] = p.RefundDesc
	param["refund_fee"] = p.RefundFee
	param["total_fee"] = p.TotalFee
	param["sign_type"] = p.SignType
	param["transaction_id"] = p.TransactionID

	if p.OutTradeNo != "" {
		param["out_trade_no"] = p.OutTradeNo
	}
	if p.NotifyURL != "" {
		param["notify_url"] = p.NotifyURL
	}

	sign, err := util.ParamSign(param, refund.Key)
	if err != nil {
		return
	}

	req := request{
		AppID:         refund.AppID,
		MchID:         refund.MchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		SignType:      p.SignType,
		TransactionID: p.TransactionID,
		OutRefundNo:   p.OutRefundNo,
		OutTradeNo:    p.OutTradeNo,
		TotalFee:      p.TotalFee,
		RefundFee:     p.RefundFee,
		RefundDesc:    p.RefundDesc,
		NotifyURL:     p.NotifyURL,
	}
	rawRet, err := util.PostXMLWithTLS(refundGateway, req, p.RootCa, refund.MchID)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &rsp)
	if err != nil {
		return
	}
	if rsp.ReturnCode == "SUCCESS" {
		if rsp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("refund error, errcode=%s,errmsg=%s", rsp.ErrCode, rsp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [sign : %s]", string(rawRet), sign)
	return
}
