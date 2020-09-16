package message

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/officialaccount/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	setIndustryURL = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"
	getIndustryURL = "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s"
)

//Industry 行业信息
type Industry struct {
	*context.Context
}

//NewIndustry 实例化
func NewIndustry(context *context.Context) *Industry {
	industry := new(Industry)
	industry.Context = context
	return industry
}
//SetIndustryResult 设置行业的返回消息
type SetIndustryResult struct {
	util.CommonError
}
//GetIndustryResult 获取行业的返回消息
type GetIndustryResult struct {
	PrimaryIndustry struct{
		FirstClass string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"primary_industry"`
	SecondaryIndustry struct{
		FirstClass string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"secondary_industry"`
}


//SetIndustry 设置行业
func (industry *Industry) SetIndustry(industryId1, industryId2 int64) (result *SetIndustryResult, err error) {
	var accessToken string
	var data = struct {
		IndustryId1 int64 `json:"industry_id1"`
		IndustryId2 int64 `json:"industry_id2"`
	}{IndustryId1: industryId1, IndustryId2: industryId2}
	accessToken, err = industry.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(setIndustryURL, accessToken)
	result = &SetIndustryResult{}
	response, err := util.PostJSON(uri, data)
	if err != nil {
		return
	}

	err = util.DecodeWithError(response, result, "SetIndustry")

	//msgID = result.MsgID
	return
}

//GetIndustry 设置行业
func (industry *Industry) GetIndustry() (result *GetIndustryResult, err error) {
	var accessToken string

	accessToken, err = industry.GetAccessToken()
	if err != nil {
		return
	}
	uri := fmt.Sprintf(getIndustryURL, accessToken)
	result = &GetIndustryResult{}
	response, err := util.PostJSON(uri, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, result)

	//msgID = result.MsgID
	return
}
