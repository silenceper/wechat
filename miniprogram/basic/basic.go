package basic

import "github.com/silenceper/wechat/v2/miniprogram/context"

//Basic struct
type Basic struct {
	*context.Context
}

//NewBasic 实例
func NewBasic(context *context.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}
