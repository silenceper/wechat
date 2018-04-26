package account

import (
	"github.com/swxctx/wechat/context"
)

// Account 账号管理
type Account struct {
	*context.Context
}

// NewAccount 实例化
func NewAccount(context *context.Context) *Account {
	account := new(Account)
	account.Context = context
	return account
}
