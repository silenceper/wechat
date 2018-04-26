package user

import (
	"github.com/swxctx/wechat/context"
	"github.com/swxctx/wechat/util"
)

const (
	userInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info"
)

//User 用户管理
type User struct {
	*context.Context
}

//NewUser 实例化
func NewUser(context *context.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

//Info 用户基本信息
type Info struct {
	util.CommonError

	Subscribe     int32    `json:"subscribe"`
	OpenID        string   `json:"openid"`
	Nickname      string   `json:"nickname"`
	Sex           int32    `json:"sex"`
	City          string   `json:"city"`
	Country       string   `json:"country"`
	Province      string   `json:"province"`
	Language      string   `json:"language"`
	Headimgurl    string   `json:"headimgurl"`
	SubscribeTime int32    `json:"subscribe_time"`
	UnionID       string   `json:"unionid"`
	Remark        string   `json:"remark"`
	GroupID       int32    `json:"groupid"`
	TagidList     []string `json:"tagid_list"`
}
