package card

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	createCardURL     = "https://api.weixin.qq.com/card/create"
	updateCardUrl     = "https://api.weixin.qq.com/card/update"
	updateUserCardUrl = "https://api.weixin.qq.com/card/membercard/updateuser"
	getTicketURL      = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

func (c *Card) Create(params interface{}) (resp *CreateMemberResponse, err error) {
	var accessToken string
	accessToken, err = c.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", createCardURL, accessToken)
	var resBAry []byte
	resBAry, err = util.PostJSON(uri, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(resBAry, &resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = fmt.Errorf("create member card error : errcode=%v , errmsg=%v", resp.ErrCode, resp.ErrMsg)
		return
	}
	return
}

func (c *Card) Update(params interface{}) ( err error) {
	var accessToken string
	accessToken, err = c.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", updateCardUrl, accessToken)
	_, err = util.PostJSON(uri, params)
	return err
}

func (c *Card) UpdateUserCardInfo(params interface{}) (err error) {
	var accessToken string
	accessToken, err = c.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", updateUserCardUrl, accessToken)
	_, err = util.PostJSON(uri, params)
	return err
}

func (c *Card) GetTicket() (ticket string, err error) {
	var accessToken string
	accessToken, err = c.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s&type=wx_card", getTicketURL, accessToken)
	var bAry []byte
	bAry,err=util.HTTPGet(uri)
	if err != nil {
		return
	}

	var resp GetTicketResponse
	err = json.Unmarshal(bAry, &resp)
	if err != nil {
		return
	}
	if resp.ErrCode != 0 {
		err = fmt.Errorf("get card ticket error : errcode=%v , errmsg=%v", resp.ErrCode, resp.ErrMsg)
		return
	}

	return resp.Ticket, nil
}

