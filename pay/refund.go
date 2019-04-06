package pay

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"github.com/akikistyle/wechat/util"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"net/http"
)

var refundGateway = "https://api.mch.weixin.qq.com/secapi/pay/refund"

//Refund Parameter
type RefundParams struct {
	TransactionId string
	OutRefundNo   string
	TotalFee      string
	RefundFee     string
	RefundDesc    string
	RootCa        string //ca证书
}

//Refund request
type refundRequest struct {
	AppID         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	SignType      string `xml:"sign_type,omitempty"`
	TransactionId string `xml:"transaction_id"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      string `xml:"total_fee"`
	RefundFee     string `xml:"refund_fee"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	//NotifyUrl     string `xml:"notify_url,omitempty"`
}

//Refund Response
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
	TransactionId       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty"`
	RefundFee           string `xml:"refund_fee,omitempty"`
	SettlementRefundFee string `xml:"settlement_refund_fee,omitempty"`
	TotalFee            string `xml:"total_fee,omitempty"`
	SettlementTotalFee  string `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             string `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`
}

//退款申请
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
	param["transaction_id"] = p.TransactionId

	bizKey := "&key=" + pcf.PayKey
	str := orderParam(param, bizKey)
	sign := util.MD5Sum(str)
	request := refundRequest{
		AppID:         pcf.AppID,
		MchID:         pcf.PayMchID,
		NonceStr:      nonceStr,
		Sign:          sign,
		SignType:      "MD5",
		TransactionId: p.TransactionId,
		OutRefundNo:   p.OutRefundNo,
		TotalFee:      p.TotalFee,
		RefundFee:     p.RefundFee,
		RefundDesc:    p.RefundDesc,
	}
	rawRet, err := postXMLWithTLS(refundGateway, request, p.RootCa, pcf.PayMchID)
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

//http TLS
func httpWithTLS(rootCa, key string) (*http.Client, error) {
	var client *http.Client
	certData, err := ioutil.ReadFile(rootCa)
	if err != nil {
		return nil, fmt.Errorf("unable to find cert path=%s, error=%v", rootCa, err)
	}
	cert := pkcs12ToPem(certData, key)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	tr := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	client = &http.Client{Transport: tr}
	return client, nil
}

//将Pkcs12转成Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {
	blocks, err := pkcs12.ToPEM(p12, password)
	defer func() {
		if x := recover(); x != nil {
			log.Print(x)
		}
	}()
	if err != nil {
		panic(err)
	}
	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}

//Post XML with TLS
func postXMLWithTLS(uri string, obj interface{}, ca, key string) ([]byte, error) {
	xmlData, err := xml.Marshal(obj)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(xmlData)
	client, err := httpWithTLS(ca, key)
	if err != nil {
		return nil, err
	}
	response, err := client.Post(uri, "application/xml;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http code error : uri=%v , statusCode=%v", uri, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
