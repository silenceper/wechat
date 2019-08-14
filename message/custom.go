package message

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/context"
	"github.com/silenceper/wechat/util"
)

const (
	sendURL = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

//Custom 客服消息
type Custom struct {
	*context.Context
}

//NewCustom 实例化
func NewCustom(context *context.Context) *Custom {
	tpl := new(Custom)
	tpl.Context = context
	return tpl
}

// CustomArticle 单篇文章
type CustomArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicURL      string `json:"picurl"`
	URL         string `json:"url"`
}

// CustomArticle 单篇文章
type CustomMenuList struct {
	Id       string `json:"id"`
	Content string `json:"content"`
}

//Message 发送的模板消息内容
type CustomMsg struct {
	ToUser     string               `json:"touser"`          // 必须, 接受者OpenID
	MsgType string               `json:"msgtype"`     // 必须, 消息类型
	Text struct {
		Content    string `json:"content"`
	} `json:"text"` //可选, 文本
	Image struct {
		MediaId    string `json:"media_id"`
	} `json:"image"` //可选, 图片
	Voice struct {
		MediaId    string `json:"media_id"`
	} `json:"voice"` //可选, 语音
	Video struct {
		MediaId    string `json:"media_id"`
	} `json:"video"` //可选, 视频
	Music struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		MusicURL      string `json:"musicurl"`
		HQMusicURL      string `json:"hqmusicurl"`
		ThumbMediaId    string `json:"thumb_media_id"`
	} `json:"music"` //可选, 音乐
	News struct {
		Articles     []*CustomArticle `json:"articles"`
	} `json:"news"` //可选, 图文
	MpNews struct {
		MediaId    string `json:"media_id"`
	} `json:"mpnews"` //可选, 公众号图文
	MsgMenu struct {
		HeadContent string `json:"head_content"`
		List []*CustomMenuList `json:"list"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu"` //可选, 菜单消息
	MiniProgramPage struct {
		Title       string `json:"title"`
		Appid string `json:"appid"`
		PagePath       string `json:"pagepath"`
		ThumbMediaId    string `json:"thumb_media_id"`
	} `json:"miniprogrampage"` //可选, 小程序卡片
}

//Send 发送客服消息
func (tpl *Custom) Send(msg *CustomMsg) (err error) {
	var accessToken string
	accessToken, err = tpl.GetAccessToken()
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("%s?access_token=%s", sendURL, accessToken)
	response, err := util.PostJSON(uri, msg)
	if err != nil {
		return err
	}

	var result util.CommonError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.ErrCode != 0 {
		err := fmt.Errorf("custom msg send error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return err
	}
	return nil
}