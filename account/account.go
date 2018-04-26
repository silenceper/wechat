package account

import (
	"github.com/swxctx/wechat/context"
)

// Account 账号管理
type Account struct {
	*context.Context
}

// NewComment 实例化
func NewComment(context *context.Context) *Account {
	account := new(Account)
	account.Context = context
	return account
}
