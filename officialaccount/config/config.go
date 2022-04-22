package config

import (
	"github.com/silenceper/wechat/v2/core"
)

// Config .config for 微信公众号
type Config struct {
	*core.Config
	Token          string `json:"token"`            // token
	EncodingAESKey string `json:"encoding_aes_key"` // EncodingAESKey
}
