package message

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/miniprogram/security"
	"github.com/silenceper/wechat/v2/util"
)

// ConfirmReceiveMethod 确认收货方式
type ConfirmReceiveMethod int8

const (
	// EventTypeTradeManageRemindAccessAPI 提醒接入发货信息管理服务API
	// 小程序完成账期授权时/小程序产生第一笔交易时/已产生交易但从未发货的小程序，每天一次
	EventTypeTradeManageRemindAccessAPI EventType = "trade_manage_remind_access_api"
	// EventTypeTradeManageRemindShipping 提醒需要上传发货信息
	//	曾经发过货的小程序，订单超过48小时未发货时
	EventTypeTradeManageRemindShipping EventType = "trade_manage_remind_shipping"
	// EventTypeTradeManageOrderSettlement 订单将要结算或已经结算
	// 订单完成发货时/订单结算时
	EventTypeTradeManageOrderSettlement EventType = "trade_manage_order_settlement"
	// EventTypeWxaMediaCheck 媒体内容安全异步审查结果通知
	EventTypeWxaMediaCheck EventType = "wxa_media_check"
	// ConfirmReceiveMethodAuto 自动确认收货
	ConfirmReceiveMethodAuto ConfirmReceiveMethod = 1
	// ConfirmReceiveMethodManual 手动确认收货
	ConfirmReceiveMethodManual ConfirmReceiveMethod = 2
)

// PushReceiver 接收消息推送
// 暂仅支付Aes加密方式
type PushReceiver struct {
	*context.Context
	Token  string // Token(令牌)
	AesKey string // 用于消息解密(消息加密密钥)
}

// NewPushReceiver 实例化
func NewPushReceiver(ctx *context.Context, token, aesKey string) *PushReceiver {
	return &PushReceiver{
		Context: ctx,
		Token:   token,
		AesKey:  aesKey,
	}
}

// GetMsg 获取接收到的消息(如果是加密的返回解密数据)
func (receiver *PushReceiver) GetMsg(r *http.Request) ([]byte, error) {
	// 读取参数
	signature := r.FormValue("signature")
	timestamp := r.FormValue("timestamp")
	nonce := r.FormValue("nonce")
	encryptType := r.FormValue("encrypt_type")
	// 验证签名
	tmpArr := []string{
		receiver.Token,
		timestamp,
		nonce,
	}

	sort.Strings(tmpArr)
	tmpSignature := util.Signature(tmpArr...)
	if tmpSignature != signature {
		return nil, errors.New("signature error")
	}

	if encryptType == "aes" {
		// 解密
		var reqData DataReceived
		if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
			return nil, err
		}
		if _, rawMsgBytes, err := util.DecryptMsg(receiver.AppID, reqData.Encrypt, receiver.AesKey); err != nil {
			return nil, err
		} else {
			return rawMsgBytes, nil
		}
	}
	// 不加密
	return io.ReadAll(r.Body)
}

// GetMsgData 获取接收到的消息(解密数据)
func (receiver *PushReceiver) GetMsgData(r *http.Request) (MsgType, EventType, PushData, error) {
	decryptMsg, err := receiver.GetMsg(r)
	if err != nil {
		return "", "", nil, err
	}

	dataMap := make(map[string]interface{})
	if err := json.Unmarshal(decryptMsg, &dataMap); err != nil {
		return "", "", nil, err
	}

	msgType := MsgType(dataMap["MsgType"].(string))
	eventType := EventType(dataMap["Event"].(string))
	if msgType == MsgTypeEvent {
		switch eventType {
		case EventTypeTradeManageRemindAccessAPI:
			// 提醒接入发货信息管理服务API
			var pushData PushDataRemindAccessAPIData
			if err := json.Unmarshal(decryptMsg, &pushData); err != nil {
				return msgType, eventType, nil, err
			}
			return msgType, eventType, &pushData, nil
		case EventTypeTradeManageRemindShipping:
			// 提醒需要上传发货信息
			var pushData PushDataRemindShippingData
			if err := json.Unmarshal(decryptMsg, &pushData); err != nil {
				return msgType, eventType, nil, err
			}
			return msgType, eventType, &pushData, nil
		case EventTypeTradeManageOrderSettlement:
			// 订单将要结算或已经结算
			var pushData PushDataOrderSettlementData
			if err := json.Unmarshal(decryptMsg, &pushData); err != nil {
				return msgType, eventType, nil, err
			}
			return msgType, eventType, &pushData, nil
		case EventTypeWxaMediaCheck:
			// 媒体内容安全异步审查结果通知
			var pushData MediaCheckAsyncData
			if err := json.Unmarshal(decryptMsg, &pushData); err != nil {
				return msgType, eventType, &pushData, err
			}
			return msgType, eventType, &pushData, nil
		default:
			// 暂不支持其他事件类型
			return msgType, eventType, string(decryptMsg), nil
		}
	}
	// 暂不支持其他消息类型
	return msgType, eventType, string(decryptMsg), nil
}

// DataReceived 接收到的数据
type DataReceived struct {
	Encrypt string `json:"Encrypt"` // 加密的消息体
}

// PushData 推送的数据(已转对应的结构体)
// after 1.18: interface{ MediaCheckAsyncData | PushDataRemindAccessApiData | PushDataRemindShippingData | PushDataOrderSettlementData}
type PushData interface{}

// CommonPushData 推送数据通用部分
type CommonPushData struct {
	MsgType      MsgType   `json:"MsgType"`      // 消息类型，为固定值 "event"
	Event        EventType `json:"Event"`        // 事件类型
	ToUserName   string    `json:"ToUserName"`   // 小程序的原始 ID
	FromUserName string    `json:"FromUserName"` // 发送方账号（一个 OpenID，此时发送方是系统账号）
	CreateTime   int64     `json:"CreateTime"`   // 消息创建时间 （整型），时间戳
}

// MediaCheckAsyncData 媒体内容安全异步审查结果通知
type MediaCheckAsyncData struct {
	CommonPushData
	Appid   string                `json:"appid"`
	TraceID string                `json:"trace_id"`
	Version int                   `json:"version"`
	Detail  []*MediaCheckDetail   `json:"detail"`
	Errcode int                   `json:"errcode"`
	Errmsg  string                `json:"errmsg"`
	Result  MediaCheckAsyncResult `json:"result"`
}

// MediaCheckDetail 检测结果详情
type MediaCheckDetail struct {
	Strategy string                `json:"strategy"`
	Errcode  int                   `json:"errcode"`
	Suggest  security.CheckSuggest `json:"suggest"`
	Label    int                   `json:"label"`
	Prob     int                   `json:"prob"`
}

// MediaCheckAsyncResult 检测结果
type MediaCheckAsyncResult struct {
	Suggest security.CheckSuggest `json:"suggest"`
	Label   security.CheckLabel   `json:"label"`
}

// PushDataOrderSettlementData 订单将要结算或已经结算通知
type PushDataOrderSettlementData struct {
	CommonPushData
	TransactionID           string               `json:"transaction_id"`            // 支付订单号
	MerchantID              string               `json:"merchant_id"`               // 商户号
	SubMerchantID           string               `json:"sub_merchant_id"`           // 子商户号
	MerchantTradeNo         string               `json:"merchant_trade_no"`         // 商户订单号
	PayTime                 int64                `json:"pay_time"`                  // 支付成功时间，秒级时间戳
	ShippedTime             int64                `json:"shipped_time"`              // 发货时间，秒级时间戳
	EstimatedSettlementTime int64                `json:"estimated_settlement_time"` // 预计结算时间，秒级时间戳。发货时推送才有该字段
	ConfirmReceiveMethod    ConfirmReceiveMethod `json:"confirm_receive_method"`    // 确认收货方式：1. 自动确认收货；2. 手动确认收货。结算时推送才有该字段
	ConfirmReceiveTime      int64                `json:"confirm_receive_time"`      // 确认收货时间，秒级时间戳。结算时推送才有该字段
	SettlementTime          int64                `json:"settlement_time"`           // 订单结算时间，秒级时间戳。结算时推送才有该字段
}

// PushDataRemindShippingData 提醒需要上传发货信息
type PushDataRemindShippingData struct {
	CommonPushData
	TransactionID   string `json:"transaction_id"`    // 微信支付订单号
	MerchantID      string `json:"merchant_id"`       // 商户号
	SubMerchantID   string `json:"sub_merchant_id"`   // 子商户号
	MerchantTradeNo string `json:"merchant_trade_no"` // 商户订单号
	PayTime         int64  `json:"pay_time"`          // 支付成功时间，秒级时间戳
	Msg             string `json:"msg"`               // 消息文本内容
}

// PushDataRemindAccessAPIData 提醒接入发货信息管理服务API
type PushDataRemindAccessAPIData struct {
	CommonPushData
	Msg string `json:"msg"` // 消息文本内容
}
