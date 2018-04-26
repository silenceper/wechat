package comment

import "context"

// Comment 评论、留言管理
type Comment struct {
	*context.Context
}

// NewComment 实例化
func NewComment(context *context.Context) *Comment {
	comment := new(Comment)
	comment.Context = context
	return comment
}
