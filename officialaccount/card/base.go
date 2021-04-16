package card

import "github.com/silenceper/wechat/v2/officialaccount/context"

type Card struct {
	*context.Context
}

func NewCard(context *context.Context) *Card {
	material := new(Card)
	material.Context = context
	return material
}
