package stat

import (
	"github.com/swxctx/wechat/context"
)

// Stat 统计分析
type Stat struct {
	*context.Context
}

// NewComment 实例化
func NewStat(context *context.Context) *Stat {
	stat := new(Stat)
	stat.Context = context
	return stat
}
