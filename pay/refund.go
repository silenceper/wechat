package pay

import (
	"encoding/xml"
	"fmt"

	"gitee.com/zhimiao/wechat-sdk/util"
)

var refundGateway = "https://api.mch.weixin.qq.com/secapi/pay/refund"

// RefundParams 调用参数
type RefundParams struct {
	TransactionID string
	OutRefundNo   string
	TotalFee      string
	RefundFee     string
	RefundDesc    string
	// RootCa        string //ca证书
	P12 []byte // 微信加密证书
}

// refundRequest 接口请求参数
type refundRequest struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type,omitempty"`
	TransactionID string `xml:"transaction_id"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      string `xml:"total_fee"`
	RefundFee     string `xml:"refund_fee"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	// NotifyUrl     string `xml:"notify_url,omitempty"`
}

// RefundResponse 接口返回
type RefundResponse struct {
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
func (pcf *Pay) Refund(p *RefundParams) (rsp RefundResponse, err error) {
	nonceStr := util.RandomStr(32)
	param := make(map[string]interface{})
	param["appid"] = pcf.AppID
	param["mch_id"] = pcf.PayMchID
	param["nonce_str"] = nonceStr
	param["out_refund_no"] = p.OutRefundNo
	param["refund_desc"] = p.RefundDesc
	param["refund_fee"] = p.RefundFee
	param["total_fee"] = p.TotalFee
	param["sign_type"] = "MD5"
	param["transaction_id"] = p.TransactionID

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)
	request := refundRequest{
		AppID:         pcf.AppID,
		MchID:         pcf.PayMchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		SignType:      "MD5",
		TransactionID: p.TransactionID,
		OutRefundNo:   p.OutRefundNo,
		TotalFee:      p.TotalFee,
		RefundFee:     p.RefundFee,
		RefundDesc:    p.RefundDesc,
	}
	rawRet, err := util.PostXMLWithTLS(refundGateway, request, p.P12, pcf.PayMchID)
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
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return
}
