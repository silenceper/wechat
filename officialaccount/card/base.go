package card

import (
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

type Card struct {
	*context.Context
}

func NewCard(context *context.Context) *Card {
	material := new(Card)
	material.Context = context
	return material
}


type CreateMemberResponse struct {
	util.CommonError
	CardId string `json:"card_id"`
}

type GetTicketResponse struct {
	util.CommonError
	Ticket string `json:"ticket" form:"ticket"`
}