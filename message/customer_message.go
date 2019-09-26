package message

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/util"
)

const (
	customerSendMessage = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//文本类型客服消息
type CustomerMessageText struct {
	ToUser  string    `json:"touser"`  // 必须, 接受者OpenID
	Msgtype MsgType    `json:"msgtype"` //
	Text    MediaText `json:"text"`    //
}

type MediaText struct {
	Content string `json:"content"` //
}

func NewCustomerTextMessage(toUser, text string) *CustomerMessageText {
	return &CustomerMessageText{
		ToUser:  toUser,
		Msgtype: MsgTypeText,
		Text: MediaText{
			text,
		},
	}
}

//发送文本类型客服消息
func (msg *CustomerMessageText) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//图片类型客服消息
type CustomerMessageImg struct {
	ToUser  string        `json:"touser"`  // 必须, 接受者OpenID
	Msgtype string        `json:"msgtype"` //
	Image   MediaResource `json:"image"`   //
}

type MediaResource struct {
	MediaId string `json:"media_id"` //
}

func NewCustomerImgMessage(toUser, mediaId string) *CustomerMessageImg {
	return &CustomerMessageImg{
		ToUser:  toUser,
		Msgtype: MsgTypeImage,
		Image: MediaResource{
			mediaId,
		},
	}
}

//发送图片类型客服消息
func (msg *CustomerMessageImg) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//语音类型客服消息
type CustomerMessageVoice struct {
	ToUser  string        `json:"touser"`  // 必须, 接受者OpenID
	Msgtype string        `json:"msgtype"` //
	Voice   MediaResource `json:"voice"`   //
}

func NewCustomerVoiceMessage(toUser, mediaId string) *CustomerMessageVoice {
	return &CustomerMessageVoice{
		ToUser:  toUser,
		Msgtype: MsgTypeVoice,
		Voice: MediaResource{
			mediaId,
		},
	}
}
//语音类型客服消息
func (msg *CustomerMessageVoice) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//视频类型客服消息 ，
type CustomerMessageVideo struct {
	ToUser  string     `json:"touser"`  // 必须, 接受者OpenID
	Msgtype string     `json:"msgtype"` //
	Video   MediaVideo `json:"video"`   //
}

type MediaVideo struct {
	MediaId      string `json:"media_id"`
	ThumbMediaId string `json:"thumb_media_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

//视频类型客服消息
func (msg *CustomerMessageVideo) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//音乐类型客服消息
type CustomerMessageMusic struct {
	ToUser  string     `json:"touser"`  // 必须, 接受者OpenID
	Msgtype string     `json:"msgtype"` //
	Music   MediaMusic `json:"music"`   //
}

type MediaMusic struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Musicurl     string `json:"musicurl"`
	Hqmusicurl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

//音乐类型客服消息
func (msg *CustomerMessageMusic) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//图文消息类型客服消息，点击跳转到外链
type CustomerMessageNews struct {
	ToUser  string    `json:"touser"`   // 必须, 接受者OpenID
	Msgtype MsgType    `json:"msgtype"`  //
	News    MediaNews `json:"articles"` //
}

type MediaNews struct {
	Articles []MediaArticles `json:"articles"` //
}

type MediaArticles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

//图文消息类型客服消息，点击跳转到外链
func (msg *CustomerMessageNews) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//图文消息类型客服消息,点击跳转到图文消息页面
type CustomerMessageInnerNews struct {
	ToUser  string        `json:"touser"`  // 必须, 接受者OpenID
	Msgtype string        `json:"msgtype"` //
	News    MediaResource `json:"mpnews"`  //
}

//图文消息类型客服消息,点击跳转到图文消息页面
func (msg *CustomerMessageInnerNews) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}




//菜单类型客服消息
type CustomerMessageMsgmenu struct {
	ToUser  string    `json:"touser"`   // 必须, 接受者OpenID
	Msgtype MsgType    `json:"msgtype"`  //
	News    MediaMsgmenu `json:"msgmenu"` //
}

type MediaMsgmenu struct {
	HeadContent string `json:"head_content"`
	List []MsgmenuItem `json:"list"`
	TailContent string `json:"tail_content"`
}

type MsgmenuItem struct {
	Id       string `json:"id"`
	Content string `json:"content"`
}

//发送菜单类型客服消息
func (msg *CustomerMessageMsgmenu) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//卡券类型客服消息
type CustomerMessageWxcard struct {
	ToUser  string    `json:"touser"`   // 必须, 接受者OpenID
	Msgtype MsgType    `json:"msgtype"`  //
	Wxcard    MediaWxcard `json:"wxcard"` //
}

type MediaWxcard struct {
	CardId string `json:"card_id"`
}

//发送卡券类型客服消息
func (msg *CustomerMessageWxcard) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}

//卡券类型客服消息
type CustomerMessageMiniprogrampage struct {
	ToUser  string    `json:"touser"`   // 必须, 接受者OpenID
	Msgtype MsgType    `json:"msgtype"`  //
	Miniprogrampage    MediaMiniprogrampage `json:"miniprogrampage"` //
}

type MediaMiniprogrampage struct {
	Title string `json:"title"`
	Appid string `json:"appid"`
	Pagepath string `json:"pagepath"`
	ThumbMediaId string `json:"thumb_media_id"`
}

//发送卡券类型客服消息
func (msg *CustomerMessageMiniprogrampage) SendWithToken(accessToken string) (err error) {
	return sendCustomerMessage(accessToken, msg)
}


//发送客服消息
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
