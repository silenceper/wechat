package card

import (
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
)

const (
	createCardURL     = "https://api.weixin.qq.com/card/create"
	updateCardUrl     = "https://api.weixin.qq.com/card/update"
	updateUserCardUrl = "https://api.weixin.qq.com/membercard/updateuser"
)

type CreateMemberResponse struct {
	util.CommonError
	CardId string `json:"card_id"`
}

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

func (c *Card) UpdateInfo(params interface{}) (err error) {
	var accessToken string
	accessToken, err = c.GetAccessToken()
	if err != nil {
		return
	}

	uri := fmt.Sprintf("%s?access_token=%s", updateUserCardUrl, accessToken)
	_, err = util.PostJSON(uri, params)
	return err
}
