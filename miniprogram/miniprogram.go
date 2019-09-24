package miniprogram

import (
	"github.com/machao520/wechat/wechat"
)

// MiniProgram struct extends context
type MiniProgram struct {
	*wechat.Wechat
}

// NewMiniProgram 实例化小程序接口
func NewMiniProgram(wc *wechat.Wechat) *MiniProgram {
	miniProgram := new(MiniProgram)
	miniProgram.Wechat = wc
	return miniProgram
}
