package pay

import (
	"crypto/md5"	
	"strings"
	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

var payGateway := "https://api.mch.weixin.qq.com/pay/unifiedorder"

// Pay struct extends context
type Pay struct {
	*context.Context
}

// PayParams was NEEDED when request unifiedorder
type PayParams struct {
	TotalFee    string
	CreateIP    string
	Body        string
	OutTradeNo  string
}

type PayResult struct {
	Success     bool
	PrePayID    string
}

//PayRequest
type payRequest struct {
	AppID           string `xml:"appid"`
	MchID           string `xml:"mch_id"`
	NotifyUrl       string `xml:"notify_url"` //通知地址
	DeviceInfo      string `xml:"device_info,omitempty"`
	NonceStr        string `xml:"nonce_str"`
	Sign            string `xml:"sign"`
	SignType        string `xml:"sign_type,omitempty"`
	Body            string `xml:"body"`
	Detail          string `xml:"detail,omitempty"`
	Attach          string `xml:"attach,omitempty"` //附加数据
	OutTradeNo      string `xml:"out_trade_no"` //商户订单号
	FeeType         string `xml:"fee_type,omitempty"` //标价币种
	TotalFee        string `xml:"total_fee"` //标价金额
	SpbillCreateIp  string `xml:"spbill_create_ip"` //终端IP
	TimeStart       string `xml:"time_start,omitempty"`  //交易起始时间
	TimeExpire      string `xml:"time_expire,omitempty"`  //交易结束时间
	GoodsTag        string `xml:"goods_tag,omitempty"`  //订单优惠标记
	TradeType       string `xml:"trade_type"` //交易类型
	ProductId       string `xml:"product_id,omitempty"`  //商品ID
	LimitPay        string `xml:"limit_pay,omitempty"` //
	OpenID          string `xml:"openid,omitempty"` //用户标识
	SceneInfo       string `xml:"scene_info,omitempty"` //场景信息	
}

type payResponse struct {
	
}

// NewPay return an instance of Pay package
func NewPay(ctx *context.Context) *Pay {
	pay := Pay{Context: ctx}
	return &pay
}

// PrePayId will request wechat merchant api and request for a pre payment order id
func (pcf *Pay) PrePayId(p *PayParams) payResult *PayResult {
	nonceStr := util.RandomStr(32)
	pType = "JSAPI"
	template := "appid=%s&body=%s&mch_id=%s&nonce_str=%s&notify_url=%s&out_trade_no=%s&spbill_create_ip=%s&total_fee=%s&trade_type"
	stringA := fmt.Sprintf(template, pcf.AppID, p.Body, pcf.MchID, nonceStr, pcf.NotifyUrl, p.OutTradeNo, p.CreateIP, p.TotalFee, pType)
	signature := md5.Sum(stringA + pcf.PayKey)
	sign := strings.ToUpper(signature)
	request := payRequest{
		AppID: pcf.AppID,
		MchID: pcf.MchID,
		NotifyUrl: pcf.NotifyUrl,
		NonceStr: util.RandomStr(32),
		Sign: sign,
		Body: p.Body,
		OutTradeNo: p.OutTradeNo,
		TotalFee: p.TotalFee,
		SpbillCreateIp: params.CreateIP,
		OpenID: params.OpenID,
	}
	ret, err := util.PostXML(payGateway, request)
	if err != nil {

	}
	fmt.Println(string(ret))
}