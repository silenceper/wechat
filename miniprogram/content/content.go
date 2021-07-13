package content

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	checkTextURL  = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	checkImageURL = "https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s"
)

//Content 内容安全
type Content struct {
	*context.Context
}

//ResContent 请求返回体
type ResContent struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//NewContent 内容安全接口
func NewContent(ctx *context.Context) *Content {
	return &Content{ctx}
}

//CheckText 检测文字
//@text 需要检测的文字
func (content *Content) CheckText(text string) (result ResContent, err error) {
	var accessToken string
	accessToken, err = content.GetAccessToken()
	if err != nil {
		return
	}
	response, err := util.PostJSON(
		fmt.Sprintf(checkTextURL, accessToken),
		map[string]string{
			"content": text,
		},
	)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

//CheckImage 检测图片
//所传参数为要检测的图片文件的绝对路径，图片格式支持PNG、JPEG、JPG、GIF, 像素不超过 750 x 1334，同时文件大小以不超过 300K 为宜，否则可能报错
//@media 图片文件的绝对路径
func (content *Content) CheckImage(media string) (result ResContent, err error) {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return
	}
	response, err := util.PostFile(
		"media",
		media,
		fmt.Sprintf(checkImageURL, accessToken),
	)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
