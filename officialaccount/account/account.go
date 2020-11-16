package account

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)
var(
	getFastRegister = "https://api.weixin.qq.com/cgi-bin/account/fastregister"
)
//Account 账号管理者
type Account struct {
	*context.Context
}

//NewAccountManager 实例化账号管理者
func NewAccountManager(context *context.Context) *Account {
	return &Account{
		context,
	}
}
//FastRegisterRes 快速注册小程序的返回结果
type FastRegisterRes struct {
	util.CommonError
	Appid string `json:"appid"`
	AuthorizationCode string `json:"authorization_code"`
	IsWxVerifySucc bool `json:"is_wx_verify_succ"`
	IsLinkSucc bool `json:"is_link_succ"`

}
//FastRegisterMiniProgram 通过公众号快速注册小程序
func (account *Account) FastRegisterMiniProgram (ticket string)(result *FastRegisterRes,err error){
	var (
		accessToken string
		response []byte
	)
	result = &FastRegisterRes{}
	accessToken, err = account.Context.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf("%s?access_token=%s", getFastRegister, accessToken)

	msg := struct {
		Ticket string `json:"ticket"`
	}{Ticket: ticket}

	response, err = util.PostJSON(uri, msg)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("FastRegisterMiniProgram error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
