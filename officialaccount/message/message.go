package message

import (
"encoding/xml"

"github.com/silenceper/wechat/v2/officialaccount/device"
)

// MsgType 基本消息类型
type MsgType string

// EventType 事件类型
type EventType string

// InfoType 第三方平台授权事件类型
type InfoType string

const (
// MsgTypeText 表示文本消息
MsgTypeText MsgType = "text"
// MsgTypeImage 表示图片消息
MsgTypeImage = "image"
// MsgTypeVoice 表示语音消息
MsgTypeVoice = "voice"
// MsgTypeVideo 表示视频消息
MsgTypeVideo = "video"
// MsgTypeMiniprogrampage 表示小程序卡片消息
MsgTypeMiniprogrampage = "miniprogrampage"
// MsgTypeShortVideo 表示短视频消息[限接收]
MsgTypeShortVideo = "shortvideo"
// MsgTypeLocation 表示坐标消息[限接收]
MsgTypeLocation = "location"
// MsgTypeLink 表示链接消息[限接收]
MsgTypeLink = "link"
// MsgTypeMusic 表示音乐消息[限回复]
MsgTypeMusic = "music"
// MsgTypeNews 表示图文消息[限回复]
MsgTypeNews = "news"
// MsgTypeTransfer 表示消息消息转发到客服
MsgTypeTransfer = "transfer_customer_service"
// MsgTypeEvent 表示事件推送消息
MsgTypeEvent = "event"
)

const (
// EventSubscribe 订阅
EventSubscribe EventType = "subscribe"
// EventUnsubscribe 取消订阅
EventUnsubscribe = "unsubscribe"
// EventScan 用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
EventScan = "SCAN"
// EventLocation 上报地理位置事件
EventLocation = "LOCATION"
// EventClick 点击菜单拉取消息时的事件推送
EventClick = "CLICK"
// EventView 点击菜单跳转链接时的事件推送
EventView = "VIEW"
// EventScancodePush 扫码推事件的事件推送
EventScancodePush = "scancode_push"
// EventScancodeWaitmsg 扫码推事件且弹出“消息接收中”提示框的事件推送
EventScancodeWaitmsg = "scancode_waitmsg"
// EventPicSysphoto 弹出系统拍照发图的事件推送
EventPicSysphoto = "pic_sysphoto"
// EventPicPhotoOrAlbum 弹出拍照或者相册发图的事件推送
EventPicPhotoOrAlbum = "pic_photo_or_album"
// EventPicWeixin 弹出微信相册发图器的事件推送
EventPicWeixin = "pic_weixin"
// EventLocationSelect 弹出地理位置选择器的事件推送
EventLocationSelect = "location_select"
// EventTemplateSendJobFinish 发送模板消息推送通知
EventTemplateSendJobFinish = "TEMPLATESENDJOBFINISH"
// EventMassSendJobFinish 群发消息推送通知
EventMassSendJobFinish = "MASSSENDJOBFINISH"
// EventWxaMediaCheck 异步校验图片/音频是否含有违法违规内容推送事件
EventWxaMediaCheck = "wxa_media_check"
)

const (
// 微信开放平台需要用到

// InfoTypeVerifyTicket 返回ticket
InfoTypeVerifyTicket InfoType = "component_verify_ticket"
// InfoTypeAuthorized 授权
InfoTypeAuthorized = "authorized"
// InfoTypeUnauthorized 取消授权
InfoTypeUnauthorized = "unauthorized"
// InfoTypeUpdateAuthorized 更新授权
InfoTypeUpdateAuthorized = "updateauthorized"
// notify_third_fasteregister 注册审核事件推送
InfoTypeNotifyThirdFasteregister = "notify_third_fasteregister"
)

// MixMessage 存放所有微信发送过来的消息和事件
type MixMessage struct {
CommonToken

// 基本消息
MsgID         int64   `xml:"MsgId"` // 其他消息推送过来是MsgId
TemplateMsgID int64   `xml:"MsgID"` // 模板消息推送成功的消息是MsgID
Content       string  `xml:"Content"`
Recognition   string  `xml:"Recognition"`
PicURL        string  `xml:"PicUrl"`
MediaID       string  `xml:"MediaId"`
Format        string  `xml:"Format"`
ThumbMediaID  string  `xml:"ThumbMediaId"`
LocationX     float64 `xml:"Location_X"`
LocationY     float64 `xml:"Location_Y"`
Scale         float64 `xml:"Scale"`
Label         string  `xml:"Label"`
Title         string  `xml:"Title"`
Description   string  `xml:"Description"`
URL           string  `xml:"Url"`

// 事件相关
Event       EventType `xml:"Event"`
EventKey    string    `xml:"EventKey"`
Ticket      string    `xml:"Ticket"`
Latitude    string    `xml:"Latitude"`
Longitude   string    `xml:"Longitude"`
Precision   string    `xml:"Precision"`
MenuID      string    `xml:"MenuId"`
Status      string    `xml:"Status"`
SessionFrom string    `xml:"SessionFrom"`
TotalCount  int64     `xml:"TotalCount"`
FilterCount int64     `xml:"FilterCount"`
SentCount   int64     `xml:"SentCount"`
ErrorCount  int64     `xml:"ErrorCount"`

ScanCodeInfo struct {
ScanType   string `xml:"ScanType"`
ScanResult string `xml:"ScanResult"`
} `xml:"ScanCodeInfo"`

SendPicsInfo struct {
Count   int32      `xml:"Count"`
PicList []EventPic `xml:"PicList>item"`
} `xml:"SendPicsInfo"`

SendLocationInfo struct {
LocationX float64 `xml:"Location_X"`
LocationY float64 `xml:"Location_Y"`
Scale     float64 `xml:"Scale"`
Label     string  `xml:"Label"`
Poiname   string  `xml:"Poiname"`
}

// 第三方平台相关
InfoType                     InfoType `xml:"InfoType"`
AppID                        string   `xml:"AppId"`
ComponentVerifyTicket        string   `xml:"ComponentVerifyTicket"`
AuthorizerAppid              string   `xml:"AuthorizerAppid"`
AuthorizationCode            string   `xml:"AuthorizationCode"`
AuthorizationCodeExpiredTime int64    `xml:"AuthorizationCodeExpiredTime"`
PreAuthCode                  string   `xml:"PreAuthCode"`
AuthCode string `xml:"auth_code"`
Info struct{
Name string `xml:"name"`
Code string `xml:"code"`
CodeType int `xml:"code_type"`
LegalPersonaWechat string `xml:"legal_persona_wechat"`
LegalPersonaName string `xml:"legal_persona_name"`
ComponentPhone string `xml:"component_phone"`
} `xml:"info"`

// 卡券相关
CardID              string `xml:"CardId"`
RefuseReason        string `xml:"RefuseReason"`
IsGiveByFriend      int32  `xml:"IsGiveByFriend"`
FriendUserName      string `xml:"FriendUserName"`
UserCardCode        string `xml:"UserCardCode"`
OldUserCardCode     string `xml:"OldUserCardCode"`
OuterStr            string `xml:"OuterStr"`
IsRestoreMemberCard int32  `xml:"IsRestoreMemberCard"`
UnionID             string `xml:"UnionId"`

// 内容审核相关
IsRisky       bool   `xml:"isrisky"`
ExtraInfoJSON string `xml:"extra_info_json"`
TraceID       string `xml:"trace_id"`
StatusCode    int    `xml:"status_code"`

// 设备相关
device.MsgDevice
}

// EventPic 发图事件推送
type EventPic struct {
PicMd5Sum string `xml:"PicMd5Sum"`
}

// EncryptedXMLMsg 安全模式下的消息体
type EncryptedXMLMsg struct {
XMLName      struct{} `xml:"xml" json:"-"`
ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
EncryptedMsg string   `xml:"Encrypt"    json:"Encrypt"`
}

// ResponseEncryptedXMLMsg 需要返回的消息体
type ResponseEncryptedXMLMsg struct {
XMLName      struct{} `xml:"xml" json:"-"`
EncryptedMsg string   `xml:"Encrypt"      json:"Encrypt"`
MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

// CDATA  使用该类型,在序列化为 xml 文本时文本会被解析器忽略
type CDATA string

// MarshalXML 实现自己的序列化方法
func (c CDATA) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
return e.EncodeElement(struct {
string `xml:",cdata"`
}{string(c)}, start)
}

// CommonToken 消息中通用的结构
type CommonToken struct {
XMLName      xml.Name `xml:"xml"`
ToUserName   CDATA    `xml:"ToUserName"`
FromUserName CDATA    `xml:"FromUserName"`
CreateTime   int64    `xml:"CreateTime"`
MsgType      MsgType  `xml:"MsgType"`
}

// SetToUserName set ToUserName
func (msg *CommonToken) SetToUserName(toUserName CDATA) {
msg.ToUserName = toUserName
}

// SetFromUserName set FromUserName
func (msg *CommonToken) SetFromUserName(fromUserName CDATA) {
msg.FromUserName = fromUserName
}

// SetCreateTime set createTime
func (msg *CommonToken) SetCreateTime(createTime int64) {
msg.CreateTime = createTime
}

// SetMsgType set MsgType
func (msg *CommonToken) SetMsgType(msgType MsgType) {
msg.MsgType = msgType
}

// GetOpenID get the FromUserName value
func (msg *CommonToken) GetOpenID() string {
return string(msg.FromUserName)
}
