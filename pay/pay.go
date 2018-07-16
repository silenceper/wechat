package pay

import (
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

var payGateway = "https://api.mch.weixin.qq.com/pay/unifiedorder"

// Pay struct extends context
type Pay struct {
	*context.Context
}

// Params was NEEDED when request unifiedorder
// 传入的参数，用于生成 prepay_id 的必需参数
type Params struct {
	TotalFee   string
	CreateIP   string
	Body       string
	OutTradeNo string
	OpenID     string
	Attach     string
}

// Config 是传出用于 jsdk 用的参数
type Config struct {
	AppId     string
	Timestamp int64
	NonceStr  string
	PrePayID  string
	SignType  string
	Sign      string
}

// payResult 是 unifie order 接口的返回
type payResult struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppID      string `xml:"appid,omitempty"`
	MchID      string `xml:"mch_id,omitempty"`
	NonceStr   string `xml:"nonce_str,omitempty"`
	Sign       string `xml:"sign,omitempty"`
	ResultCode string `xml:"result_code,omitempty"`
	TradeType  string `xml:"trade_type,omitempty"`
	PrePayID   string `xml:"prepay_id,omitempty"`
	CodeURL    string `xml:"code_url,omitempty"`
	ErrCode    string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
}

//payRequest 接口请求参数
type payRequest struct {
	AppID          string `xml:"appid"`
	MchID          string `xml:"mch_id"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	SignType       string `xml:"sign_type,omitempty"`
	Body           string `xml:"body"`
	Detail         string `xml:"detail,omitempty"`
	Attach         string `xml:"attach,omitempty"`      //附加数据
	OutTradeNo     string `xml:"out_trade_no"`          //商户订单号
	FeeType        string `xml:"fee_type,omitempty"`    //标价币种
	TotalFee       string `xml:"total_fee"`             //标价金额
	SpbillCreateIP string `xml:"spbill_create_ip"`      //终端IP
	TimeStart      string `xml:"time_start,omitempty"`  //交易起始时间
	TimeExpire     string `xml:"time_expire,omitempty"` //交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty"`   //订单优惠标记
	NotifyURL      string `xml:"notify_url"`            //通知地址
	TradeType      string `xml:"trade_type"`            //交易类型
	ProductID      string `xml:"product_id,omitempty"`  //商品ID
	LimitPay       string `xml:"limit_pay,omitempty"`   //
	OpenID         string `xml:"openid,omitempty"`      //用户标识
	SceneInfo      string `xml:"scene_info,omitempty"`  //场景信息
}

// NewPay return an instance of Pay package
func NewPay(ctx *context.Context) *Pay {
	pay := Pay{Context: ctx}
	return &pay
}

// PrePayID will request wechat merchant api and request for a pre payment order id
func (pcf *Pay) PrePayID(p *Params) (prePayID string, err error) {
	nonceStr := util.RandomStr(32)
	tradeType := "JSAPI"
	template := "appid=%s&attach=%s&body=%s&mch_id=%s&nonce_str=%s&notify_url=%s&openid=%s&out_trade_no=%s&spbill_create_ip=%s&total_fee=%s&trade_type=%s&key=%s"
	str := fmt.Sprintf(template, pcf.AppID, p.Attach, p.Body, pcf.PayMchID, nonceStr, pcf.PayNotifyURL, p.OpenID, p.OutTradeNo, p.CreateIP, p.TotalFee, tradeType, pcf.PayKey)
	sign := util.MD5Sum(str)
	request := payRequest{
		AppID:          pcf.AppID,
		MchID:          pcf.PayMchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		Body:           p.Body,
		OutTradeNo:     p.OutTradeNo,
		TotalFee:       p.TotalFee,
		SpbillCreateIP: p.CreateIP,
		NotifyURL:      pcf.PayNotifyURL,
		TradeType:      tradeType,
		OpenID:         p.OpenID,
		Attach:         p.Attach,
	}
	rawRet, err := util.PostXML(payGateway, request)
	if err != nil {
		return "", errors.New(err.Error() + " parameters : " + str)
	}
	payRet := payResult{}
	err = xml.Unmarshal(rawRet, &payRet)
	if err != nil {
		return "", errors.New(err.Error())
	}
	if payRet.ReturnCode == "SUCCESS" {
		//pay success
		if payRet.ResultCode == "SUCCESS" {
			return payRet.PrePayID, nil
		}
		return "", errors.New(payRet.ErrCode + payRet.ErrCodeDes)
	}
	return "", errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [params : " + str + "] [sign : " + sign + "]")
}

// 获取js配置
func (pcf *Pay) GetPayJsConf(prePayID string) (config *Config) {
	config = new(Config)
	nonceStr := util.RandomStr(32)
	timeStamp := util.GetCurrTs()
	config.NonceStr = nonceStr
	config.Timestamp = timeStamp
	config.PrePayID = prePayID
	config.SignType = "MD5"
	config.AppId = pcf.AppID
	template := "appId=%s&nonceStr=%s&package=%s&signType=%s&timeStamp=%d&key=%s"
	str := fmt.Sprintf(template, pcf.AppID, nonceStr, "prepay_id="+prePayID, "MD5", timeStamp, pcf.PayKey)
	sign := util.MD5Sum(str)
	config.Sign = sign
	return
}
