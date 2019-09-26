package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/util"
)

const (
	customerSendMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//CustomerMessageText 文本类型客服消息
type CustomerMessageText struct {
	ToUser  string    `json:"touser"` //接受者OpenID
	Msgtype MsgType   `json:"msgtype"`
	Text    MediaText `json:"text"`
}

//MediaText 文本消息的文字
type MediaText struct {
	Content string `json:"content"`
}

//NewCustomerTextMessage 文本消息结构体构造方法
func NewCustomerTextMessage(toUser, text string) *CustomerMessageText {
	return &CustomerMessageText{
		ToUser:  toUser,
		Msgtype: MsgTypeText,
		Text: MediaText{
			text,
		},
	}
}

//SendWithToken 发送文本类型客服消息
func (msg *CustomerMessageText) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageImg 图片类型客服消息
type CustomerMessageImg struct {
	ToUser  string        `json:"touser"` //接受者OpenID
	Msgtype string        `json:"msgtype"`
	Image   MediaResource `json:"image"`
}

//MediaResource 图片消息的资源id
type MediaResource struct {
	MediaID string `json:"media_id"`
}

//NewCustomerImgMessage 图片消息的构造方法
func NewCustomerImgMessage(toUser, mediaID string) *CustomerMessageImg {
	return &CustomerMessageImg{
		ToUser:  toUser,
		Msgtype: MsgTypeImage,
		Image: MediaResource{
			mediaID,
		},
	}
}

//SendWithToken 发送图片类型客服消息
func (msg *CustomerMessageImg) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageVoice 语音类型客服消息
type CustomerMessageVoice struct {
	ToUser  string        `json:"touser"` //接受者OpenID
	Msgtype string        `json:"msgtype"`
	Voice   MediaResource `json:"voice"`
}

//NewCustomerVoiceMessage 语音消息的构造方法
func NewCustomerVoiceMessage(toUser, mediaID string) *CustomerMessageVoice {
	return &CustomerMessageVoice{
		ToUser:  toUser,
		Msgtype: MsgTypeVoice,
		Voice: MediaResource{
			mediaID,
		},
	}
}

//SendWithToken 语音类型客服消息
func (msg *CustomerMessageVoice) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageVideo 视频类型客服消息 ，
type CustomerMessageVideo struct {
	ToUser  string     `json:"touser"` //接受者OpenID
	Msgtype string     `json:"msgtype"`
	Video   MediaVideo `json:"video"`
}

//MediaVideo 视频消息包含的内容
type MediaVideo struct {
	MediaID      string `json:"media_id"`
	ThumbMediaID string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

//SendWithToken 视频类型客服消息
func (msg *CustomerMessageVideo) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageMusic 音乐类型客服消息
type CustomerMessageMusic struct {
	ToUser  string     `json:"touser"` //接受者OpenID
	Msgtype string     `json:"msgtype"`
	Music   MediaMusic `json:"music"`
}

//MediaMusic 音乐消息包括的内容
type MediaMusic struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Musicurl     string `json:"musicurl"`
	Hqmusicurl   string `json:"hqmusicurl"`
	ThumbMediaID string `json:"thumb_media_id"`
}

//SendWithToken 音乐类型客服消息
func (msg *CustomerMessageMusic) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageNews 图文消息类型客服消息，点击跳转到外链
type CustomerMessageNews struct {
	ToUser  string    `json:"touser"` //接受者OpenID
	Msgtype MsgType   `json:"msgtype"`
	News    MediaNews `json:"articles"`
}

//MediaNews 图文消息的内容
type MediaNews struct {
	Articles []MediaArticles `json:"articles"`
}

//MediaArticles 图文消息的内容的文章列表中的单独一条
type MediaArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Picurl      string `json:"picurl"`
}

//SendWithToken 图文消息类型客服消息，点击跳转到外链
func (msg *CustomerMessageNews) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageInnerNews 图文消息类型客服消息,点击跳转到图文消息页面
type CustomerMessageInnerNews struct {
	ToUser  string        `json:"touser"` //接受者OpenID
	Msgtype string        `json:"msgtype"`
	News    MediaResource `json:"mpnews"`
}

//SendWithToken 图文消息类型客服消息,点击跳转到图文消息页面
func (msg *CustomerMessageInnerNews) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageMsgmenu 菜单类型客服消息
type CustomerMessageMsgmenu struct {
	ToUser  string       `json:"touser"` //接受者OpenID
	Msgtype MsgType      `json:"msgtype"`
	News    MediaMsgmenu `json:"msgmenu"`
}

//MediaMsgmenu 菜单消息的内容
type MediaMsgmenu struct {
	HeadContent string        `json:"head_content"`
	List        []MsgmenuItem `json:"list"`
	TailContent string        `json:"tail_content"`
}

//MsgmenuItem 菜单消息的菜单按钮
type MsgmenuItem struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

//SendWithToken 发送菜单类型客服消息
func (msg *CustomerMessageMsgmenu) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageWxcard 卡券类型客服消息
type CustomerMessageWxcard struct {
	ToUser  string      `json:"touser"` //接受者OpenID
	Msgtype MsgType     `json:"msgtype"`
	Wxcard  MediaWxcard `json:"wxcard"`
}

//MediaWxcard 卡券的id
type MediaWxcard struct {
	CardID string `json:"card_id"`
}

//SendWithToken 发送卡券类型客服消息
func (msg *CustomerMessageWxcard) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//CustomerMessageMiniprogrampage 卡券类型客服消息
type CustomerMessageMiniprogrampage struct {
	ToUser          string               `json:"touser"` // 接受者OpenID
	Msgtype         MsgType              `json:"msgtype"`
	Miniprogrampage MediaMiniprogrampage `json:"miniprogrampage"` //
}

//MediaMiniprogrampage 小程序消息
type MediaMiniprogrampage struct {
	Title        string `json:"title"`
	Appid        string `json:"appid"`
	Pagepath     string `json:"pagepath"`
	ThumbMediaID string `json:"thumb_media_id"`
}

//SendWithToken 发送卡券类型客服消息
func (msg *CustomerMessageMiniprogrampage) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//sendCustomerMessage 发送客服消息
func sendCustomerMessage(accessToken string, msg interface{}) (err error) {

	uri := fmt.Sprintf("%s?access_token=%s", customerSendMessage, accessToken)
	response, err := util.PostJSON(uri, msg)
	fmt.Println("SendTem uri", uri)
	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("template msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	return nil
}
