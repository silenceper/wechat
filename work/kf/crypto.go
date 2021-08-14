package kf

import (
	"errors"

	"github.com/silenceper/wechat/v2/work/crypto"
)

// CryptoOptions 微信服务器验证参数
type CryptoOptions struct {
	Signature string `form:"msg_signature"`
	TimeStamp string `form:"timestamp"`
	Nonce     string `form:"nonce"`
	EchoStr   string `form:"echostr"`
}

// VerifyURL 验证请求参数是否合法
func (r *Client) VerifyURL(options CryptoOptions) (string, error) {
	wxCpt := crypto.NewWXBizMsgCrypt(r.token, r.encodingAESKey, r.corpID, crypto.XMLType)
	data, err := wxCpt.VerifyURL(options.Signature, options.TimeStamp, options.Nonce, options.EchoStr)
	if err != nil {
		return "", errors.New(err.ErrMsg)
	}
	return string(data), nil
}

// DecryptMsg 解密消息
func (r *Client) DecryptMsg(options CryptoOptions, postData []byte) ([]byte, error) {
	wxCpt := crypto.NewWXBizMsgCrypt(r.token, r.encodingAESKey, r.corpID, crypto.XMLType)
	message, status := wxCpt.DecryptMsg(options.Signature, options.TimeStamp, options.Nonce, postData)
	if status != nil && status.ErrCode != 0 {
		return nil, errors.New(status.ErrMsg)
	}
	return message, nil
}
