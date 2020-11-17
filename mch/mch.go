package mch

import "github.com/silenceper/wechat/context"

type Mch struct {
	*context.Context
}

func NewTransfers(ctx *context.Context) *Mch {
	mch := Mch{Context: ctx}
	return &mch
}
