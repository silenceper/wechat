package scheme

import (
	"fmt"
	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	generateUrlScheme = "https://api.weixin.qq.com/wxa/generatescheme?access_token=%v"
)

//UrlScheme struct
type UrlScheme struct {
	*context.Context
}
type MiniProgram struct {
	Path  string `json:"path"`
	Query string `json:"query"`
}
type GenerateBody struct {
	JumpWxa    *MiniProgram `json:"jump_wxa"`
	IsExpire   bool         `json:"is_expire"`
	ExpireTime int64        `json:"expire_time"`
}
type GenerateRes struct {
	util.CommonError
	Openlink string `json:"openlink"`
}

//NewUrlScheme 实例
func NewUrlScheme(context *context.Context) *UrlScheme {
	urlScheme := new(UrlScheme)
	urlScheme.Context = context
	return urlScheme
}

// Generate 获取小程序scheme码
func (urlScheme *UrlScheme) Generate(miniprogram *MiniProgram, isExpire bool, expireTime int64) (rd string, err error) {
	var (
		accessToken string
		urlStr      string
		body        = &GenerateBody{}
		res         = GenerateRes{}
		response    []byte
	)

	accessToken, err = urlScheme.GetAccessToken()
	if err != nil {
		return
	}
	urlStr = fmt.Sprintf(generateUrlScheme, accessToken)
	body.JumpWxa = miniprogram
	body.IsExpire = isExpire
	body.ExpireTime = expireTime
	response, err = util.PostJSON(urlStr, body)
	err = util.DecodeWithError(response, &res, "GenerateUrlScheme")
	rd = res.Openlink
	return

}
