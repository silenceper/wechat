package order

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"hash"
	"strconv"
	"strings"
	"time"

	"github.com/silenceper/wechat/v2/pay/config"
	"github.com/silenceper/wechat/v2/util"
)

//https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
var payGateway = "https://api.mch.weixin.qq.com/pay/unifiedorder"

// Order struct extends context
type Order struct {
	*config.Config
}

// NewOrder return an instance of order package
func NewOrder(cfg *config.Config) *Order {
	order := Order{cfg}
	return &order
}

// Params was NEEDED when request unifiedorder
// 传入的参数，用于生成 prepay_id 的必需参数
type Params struct {
	TotalFee   string
	CreateIP   string
	Body       string
	OutTradeNo string
	OpenID     string
	TradeType  string
	SignType   string
	Detail     string
	Attach     string
	GoodsTag   string
	NotifyURL  string
}

// Config 是传出用于 js sdk 用的参数
type Config struct {
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	PrePayID  string `json:"prePayId"`
	SignType  string `json:"signType"`
	Package   string `json:"package"`
	PaySign   string `json:"paySign"`
}

// PreOrder 是 unifie order 接口的返回
type PreOrder struct {
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

// payRequest 接口请求参数
type payRequest struct {
	AppID          string `xml:"appid"`
	MchID          string `xml:"mch_id"`
	DeviceInfo     string `xml:"device_info,omitempty"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	SignType       string `xml:"sign_type,omitempty"`
	Body           string `xml:"body"`
	Detail         string `xml:"detail,omitempty"`
	Attach         string `xml:"attach,omitempty"`      // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`          // 商户订单号
	FeeType        string `xml:"fee_type,omitempty"`    // 标价币种
	TotalFee       string `xml:"total_fee"`             // 标价金额
	SpbillCreateIP string `xml:"spbill_create_ip"`      // 终端IP
	TimeStart      string `xml:"time_start,omitempty"`  // 交易起始时间
	TimeExpire     string `xml:"time_expire,omitempty"` // 交易结束时间
	GoodsTag       string `xml:"goods_tag,omitempty"`   // 订单优惠标记
	NotifyURL      string `xml:"notify_url"`            // 通知地址
	TradeType      string `xml:"trade_type"`            // 交易类型
	ProductID      string `xml:"product_id,omitempty"`  // 商品ID
	LimitPay       string `xml:"limit_pay,omitempty"`   //
	OpenID         string `xml:"openid,omitempty"`      // 用户标识
	SceneInfo      string `xml:"scene_info,omitempty"`  // 场景信息
}

// BridgeConfig get js bridge config
func (o *Order) BridgeConfig(p *Params) (cfg Config, err error) {
	var (
		buffer    strings.Builder
		h         hash.Hash
		timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	)
	order, err := o.PrePayOrder(p)
	if err != nil {
		return
	}
	buffer.WriteString("appId=")
	buffer.WriteString(order.AppID)
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(order.NonceStr)
	buffer.WriteString("&package=")
	buffer.WriteString("prepay_id=" + order.PrePayID)
	buffer.WriteString("&signType=")
	buffer.WriteString(p.SignType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timestamp)
	buffer.WriteString("&key=")
	buffer.WriteString(o.Key)
	if p.SignType == "MD5" {
		h = md5.New()
	} else {
		h = hmac.New(sha256.New, []byte(o.Key))
	}
	h.Write([]byte(buffer.String()))
	// 签名
	cfg.PaySign = strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	cfg.NonceStr = order.NonceStr
	cfg.Timestamp = timestamp
	cfg.PrePayID = order.PrePayID
	cfg.SignType = p.SignType
	cfg.Package = "prepay_id=" + order.PrePayID
	return
}

// PrePayOrder return data for invoke wechat payment
func (o *Order) PrePayOrder(p *Params) (payOrder PreOrder, err error) {
	nonceStr := util.RandomStr(32)
	notifyURL := o.NotifyURL
	// 签名类型
	if p.SignType == "" {
		p.SignType = "MD5"
	}
	// 通知地址
	if p.NotifyURL != "" {
		notifyURL = p.NotifyURL
	}
	param := make(map[string]interface{})
	param["appid"] = o.AppID
	param["body"] = p.Body
	param["mch_id"] = o.MchID
	param["nonce_str"] = nonceStr
	param["out_trade_no"] = p.OutTradeNo
	param["spbill_create_ip"] = p.CreateIP
	param["total_fee"] = p.TotalFee
	param["trade_type"] = p.TradeType
	param["openid"] = p.OpenID
	param["sign_type"] = p.SignType
	param["detail"] = p.Detail
	param["attach"] = p.Attach
	param["goods_tag"] = p.GoodsTag
	param["notify_url"] = notifyURL

	bizKey := "&key=" + o.Key
	str := util.OrderParam(param, bizKey)
	sign := util.MD5Sum(str)
	request := payRequest{
		AppID:          o.AppID,
		MchID:          o.MchID,
		NonceStr:       nonceStr,
		Sign:           sign,
		Body:           p.Body,
		OutTradeNo:     p.OutTradeNo,
		TotalFee:       p.TotalFee,
		SpbillCreateIP: p.CreateIP,
		NotifyURL:      notifyURL,
		TradeType:      p.TradeType,
		OpenID:         p.OpenID,
		SignType:       p.SignType,
		Detail:         p.Detail,
		Attach:         p.Attach,
		GoodsTag:       p.GoodsTag,
	}
	rawRet, err := util.PostXML(payGateway, request)
	if err != nil {
		return
	}
	err = xml.Unmarshal(rawRet, &payOrder)
	if err != nil {
		return
	}
	if payOrder.ReturnCode == "SUCCESS" {
		// pay success
		if payOrder.ResultCode == "SUCCESS" {
			err = nil
			return
		}
		err = errors.New(payOrder.ErrCode + payOrder.ErrCodeDes)
		return
	}
	err = errors.New("[msg : xmlUnmarshalError] [rawReturn : " + string(rawRet) + "] [params : " + str + "] [sign : " + sign + "]")
	return
}

// PrePayID will request wechat merchant api and request for a pre payment order id
func (o *Order) PrePayID(p *Params) (prePayID string, err error) {
	order, err := o.PrePayOrder(p)
	if err != nil {
		return
	}
	if order.PrePayID == "" {
		err = errors.New("empty prepayid")
	}
	prePayID = order.PrePayID
	return
}
