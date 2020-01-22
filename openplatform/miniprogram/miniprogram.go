package miniprogram

import (
	opContext "github.com/silenceper/wechat/openplatform/context"
	"github.com/silenceper/wechat/openplatform/miniprogram/basic"
	"github.com/silenceper/wechat/openplatform/miniprogram/component"
)

//MiniProgram 代小程序实现业务
type MiniProgram struct {
	AppID     string
	opContext *opContext.Context
}

//NewMiniProgram 实例化
func NewMiniProgram(opCtx *opContext.Context, appID string) *MiniProgram {
	return &MiniProgram{
		opContext: opCtx,
		AppID:     appID,
	}
}

//GetComponent get component
//快速注册小程序相关
func (miniProgram *MiniProgram) GetComponent() *component.Component {
	return component.NewComponent(miniProgram.opContext)
}

//GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetBasic() *basic.Basic {
	return basic.NewBasic(miniProgram.opContext, miniProgram.AppID)
}
