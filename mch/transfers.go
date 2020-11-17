package mch

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
	"sort"
	"strconv"
)

//https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
//https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=4_3
var transfersGateway = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"

type Transfers struct {
	*context.Context
}

type TransfersParams struct {
	//MchAppid       string `xml:"mch_appid"`
	//Mchid          string `xml:"mchid"`
	DeviceInfo string `xml:"device_info"`
	//NonceStr       string `xml:"nonce_str"`
	//Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	OpenId         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name"`
	Amount         string `xml:"amount"`
	Desc           string `xml:"desc"`
	SpbillCreateIp string `xml:"spbill_create_ip"`
	RootCa         string //ca证书
}

type transfersRequest struct {
	MchAppid       string `xml:"mch_appid"`
	Mchid          string `xml:"mchid"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	OpenId         string `xml:"openid"`
	CheckName      string `xml:"check_name"`
	ReUserName     string `xml:"re_user_name,omitempty"`
	Amount         string `xml:"amount"`
	Desc           string `xml:"desc"`
	SpbillCreateIp string `xml:"spbill_create_ip,omitempty"`
}

type TransferResponse struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	MchAppid       string `xml:"mch_appid"`
	MchId          string `xml:"mchid"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str"`
	ResultCode     string `xml:"result_code"`
	ErrCode        string `xml:"err_code,omitempty"`
	ErrCodeDes     string `xml:"err_code_des,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	PaymentNo      string `xml:"payment_no,omitempty"`
	PaymentTime    string `xml:"payment_time,omitempty"`
}

type TransfersConfig struct {
}

func makeParam(source interface{}, bizKey string) (returnStr string) {
	switch v := source.(type) {
	case map[string]string:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v[k])
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			if k == "sign" {
				continue
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var buf bytes.Buffer
		for _, k := range keys {
			if v[k] == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			switch vv := v[k].(type) {
			case string:
				buf.WriteString(vv)
			case int:
				buf.WriteString(strconv.FormatInt(int64(vv), 10))
			default:
				panic("params type not supported")
			}
		}
		buf.WriteString(bizKey)
		returnStr = buf.String()
	}
	return
}

func (m *Mch) Transfers(p *TransfersParams) (resp TransferResponse, err error) {
	nonceStr := util.RandomStr(32)
	param := make(map[string]interface{})
	param["mch_appid"] = m.AppID
	param["mchid"] = m.PayMchID
	param["device_info"] = p.DeviceInfo
	param["nonce_str"] = nonceStr
	param["partner_trade_no"] = p.PartnerTradeNo
	param["openid"] = p.OpenId
	param["check_name"] = p.CheckName
	param["re_user_name"] = p.ReUserName
	param["amount"] = p.Amount
	param["desc"] = p.Desc
	param["spbill_create_ip"] = p.SpbillCreateIp

	bizKey := "&key=" + m.PayKey
	str := makeParam(param, bizKey)
	sign := util.MD5Sum(str)
	request := transfersRequest{
		MchAppid:       m.AppID,
		Mchid:          m.PayMchID,
		DeviceInfo:     p.DeviceInfo,
		NonceStr:       nonceStr,
		Sign:           sign,
		PartnerTradeNo: p.PartnerTradeNo,
		OpenId:         p.OpenId,
		CheckName:      p.CheckName,
		ReUserName:     p.ReUserName,
		Amount:         p.Amount,
		Desc:           p.Desc,
		SpbillCreateIp: p.SpbillCreateIp,
	}
	rawRet, err := util.PostXMLWithTLS(transfersGateway, request, p.RootCa, m.PayMchID)
	if err != nil {
		return
	}

	err = xml.Unmarshal(rawRet, &resp)
	if err != nil {
		return
	}
	if resp.ReturnCode == "SUCCESS" {
		if resp.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = fmt.Errorf("refund error, errcode=%s,errmsg=%s", resp.ErrCode, resp.ErrCodeDes)
		return
	}
	err = fmt.Errorf("[msg : xmlUnmarshalError] [rawReturn : %s] [params : %s] [sign : %s]",
		string(rawRet), str, sign)
	return
}
