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
	// 曾经发过货的小程序，订单超过48小时未发货时
	EventTypeTradeManageRemindShipping EventType = "trade_manage_remind_shipping"
	// EventTypeTradeManageOrderSettlement 订单将要结算或已经结算
	// 订单完成发货时/订单结算时
	EventTypeTradeManageOrderSettlement EventType = "trade_manage_order_settlement"
	// EventTypeAddExpressPath 运单轨迹更新事件
	EventTypeAddExpressPath EventType = "add_express_path"
	// EventTypeSecvodUpload 短剧媒资上传完成事件
	EventTypeSecvodUpload EventType = "secvod_upload_event"
	// EventTypeSecvodAudit 短剧媒资审核状态事件
	EventTypeSecvodAudit EventType = "secvod_audit_event"
	// EventTypeWxaMediaCheck 媒体内容安全异步审查结果通知
	EventTypeWxaMediaCheck EventType = "wxa_media_check"
	// EventTypeXpayGoodsDeliverNotify 道具发货推送事件
	EventTypeXpayGoodsDeliverNotify EventType = "xpay_goods_deliver_notify"
	// EventTypeXpayCoinPayNotify 代币支付推送事件
	EventTypeXpayCoinPayNotify EventType = "xpay_coin_pay_notify"
	// ConfirmReceiveMethodAuto 自动确认收货
	ConfirmReceiveMethodAuto ConfirmReceiveMethod = 1
	// ConfirmReceiveMethodManual 手动确认收货
	ConfirmReceiveMethodManual ConfirmReceiveMethod = 2
)

// PushReceiver 接收消息推送
// 暂仅支付Aes加密方式
type PushReceiver struct {
	*context.Context
}

// NewPushReceiver 实例化
func NewPushReceiver(ctx *context.Context) *PushReceiver {
	return &PushReceiver{
		Context: ctx,
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
		_, rawMsgBytes, err := util.DecryptMsg(receiver.AppID, reqData.Encrypt, receiver.EncodingAESKey)
		return rawMsgBytes, err
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
			var pushData PushDataRemindAccessAPI
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeTradeManageRemindShipping:
			// 提醒需要上传发货信息
			var pushData PushDataRemindShipping
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeTradeManageOrderSettlement:
			// 订单将要结算或已经结算
			var pushData PushDataOrderSettlement
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeWxaMediaCheck:
			// 媒体内容安全异步审查结果通知
			var pushData MediaCheckAsyncData
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeAddExpressPath:
			// 运单轨迹更新
			var pushData PushDataAddExpressPath
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeSecvodUpload:
			// 短剧媒资上传完成
			var pushData PushDataSecVodUpload
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeSecvodAudit:
			// 短剧媒资审核状态
			var pushData PushDataSecVodAudit
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeXpayGoodsDeliverNotify:
			// 道具发货推送事件
			var pushData PushDataXpayGoodsDeliverNotify
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		case EventTypeXpayCoinPayNotify:
			// 代币支付推送事件
			var pushData PushDataXpayCoinPayNotify
			err := json.Unmarshal(decryptMsg, &pushData)
			return msgType, eventType, &pushData, err
		}
		// 暂不支持其他事件类型
		return msgType, eventType, decryptMsg, nil
	}
	// 暂不支持其他消息类型
	return msgType, eventType, decryptMsg, nil
}

// DataReceived 接收到的数据
type DataReceived struct {
	Encrypt string `json:"Encrypt"` // 加密的消息体
}

// PushData 推送的数据(已转对应的结构体)
// MediaCheckAsyncData | PushDataRemindAccessAPI | PushDataRemindShipping | PushDataOrderSettlement | []byte
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

// PushDataOrderSettlement 订单将要结算或已经结算通知
type PushDataOrderSettlement struct {
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

// PushDataRemindShipping 提醒需要上传发货信息
type PushDataRemindShipping struct {
	CommonPushData
	TransactionID   string `json:"transaction_id"`    // 微信支付订单号
	MerchantID      string `json:"merchant_id"`       // 商户号
	SubMerchantID   string `json:"sub_merchant_id"`   // 子商户号
	MerchantTradeNo string `json:"merchant_trade_no"` // 商户订单号
	PayTime         int64  `json:"pay_time"`          // 支付成功时间，秒级时间戳
	Msg             string `json:"msg"`               // 消息文本内容
}

// PushDataRemindAccessAPI 提醒接入发货信息管理服务API信息
type PushDataRemindAccessAPI struct {
	CommonPushData
	Msg string `json:"msg"` // 消息文本内容
}

// PushDataAddExpressPath 运单轨迹更新信息
type PushDataAddExpressPath struct {
	CommonPushData
	DeliveryID string                          `json:"DeliveryID"` // 快递公司ID
	WayBillID  string                          `json:"WaybillId"`  // 运单ID
	OrderID    string                          `json:"OrderId"`    // 订单ID
	Version    int                             `json:"Version"`    // 轨迹版本号（整型）
	Count      int                             `json:"Count"`      // 轨迹节点数（整型）
	Actions    []*PushDataAddExpressPathAction `json:"Actions"`    // 轨迹节点列表
}

// PushDataAddExpressPathAction 轨迹节点
type PushDataAddExpressPathAction struct {
	ActionTime int64  `json:"ActionTime"` // 轨迹节点 Unix 时间戳
	ActionType int    `json:"ActionType"` // 轨迹节点类型
	ActionMsg  string `json:"ActionMsg"`  // 轨迹节点详情
}

// PushDataSecVodUpload 短剧媒资上传完成
type PushDataSecVodUpload struct {
	CommonPushData
	UploadEvent SecVodUploadEvent `json:"upload_event"` // 上传完成事件
}

// SecVodUploadEvent 短剧媒资上传完成事件
type SecVodUploadEvent struct {
	MediaID       string `json:"media_id"`       // 媒资id
	SourceContext string `json:"source_context"` // 透传上传接口中开发者设置的值。
	Errcode       int    `json:"errcode"`        // 错误码，上传失败时该值非
	Errmsg        string `json:"errmsg"`         // 错误提示
}

// PushDataSecVodAudit 短剧媒资审核状态
type PushDataSecVodAudit struct {
	CommonPushData
	AuditEvent SecVodAuditEvent `json:"audit_event"` // 审核状态事件
}

// SecVodAuditEvent 短剧媒资审核状态事件
type SecVodAuditEvent struct {
	DramaID       string           `json:"drama_id"`       // 剧目id
	SourceContext string           `json:"source_context"` // 透传上传接口中开发者设置的值
	AuditDetail   DramaAuditDetail `json:"audit_detail"`   // 剧目审核结果，单独每一集的审核结果可以根据drama_id查询剧集详情得到
}

// DramaAuditDetail 剧目审核结果
type DramaAuditDetail struct {
	Status     int   `json:"status"`      // 审核状态，0为无效值；1为审核中；2为最终失败；3为审核通过；4为驳回重填
	CreateTime int64 `json:"create_time"` // 提审时间戳
	AuditTime  int64 `json:"audit_time"`  // 审核时间戳
}

// PushDataXpayGoodsDeliverNotify 道具发货推送
type PushDataXpayGoodsDeliverNotify struct {
	CommonPushData
	OpenID        string        `json:"OpenId"`        // 用户openid
	OutTradeNo    string        `json:"OutTradeNo"`    // 业务订单号
	Env           int           `json:"Env"`           //，环境配置 0：现网环境（也叫正式环境）1：沙箱环境
	WeChatPayInfo WeChatPayInfo `json:"WeChatPayInfo"` // 微信支付信息 非微信支付渠道可能没有
	GoodsInfo     struct {
		GoodsId string `json:"goods_id"` // 道具id
	} `json:"GoodsInfo"` // 道具参数信息
}

// WeChatPayInfo 微信支付信息
type WeChatPayInfo struct {
	MchOrderNo    string `json:"MchOrderNo"`    // 微信支付商户单号
	TransactionID string `json:"TransactionId"` // 交易单号（微信支付订单号）
	PaidTime      int64  `json:"PaidTime"`      // 用户支付时间，Linux秒级时间戳
}

// GoodsInfo 道具参数信息
type GoodsInfo struct {
	ProductID   string `json:"ProductId"`   // 道具ID
	Quantity    int    `json:"Quantity"`    // 数量
	OrigPrice   int64  `json:"OrigPrice"`   // 物品原始价格 （单位：分）
	ActualPrice int64  `json:"ActualPrice"` // 物品实际支付价格（单位：分）
	Attach      string `json:"Attach"`      // 透传信息
}

// PushDataXpayCoinPayNotify 代币支付推送
type PushDataXpayCoinPayNotify struct {
	CommonPushData
	OpenID        string        `json:"OpenId"`        // 用户openid
	OutTradeNo    string        `json:"OutTradeNo"`    // 业务订单号
	Env           int           `json:"Env"`           //，环境配置 0：现网环境（也叫正式环境）1：沙箱环境
	WeChatPayInfo WeChatPayInfo `json:"WeChatPayInfo"` // 微信支付信息 非微信支付渠道可能没有
	CoinInfo      CoinInfo      `json:"CoinInfo"`      // 代币参数信息
}

// CoinInfo 代币参数信息
type CoinInfo struct {
	Quantity    int    `json:"Quantity"`    // 数量
	OrigPrice   int64  `json:"OrigPrice"`   // 物品原始价格 （单位：分）
	ActualPrice int64  `json:"ActualPrice"` // 物品实际支付价格（单位：分）
	Attach      string `json:"Attach"`      // 透传信息
}
