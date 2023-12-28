package redpacketcover

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/silenceper/wechat/v2/miniprogram/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	getRedPacketCoverURL = "https://api.weixin.qq.com/redpacketcover/wxapp/cover_url/get_by_token?access_token=%s"
)

// RedPacketCover struct
type RedPacketCover struct {
	*context.Context
}

// NewRedPacketCover 实例
func NewRedPacketCover(context *context.Context) *RedPacketCover {
	redPacketCover := new(RedPacketCover)
	redPacketCover.Context = context
	return redPacketCover
}

// CoverParma 小程序码参数
type CoverParma struct {
	// openid 可领取用户的openid
	OpenID string `json:"openid"`
	// ctoken 在红包封面平台获取发放ctoken（需要指定可以发放的appid）
	CToken string `json:"ctoken"`
}

// fetchCode 请求并返回二维码二进制数据
func (qrCode *RedPacketCover) fetchCode(urlStr string, body interface{}) (response []byte, err error) {
	var accessToken string
	accessToken, err = qrCode.GetAccessToken()
	if err != nil {
		return
	}

	urlStr = fmt.Sprintf(urlStr, accessToken)
	var contentType string
	response, contentType, err = util.PostJSONWithRespContentType(urlStr, body)
	if err != nil {
		return
	}
	if strings.HasPrefix(contentType, "application/json") {
		// 返回错误信息
		var result util.CommonError
		err = json.Unmarshal(response, &result)
		if err == nil && result.ErrCode != 0 {
			err = fmt.Errorf("error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
			return nil, err
		}

		return response, nil
	}
	err = fmt.Errorf("error : unknown response content type - %v", contentType)
	return nil, err
}

// GetRedPacketCoverURL 获得指定用户可以领取的红包封面链接。获取参数ctoken参考微信红包封面开放平台
// 文档地址： https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/red-packet-cover/getRedPacketCoverUrl.html
func (qrCode *RedPacketCover) GetRedPacketCoverURL(coderParams CoverParma) (response []byte, err error) {
	return qrCode.fetchCode(getRedPacketCoverURL, coderParams)
}
