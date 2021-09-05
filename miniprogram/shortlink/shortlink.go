package shortlink

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
	"strings"
)

const (
	generateShortLinkURL = "https://api.weixin.qq.com/wxa/genwxashortlink?access_token=%s"
)

type ShortLink struct {
	*context.Context
}

// NewShortLink 实例
func NewShortLink(ctx *context.Context) *ShortLink {
	return &ShortLink{ctx}
}

type ShortLinker struct {

	// pageUrl 通过 Short Link 进入的小程序页面路径，必须是已经发布的小程序存在的页面，可携带 query，最大1024个字符
	PageUrl string `json:"page_url"`

	// pageTitle 页面标题，不能包含违法信息，超过20字符会用... 截断代替
	PageTitle string `json:"page_title"`

	// isPermanent 生成的 Short Link 类型，短期有效：false，永久有效：true
	IsPermanent bool `json:"is_permanent,omitempty"`

}

// Generate 生成 shortLink
func (shortLink *ShortLink) generate(shortLinkParams ShortLinker) (response []byte, err error) {
	var accessToken string
	accessToken, err = shortLink.GetAccessToken()
	if err != nil {
		return
	}

	urlStr := fmt.Sprintf(generateShortLinkURL, accessToken)
	var contentType string
	response, contentType, err = util.PostJSONWithRespContentType(urlStr, shortLinkParams)
	if err != nil {
		return
	}
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result util.CommonError
		err = json.Unmarshal(response, &result)
		if err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("fetchCode error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return
		}
	}
	return response,nil
}

// GenerateShortLinkPermanent 生成永久shortLink
func (shortLink *ShortLink) GenerateShortLinkPermanent(pageUrl,pageTitle string) (response []byte, err error) {
	return shortLink.generate(ShortLinker{
		PageUrl:     pageUrl,
		PageTitle:   pageTitle,
		IsPermanent: true,
	})
}

// GenerateShortLinkTemp 生成临时shortLink
func (shortLink *ShortLink) GenerateShortLinkTemp(pageUrl,pageTitle string) (response []byte, err error) {
	return shortLink.generate(ShortLinker{
		PageUrl:     pageUrl,
		PageTitle:   pageTitle,
		IsPermanent: false,
	})
}