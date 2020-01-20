package miniprogram

import (
	"sync"

	"github.com/silenceper/wechat/miniprogram/analysis"
	"github.com/silenceper/wechat/miniprogram/auth"
	"github.com/silenceper/wechat/miniprogram/basic"
	"github.com/silenceper/wechat/miniprogram/config"
	"github.com/silenceper/wechat/miniprogram/context"
	"github.com/silenceper/wechat/miniprogram/qrcode"
	"github.com/silenceper/wechat/miniprogram/tcb"
)

//MiniProgram 微信小程序相关API
type MiniProgram struct {
	ctx *context.Context
}

//NewMiniProgram 实例化小程序API
func NewMiniProgram(cfg *config.Config) *MiniProgram {
	if cfg.Cache == nil {
		panic("cache未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	ctx.SetAccessTokenLock(new(sync.RWMutex))
	return &MiniProgram{ctx}
}

// GetContext get Context
func (miniProgram *MiniProgram) GetContext() *context.Context {
	return miniProgram.ctx
}

// GetBasic  基础接口(小程序加解密)
func (miniProgram *MiniProgram) GetBasic() *basic.Basic {
	return basic.NewBasic(miniProgram.ctx)
}

//GetAuth 登录/用户信息相关接口
func (miniProgram *MiniProgram) GetAuth() *auth.Auth {
	return auth.NewAuth(miniProgram.ctx)
}

//GetAnalysis 数据分析
func (miniProgram *MiniProgram) GetAnalysis() *analysis.Analysis {
	return analysis.NewAnalysis(miniProgram.ctx)
}

//GetQRCode 小程序码相关API
func (miniProgram *MiniProgram) GetQRCode() *qrcode.QRCode {
	return qrcode.NewQRCode(miniProgram.ctx)
}

//GetTcb 小程序云开发API
func (miniProgram *MiniProgram) GetTcb() *tcb.Tcb {
	return tcb.NewTcb(miniProgram.ctx)
}
