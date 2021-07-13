package content_security

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	checkTextUrl  = "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s"
	checkImageUrl = "https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s"
)

//内容安全
type ContentSecurity struct {
	*context.Context
}

type ContentSecurityRes struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

//NewContentSecurity
func NewContentSecurity(ctx *context.Context) *ContentSecurity {
	return &ContentSecurity{ctx}
}

//检测文字
//@text 需要检测的文字
func (content *ContentSecurity) CheckText(text string) (result ContentSecurityRes, err error) {
	var accessToken string
	accessToken, err = content.GetAccessToken()
	if err != nil {
		return
	}
	response, _, err := util.PostJSONWithRespContentType(
		fmt.Sprintf(checkTextUrl, accessToken),
		map[string]string{
			"content": text,
		},
	)
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}

//检测图片
//所传参数为要检测的图片文件的绝对路径，图片格式支持PNG、JPEG、JPG、GIF, 像素不超过 750 x 1334，同时文件大小以不超过 300K 为宜，否则可能报错
//@media 图片文件的绝对路径
func (content *ContentSecurity) CheckImage(media string) (result ContentSecurityRes, err error) {
	accessToken, err := content.GetAccessToken()
	if err != nil {
		return
	}
	response, err := util.PostFile(
		"media",
		media,
		fmt.Sprintf(checkImageUrl, accessToken),
	)
	fmt.Println(string(response))
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	return
}
