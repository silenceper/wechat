package basic

import "github.com/silenceper/wechat/v2/officialaccount/context"

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
