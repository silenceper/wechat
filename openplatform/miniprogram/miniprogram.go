package miniprogram

import (
	openContext "github.com/silenceper/wechat/v2/openplatform/context"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/basic"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/code"
	"github.com/silenceper/wechat/v2/openplatform/miniprogram/component"
)

//MiniProgram 代小程序实现业务
type MiniProgram struct {
	AppID       string
	openContext *openContext.Context
}

//NewMiniProgram 实例化
func NewMiniProgram(opCtx *openContext.Context, appID string) *MiniProgram {
	return &MiniProgram{
		openContext: opCtx,
		AppID:       appID,
	}
}

//GetComponent get component
//快速注册小程序相关
func (miniProgram *MiniProgram) GetComponent() *component.Component {
	return component.NewComponent(miniProgram.openContext)
}

//GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetBasic() *basic.Basic {
	return basic.NewBasic(miniProgram.openContext, miniProgram.AppID)
}

//GetBasic 基础信息设置
func (miniProgram *MiniProgram) GetCode() *code.Code {
	return code.NewCode(miniProgram.openContext, miniProgram.AppID)
}
